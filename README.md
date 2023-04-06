# goodhostsbug
Demonstrate goodhosts bug

Build with `go install`

Run with `goodhostsbug`

macOS: Install the default /etc/hosts file (don't forget to save your normal one first)
```
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost

```

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
