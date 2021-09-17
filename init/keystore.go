package init

import (
	"amnesia-db/constants"
	"amnesia-db/internal"
	"fmt"
	"sync"
)

func KeyStore() {
	fmt.Print("Loading -> Keystore...")
	internal.GlobalKeyStore.Init()
	fmt.Println("DONE")

	wg := sync.WaitGroup{}
	for i := 0; i < constants.NNFetchWorkers; i++ {
		wg.Add(1)
		go internal.GlobalKeyStore.NFetchWorker(i, &wg)
	}
	wg.Wait()
}
