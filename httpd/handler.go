package httpd

import (
	"fmt"
	"log"
	"net"
	"os"
)

func HttpHandler(conn net.Conn, dtdFile string) error {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	log.Printf("[HTTP] Connection built with %s\n", remoteAddr)

	buf := make([]byte, 4096)
	_, err := conn.Read(buf)
	if err != nil {
		return err
	}

	log.Printf("[HTTP] Receive request : \n%s\n", buf)

	dtd, err := os.ReadFile(dtdFile)
	if err != nil {
		return err
	}

	resp := []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/xml\r\nContent-length: %d\r\n\r\n%s\r\n\r\n", len(dtd), string(dtd)))
	_, err = conn.Write(resp)
	if err != nil {
		return err
	}

	log.Printf("[HTTP] Send response : \n%s\n", resp)

	return nil
}
