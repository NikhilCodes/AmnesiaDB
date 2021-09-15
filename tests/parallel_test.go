package tests

import (
	sysInit "amnesia-db/init"
	"amnesia-db/internal"
	"fmt"
	"testing"
)

//func TestMain(m *testing.M) {
//}

func TestParallel1(t *testing.T) {
	fmt.Println("STARTING1")
	sysInit.KeyStore()
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel2(t *testing.T) {
	fmt.Println("STARTING2")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel3(t *testing.T) {
	fmt.Println("STARTING3")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel4(t *testing.T) {
	fmt.Println("STARTING4")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel5(t *testing.T) {
	fmt.Println("STARTING5")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel6(t *testing.T) {
	fmt.Println("STARTING6")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel7(t *testing.T) {
	fmt.Println("STARTING7")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel8(t *testing.T) {
	fmt.Println("STARTING8")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel9(t *testing.T) {
	fmt.Println("STARTING9")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}

func TestParallel10(t *testing.T) {
	fmt.Println("STARTING10")
	t.Parallel()
	//q1 := internal.ParseFromString("SET a AS b WHERE TTL=10hr NFETCH=4")
	q2 := internal.ParseFromString("SET a AS d WHERE TTL=10hr NFETCH=2")

	//internal.GlobalKeyStore.Insert(q1.Key, q1.Value, q1.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	internal.GlobalKeyStore.Insert(q2.Key, q2.Value, q2.Options)
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
	fmt.Println(internal.GlobalKeyStore.Get("a"))
}
