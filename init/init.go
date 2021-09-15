package init

import "amnesia-db/config"

func All() {
	config.InitConfig()

	KeyStore()
	RandomSeed()
	Socket()
}
