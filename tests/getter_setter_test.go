package tests

import (
	sysInit "amnesia-db/init"
	"amnesia-db/internal"
	"testing"
	"time"
)

func TestGetterSetter(t *testing.T) {
	sysInit.KeyStore()
	q := internal.ParseFromString("SET a AS d WHERE NFETCH=2")
	internal.GlobalKeyStore.Insert(q.Key, q.Value, q.Options)
	out := internal.GlobalKeyStore.Get("a")
	if out != "d" {
		t.Errorf("Test Failed on GET 1st Iter: %s inputted, %s expected, %s received", "a", "d", out)
	}

	out = internal.GlobalKeyStore.Get("a")
	if out != "d" {
		t.Errorf("Test Failed on GET 2nd Iter: %s inputted, %s expected, %s received", "a", "d", out)
	}

	time.Sleep(time.Duration(2) * time.Second)
	out = internal.GlobalKeyStore.Get("a")
	if out != "<NIL>" {
		t.Errorf("Test Failed on GET 3rd Iter: %s inputted, %s expected, %s received", "a", "<NIL>", out)
	}
}