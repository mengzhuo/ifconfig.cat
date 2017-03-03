VERSION := $(shell git describe --tags)
DESTDIR?=.tmpBuildRoot

.PHONY: binary
binary: clean build

.PHONY: clean
clean:
	rm -rf *.deb
	rm -rf *.rpm
	rm -rf ifc
	rm -rf ${DESTDIR} 

.PHONY: build
build:
	go build -o ifc -ldflags '-X main.Version=${VERSION}' 

.PHONY: pkg
pkg:
	rm -rf ${DESTDIR}
	mkdir ${DESTDIR}
	mkdir -p ${DESTDIR}/usr/share/ifc/templates
	cp -r templates/*.tpl ${DESTDIR}/usr/share/ifc/templates/
	mkdir -p ${DESTDIR}/usr/local/bin
	cp ifc ${DESTDIR}/usr/local/bin/
	mkdir -p ${DESTDIR}/lib/systemd/system/
	cp ifc.service ${DESTDIR}/lib/systemd/system/
	mkdir -p ${DESTDIR}/etc/default/
	cp default ${DESTDIR}/etc/default/ifc
	mkdir -p ${DESTDIR}/var/log/ifc
	mkdir -p ${DESTDIR}/etc/logrotate.d/
	cp logrote ${DESTDIR}/etc/logrotate.d/ifc

deb: clean build pkg
	fpm -t deb -s dir -n ifc -v $(VERSION:v%=%) -C ${DESTDIR}

rpm: clean build pkg
	fpm -t rpm -s dir -n ifc -v ${VERSION} -C ${DESTDIR} 
