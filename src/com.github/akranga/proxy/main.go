package main

import (
	"bufio"
	"flag"
	"fmt"
	//	"io"
	"log"
	"net"
	//	"net/http"
	"os"
	"regexp"
)

var (
	url     = os.Getenv("PROBE_ENDPOINT")
	backend = os.Getenv("CONSUL_URL")
)

func exitGracefully(err error) {
	fmt.Fprintf(os.Stdout, "unknown")
	os.Exit(0) // we always exit normally to avoid confusion of consul-proxy
}

func probeWWW(conn net.Conn) bool {
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		exitGracefully(err)
	}
	log.Println("status: ", status)
	matched, err := regexp.MatchString("HTTP/\\d?", status)
	if err != nil {
		exitGracefully(err)
	}
	return matched
}

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var (
	versionFlag *bool = flag.Bool("v", false, "Print the version number.")
)

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	if url == "" {
		url = "localhost:8080"
	}
	conn, err := net.Dial("tcp", url)
	log.Println("status: ", conn)
	if err != nil {
		exitGracefully(err)
	}

	if probeWWW(conn) {
		fmt.Fprintf(os.Stdout, "www")
	} else {
		fmt.Fprintf(os.Stdout, "unknown")
	}
}
