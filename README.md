# goodhostsbug
Demonstrate goodhosts comments bug

`go get -u && go mod vendor && go mod tidy`

Build with `go install`

macOS: Install the default macOS /etc/hosts file from [here](hosts)

Run with `sudo ~/go/bin/goodhostsbug 127.0.0.1 example.com`

Now inspect the /etc/hosts file. It will have incorrect information, with the original comments mangled and mostly destroyed, and in the wrong place.

The original /etc/hosts was 
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

But now it's pretty destroyed, with only one small part of comments remaining, in the wrong place:

```
127.0.0.1 localhost xxx.ddev.site
255.255.255.255 broadcasthost
::1 localhost
  # Host Database
```
