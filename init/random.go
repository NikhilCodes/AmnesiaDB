package init

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomSeed() {
	fmt.Print("Randomizing Seed...")
	rand.Seed(time.Now().Unix())
	fmt.Println("DONE")
}
