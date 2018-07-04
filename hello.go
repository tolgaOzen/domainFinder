package main

// import (
// 	"fmt"
// 	"net"
// )

// func handler(c net.Conn, ch chan string) {
// 	ch <- c.RemoteAddr().String()
// 	c.Write([]byte("ok"))
// 	c.Close()
// }

// func logger(ch chan string) {
// 	for {
// 		fmt.Println(<-ch)
// 	}
// }

// func server(l net.Listener, ch chan string) {
// 	for {
// 		c, err := l.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		go handler(c, ch)
// 	}
// }

// func main() {
// 	l, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		panic(err)
// 	}
// 	ch := make(chan string)
// 	go server(l, ch)
// 	logger(ch)

// 	//time.Sleep(10 * time.Second)
// }
