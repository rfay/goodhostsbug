package main

import (
	"fmt"
	"github.com/goodhosts/hostsfile"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <ip_address> <hostname>\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]
	hostname := os.Args[2]

	hf, err := hostsfile.NewCustomHosts("/etc/hosts")
	if err != nil {
		panic(err)
	}

	err = hf.Add(ip, hostname)
	if err != nil {
		panic(err)
	}

	// If HostsPerLine() is called before hf.Add(), we get a panic on hf.Add().
	hf.HostsPerLine(8)

	err = hf.Flush()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hosts file updated successfully!")
}
