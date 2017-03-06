# ifconfig.cat
ifconfig.cat source code

![](https://raw.githubusercontent.com/mengzhuo/ifconfig.cat/master/ifc_screenshot.png)

### GeoLite DB

http://dev.maxmind.com/geoip/geoip2/geolite2/

### Build

```
make
```

### Pkg

```
make deb
make rpm
```

### Params
```
Ifconfig.cat service version:v0.1
  -addr string
    	listen addr (default ":8080")
  -cert string
    	TLS certfile
  -favicon string
    	favicon path (default "favicon.ico")
  -geo string
    	geo DB
  -key string
    	TLS keyfile
  -prometheus string
    	prometheus listener
  -tlsaddr string
    	tls listen addr
  -tmpl string
    	template path glob (default "templates/*.tpl")
  -v	version
```
