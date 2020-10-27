
**Install Golang on Linux**

```
sudo apt install golang
```

**Upgrade Golang manually**

```
# 1/ download the desired packages from (for example):
# https://launchpad.net/ubuntu/+source/golang-1.14
# 2/ install them:
sudo dpkg -i *.deb
# 3/ redirect links to the new version:
sudo rm /usr/lib/go && sudo ln -s go-1.14 /usr/lib/go
sudo ln -sf ../lib/go/bin/go /usr/bin/go
sudo ln -sf ../lib/go/bin/gofmt /usr/bin/gofmt
```

**[Install TinyGo on Linux](https://tinygo.org/getting-started/linux/)**

```
wget https://github.com/tinygo-org/tinygo/releases/download/v0.15.0/tinygo_0.15.0_amd64.deb
sudo dpkg -i tinygo_0.15.0_amd64.deb
```

**Setup your `/etc/hosts`**

```
### How to Go ###

127.1.0.0       how-srv-hello
127.1.1.0       how-srv-simple
127.1.2.0       how-srv-system0
127.1.2.1       how-srv-system1
127.1.3.0       how-srv-crud0
127.1.3.1       how-srv-crud1
```
