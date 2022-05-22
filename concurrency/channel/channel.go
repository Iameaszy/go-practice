package channel

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		//done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	defer conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
