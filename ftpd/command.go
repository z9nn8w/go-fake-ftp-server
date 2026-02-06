package ftpd

import (
	"net"
	"strings"
)

func HandleCommand(conn net.Conn, cmd string) {
	switch {
	case strings.HasPrefix(cmd, "USER"):
		conn.Write([]byte("331 Please specify the password.\r\n"))

	case strings.HasPrefix(cmd, "PASS"):
		conn.Write([]byte("230 Login successful.\r\n"))

	case strings.HasPrefix(cmd, "CWD"):
		conn.Write([]byte("250 Directory changed successfully.\r\n"))

	case strings.HasPrefix(cmd, "SYST"):
		conn.Write([]byte("215 UNIX Type: L8\r\n"))

	case strings.HasPrefix(cmd, "TYPE"):
		conn.Write([]byte("200 Switching to Binary mode.\r\n"))

	case strings.HasPrefix(cmd, "OPTS"):
		conn.Write([]byte("200 Always in UTF8 mode.\r\n"))

	case strings.HasPrefix(cmd, "LIST"):
		conn.Write([]byte("150 Here comes the directory listing\r\n"))
		conn.Write([]byte("drwxrwxrwx 1 user group 4096 Feb 05 1.txt\r\n"))
		conn.Write([]byte("226 Transfer complete\r\n"))

	case strings.HasPrefix(cmd, "EPSV"):
		conn.Write([]byte("229 Entering Extended Passive Mode (|||2121|)\r\n"))

	case strings.HasPrefix(cmd, "RETR"):
		conn.Write([]byte("550 File not found.\r\n"))

	case strings.HasPrefix(cmd, "QUIT"):
		conn.Write([]byte("221 Goodbye.\r\n"))
		conn.Close()

	default:
		conn.Write([]byte("230 more data please!\r\n"))
	}
}
