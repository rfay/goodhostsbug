# goodhostsbug
Demonstrate goodhosts bug

Build with `go install`

Run with `goodhostsbug`

macOS: Install the default macOS /etc/hosts file from [here](hosts)

On the "broken" branch you get this:

Result is 
```
$ goodhostsbug
panic: runtime error: index out of range [6] with length 4

goroutine 1 [running]:
github.com/goodhosts/hostsfile.(*Hosts).Add(0x140002ad8f0, {0x10109e3a0, 0x9}, {0x140001d3850?, 0x1, 0x1})
	/Users/rfay/workspace/goodhostsbug/vendor/github.com/goodhosts/hostsfile/hosts.go:167 +0x3ec
main.main()
	/Users/rfay/workspace/goodhostsbug/main.go:15 +0x80
```

On the "main" branch, with the HostsPerLine() later, you get predictable behavior.
