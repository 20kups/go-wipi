# go-wipi

A golang wrapper around the C library [WiringPi](https://github.com/WiringPi/WiringPi)

## Setup

First, follow their guide to install [WiringPi](https://github.com/WiringPi/WiringPi). A `libwiringPI.so` is expected under `/usr/lib`

```
$ lh /usr/lib | grep wiringPi
lrwxrwxrwx  1 root root   37 Jan 01 XX:XX libwiringPiDev.so -> /usr/local/lib/libwiringPiDev.so.3.16
lrwxrwxrwx  1 root root   34 Jan 01 XX:XX libwiringPi.so -> /usr/local/lib/libwiringPi.so.3.16
```

Then install `github.com/20kups/go-wipi` in your project

```bash
go get github.com/20kups/go-wipi
```