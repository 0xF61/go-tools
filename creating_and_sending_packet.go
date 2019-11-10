// Need to work
package main

import (
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers" // Creating the byte structure for the several layers
	"github.com/google/gopacket/pcap"   // as an easy way to send bytes
)

var (
	device      string = "eth0"
	snapshotLen int32  = 1024
	promiscuous        = false
	err         error
	timeout     = -1 * time.Second
	handle      *pcap.Handle
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions
)

func main() {
	// Open Device
	handler, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal("Error opening device.", err)
	}
	defer handler.Close()
	payload := "OUR Pancake! Pancake! Pancake! Pancake!"

	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0x80, 0x00, 0x27, 0x51, 0x1c, 0x5c},
		DstMAC: net.HardwareAddr{0x80, 0x00, 0x27, 0x24, 0xfd, 0x11},
	}

	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{8, 8, 8, 8},
	}

	// TCP Layer struct has Boolean fields for the SYN, FIN and ACK flags
	// Good for manipulating and fuzzing TCP handshakes, sessions and port scan
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(1337),
		DstPort: layers.TCPPort(80),
	}

	// And create the packet with the layers
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(payload),
	)
	outgoingPacket := buffer.Bytes()

	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal("Error sending packet to network device.", err)
	}
}
