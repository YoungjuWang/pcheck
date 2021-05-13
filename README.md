### Info
---

'pcheck'는 tcp port가 open 됐는지 확인하는 간단한 프로그램입니다.

### Download / Installation
---

```
# wget https://raw.githubusercontent.com/YoungjuWang/pcheck/master/pcheck/pcheck
# chmod +x pcheck; mv pcheck /usr/local/bin/
```

### Help
---

```
# pcheck --help
"pcheck" check certain TCP port is opend

Usage:
  pcheck [flags]

Examples:
pcheck -i 192.168.0.10 -p 8888

Flags:
  -h, --help          help for pcheck
  -i, --host string   destination address for checking (required)
  -p, --port string   destination port for checking (required)
```

### Expected Output
---

```
# pcheck -i 192.168.122.1 -p 8888
192.168.122.1:8888 is opened

# pcheck -i 192.168.122.1 -p 8080
dial tcp 192.168.122.1:8080: connect: connection refused
```
