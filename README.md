# Dependencies

* libpcap

# Build
```
go build -ldconfig '-s -w' -o mcf main.go
```


# Examples

Forward SSDP from eth0 to eth1
```
./mcf \
  -bpf '(dst host 239.255.255.250 or dst host FF02::C or dst host FF05::C or dst host FF08::C or dst host FF0E::C) and dst port 1900 and udp' \
  -read-if eth0 \
  -write-if eth1
```

Forward multicast CoAP
```
./mcf \
  -bpf 'dst host 224.0.1.187 and dst port 5683 and udp' \
  -read-if eth0 \
  -write-if eth1
```
