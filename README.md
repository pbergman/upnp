## UPnP

An simple UPnP discovery implementation that uses the standard http.Request and http.Response for sending and receiving available devices.

#### usage

```
# this can also be nil
ifi, err := net.InterfaceByName("enp3s0")

if nil != err {
    panic(err)
}

var queue = make(chan *http.Response, 10)
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for response := range queue {
            out, _ := httputil.DumpResponse(response, true)
            fmt.Println(string(out))
        }
    }()
}

if err := upnp.Discover(queue, 5*time.Second, ifi, nil, nil); err != nil {
    panic(err)
}

close(queue)
wg.Wait()


```

by default it will search for all UPnP devices and services but can be changed with the headers argument:

```
if err := upnp.Discover(queue, 5*time.Second, nil, map[string]string{"ST": "upnp:rootdevice"}, nil); err != nil {
    panic(err)
}
```

to debug input/output for connection you can give the last argument the os.Stdout or any other writer to log the request and response.
