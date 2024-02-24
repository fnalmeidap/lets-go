# Go RPC

If you get the following error:
```
client.go:l:c: package gorpc/impl is not in std (/usr/local/go/src/gorpc/impl)
```
That is because `gorpc` project has one internal module (`impl`) that defines communication structure and interface.
You need initialize your project as a **Go module** to be able to use these definitions. To do so, go to `/gorpc` and use:

```bash
    go mod init gorpc
```