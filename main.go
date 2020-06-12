package main

import (
	"flag"
	"github.com/google/gopacket/pcap"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	flgReadIf := flag.String("read-if", "eth0", "Interface to read from")
	flgWriteIf := flag.String("write-if", "eth1", "Interface to write on")
	flgBPF := flag.String("bpf", "", "BPF filter")
	flag.Parse()

	read_h, err := pcap.OpenLive(*flgReadIf, 1024, false, pcap.BlockForever)
	check(err)
	check(read_h.SetDirection(pcap.DirectionIn))
	defer read_h.Close()
	write_h, err := pcap.OpenLive(*flgWriteIf, 1024, false, pcap.BlockForever)
	check(err)
	check(write_h.SetDirection(pcap.DirectionOut))
	defer write_h.Close()

	check(read_h.SetBPFFilter(*flgBPF))
	for {
		data, _, err := read_h.ReadPacketData()
		check(err)
		check(write_h.WritePacketData(data))
	}

}