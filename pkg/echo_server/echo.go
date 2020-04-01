package echo_server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Echo() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		panic("Could connect to server %s")
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("REQUEST>")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text)
		go copyMsg(conn, os.Stdout)
	}
}

func copyMsg(src io.ReadCloser, dst io.WriteCloser) {
	defer func() {
		src.Close()
		dst.Close()
	}()
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Errorf("Error when copy %s", err)
	}
}
