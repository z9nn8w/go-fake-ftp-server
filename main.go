package main

import (
	"flag"
	"go-fake-ftp-server/ftpd"
	"go-fake-ftp-server/httpd"
	"log"
	"net"
)

func main() {
	httpPort := flag.String("httpport", "8008", "HTTP server port")
	ftpPort := flag.String("ftpport", "2121", "FTP server port")
	dtdFile := flag.String("file", "evil.dtd", "evil dtd file")
	flag.Parse()

	go startHTTP(*httpPort, *dtdFile)
	go startFTP(*ftpPort)

	select {} // 阻塞主 goroutine
}

func startHTTP(port string, dtdFile string) {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[HTTP] Server starts on port %s\n", port)

	for {
		conn, _ := l.Accept()
		go httpd.HttpHandler(conn, dtdFile)
	}
}

func startFTP(port string) {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[FTP] Server starts on port %s\n", port)

	for {
		conn, _ := l.Accept()
		go ftpd.FTPHandler(conn)
	}
}
