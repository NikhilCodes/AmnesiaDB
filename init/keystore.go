package init

import (
	"amnesia-db/constants"
	"amnesia-db/internal"
	"sync"
)

func KeyStore() {
	internal.GlobalKeyStore.Init()
	for i := 0; i < constants.NNFetchWorkers; i++ {
		go internal.GlobalKeyStore.NFetchWorker()
	}
	wg := sync.WaitGroup{}
	wg.Wait()
}
