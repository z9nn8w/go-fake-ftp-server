package ftpd

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func FTPHandler(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	log.Printf("[FTP] Connection built with %s\n", remoteAddr)

	banner := []byte("220 ftp-server\r\n")
	conn.Write(banner)

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("[FTP] Client disconnected")
			return
		}

		line = strings.TrimSpace(line)
		log.Printf("[FTP] %s\n", line)

		HandleCommand(conn, line)
	}
}
