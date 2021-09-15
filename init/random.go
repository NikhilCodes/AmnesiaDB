package init

import (
	"math/rand"
	"time"
)

func RandomSeed() {
	rand.Seed(time.Now().Unix())
}
