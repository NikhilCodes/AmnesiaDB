package internal

import (
	"amnesia-db/constants"
	"fmt"
	"sync"
	"time"
)

type ConcurrentHashMap struct {
	data map[string]Block
	mu   sync.RWMutex
}

type Block struct {
	Value         Value
	NFetch        *int32
	IsUsingNFetch bool
	Expiry        time.Time
	IsUsingExpiry bool
	mu            *sync.RWMutex
}

func (chm *ConcurrentHashMap) Init() {
	chm.data = make(map[string]Block)
}

func (chm *ConcurrentHashMap) Insert(key Key, value Value, options Options) {
	var (
		nFetch      OptionValue
		ttl         OptionValue
		foundNFetch bool
		foundTTL    bool
	)

	nFetch, foundNFetch = options[NFETCH]
	ttl, foundTTL = options[TTL]
	delay := time.Second * -1
	if foundTTL {
		switch ttl.value {
		case "ms":
			delay = time.Millisecond * time.Duration(ttl.intValue)
			break

		case "s":
			delay = time.Second * time.Duration(ttl.intValue)
			break

		case "hr":
			delay = time.Hour * time.Duration(ttl.intValue)
			break

		case "d":
			delay = time.Hour * 24 * time.Duration(ttl.intValue)
			break
		}
	}

	mutex := &sync.RWMutex{}
	nFetchValue := &nFetch.intValue
	chm.mu.Lock()
	chm.data[string(key)] = Block{
		Value:         value,
		NFetch:        nFetchValue,
		Expiry:        time.Now().Add(delay),
		IsUsingNFetch: foundNFetch,
		IsUsingExpiry: foundTTL,
		mu:            mutex,
	}
	chm.mu.Unlock()
}

func CheckValueValidity(value Block) bool {
	if value.IsUsingExpiry && !value.Expiry.After(time.Now()) {
		return false
	}

	if value.IsUsingNFetch && *(value.NFetch) <= 0 {
		return false
	}

	return true
}

func (chm *ConcurrentHashMap) Get(key Key) Value {
	chm.mu.RLock()
	fetch, ok := chm.data[string(key)]
	isValid := CheckValueValidity(fetch)
	chm.mu.RUnlock()
	if ok {
		if isValid {
			NFetchWorkerPipe <- key
			return fetch.Value
		} else {
			chm.mu.Lock()
			delete(chm.data, string(key))
			chm.mu.Unlock()
		}
	}

	return constants.NIL
}

var NFetchWorkerPipe = make(chan Key)

func (chm *ConcurrentHashMap) NFetchWorker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Starting <> NFetchWorker #%d\n", id)
	wg.Done()
	for {
		key := <-NFetchWorkerPipe
		chm.mu.RLock()
		data := chm.data[string(key)]
		isUsingNFetch := data.IsUsingNFetch
		chm.mu.RUnlock()

		if isUsingNFetch {
			chm.mu.Lock()
			// TODO: Possible Bottleneck
			if chm.data[string(key)].NFetch != nil {
				*(chm.data[string(key)].NFetch)--
			}
			chm.mu.Unlock()
		}
	}
}
