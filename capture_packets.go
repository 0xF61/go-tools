// tcpdump -i eth0 -w captured_packets.pcap
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap" // wrapped around libcap
	"github.com/google/gopacket/pcapgo"
)

var (
	deviceName        = "eth0"
	snapshotLen int32 = 1024
	promiscuous       = false
	err         error
	timeout     = -1 * time.Second
	handle      *pcap.Handle
	packetCount = 0
)

func main() {
	// Open output pcap file and write header
	f, _ := os.Create("network_traffic.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// Open device for capturing
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		if packetCount > 100 {
			break
		}
	}
}
