package main

import (
	"fmt"
	"github.com/goodhosts/hostsfile"
)

func main() {
	hf, err := hostsfile.NewCustomHosts("/etc/hosts")
	if err != nil {
		panic(err)
	}

	hf.HostsPerLine(8)
	err = hf.Add("127.0.0.1", "something.example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("Hosts file updated successfully!")
}
