package init

import (
	"amnesia-db/config"
	"amnesia-db/internal"
	"fmt"
	"net"
	"strconv"
)

func Socket() {
	fmt.Print("Initializing TCP Listener...")
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(config.GetConfig().Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("CONNECTION__CLOSED")
	}(l)

	fmt.Println("DONE")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go internal.HandleConnection(c)
	}
}
