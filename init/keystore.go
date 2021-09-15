package init

import (
	"amnesia-db/internal"
	"sync"
)

func KeyStore() {
	internal.GlobalKeyStore.Init()
	go internal.GlobalKeyStore.NFetchWorker()
	wg := sync.WaitGroup{}
	wg.Wait()
}
