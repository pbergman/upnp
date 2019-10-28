## UPnP

An simple UPnP discovery implementation that uses the standard http.Request and http.Response for sending and receiving available devices.

#### usage

```
# this can also be nil
ifi, err := net.InterfaceByName("enp3s0")

if nil != err {
    panic(err)
}

handler := func(resp *http.Response) {
    out, _ := httputil.DumpResponse(response, true)
    fmt.Println(string(out))
}

if err := upnp.Discover(context.Background(), handler, 5*time.Second, ifi, nil, nil); err != nil {
    panic(err)
}


```

by default it will search for all UPnP devices and services but can be changed with the headers argument:

```
if err := upnp.Discover(context.Background(), handler, 5*time.Second, nil, map[string]string{"ST": "upnp:rootdevice"}, nil); err != nil {
    panic(err)
}
```

to debug input/output for connection you can give the last argument the os.Stdout or any other writer to log the request and response.
