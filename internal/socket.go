package internal

import (
	"amnesia-db/constants"
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(c net.Conn) {
	pipe := make(chan string)
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		queryString := strings.TrimSpace(netData)
		if queryString == constants.EXIT {
			break
		}

		go Eval(queryString, pipe)
		result := <- pipe + "\r\n"
		_, err = c.Write([]byte(result))
		if err != nil {
			println(err)
		}
	}
	err := c.Close()
	if err != nil {
		println(err)
	}
}
