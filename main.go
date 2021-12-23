package main

import (
	"fmt"
	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw"
	"log"
	"net"
)

func main() {
	ifi, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatalf("failed to find interface %q: %v", "eth0", err)
	}
	c, err := raw.ListenPacket(ifi, etherType, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer c.Close()

	bytestream := []byte{255, 255, 255, 255, 255, 255, 228, 21, 246, 244, 49, 166, 204, 204}
	fmt.Println(bytestream)

	addr := &raw.Addr{
		HardwareAddr: ethernet.Broadcast,
	}

	if _, err := c.WriteTo(bytestream, addr); err != nil {
		log.Fatalf("failed to write frame: %v", err)
	}
}
