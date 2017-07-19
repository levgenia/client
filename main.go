package main

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"flag"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "4001", "port for requesting")
}

func main() {
	flag.Parse()

	fmt.Println("Let a magic begins!")


	reader := bufio.NewReader(os.Stdin)
	for {
		// Read in an integer
		data, _ := reader.ReadString('\n')
		if len(data) > 2 {
			resp, err := http.Post("http://localhost:" + port +"/db/foo", "text/plain; charset=utf-8", bytes.NewReader([]byte(data)))
			if err != nil {
				fmt.Print(err)
			}
			resp.Body.Close()
		}

		resp, err := http.Get("http://localhost:4001/db/foo")
		if err != nil {
			fmt.Print(err)
		}
		dd, _ := ioutil.ReadAll(resp.Body)
		os.Stdout.Write(dd)
		resp.Body.Close()
	}
}
