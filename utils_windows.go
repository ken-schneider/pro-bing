//go:build windows
// +build windows

package probing

import (
	"math"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

const (
	minimumBufferLength = 2048
)

// Returns the length of an ICMP message, plus the IP packet header.
// Calculated as:
// len(response ICMP header) + len(request IP header)
// + len(request ICMP header) + len(request ICMP data)
func (p *Pinger) getMessageLength() int {
	if p.ipv4 {
		calculatedLength := 8 + ipv4.HeaderLen + 8 + p.Size
		return int(math.Max(float64(calculatedLength), float64(minimumBufferLength)))
	}
	calculatedLength := 8 + ipv6.HeaderLen + 8 + p.Size
	return int(math.Max(float64(calculatedLength), float64(minimumBufferLength)))
}

// Attempts to match the ID of an ICMP packet.
func (p *Pinger) matchID(ID int) bool {
	if ID != p.id {
		return false
	}
	return true
}

// SetMark sets the SO_MARK socket option on outgoing ICMP packets.
// Setting this option requires CAP_NET_ADMIN.
func (c *icmpConn) SetMark(mark uint) error {
	return ErrMarkNotSupported
}

// SetMark sets the SO_MARK socket option on outgoing ICMP packets.
// Setting this option requires CAP_NET_ADMIN.
func (c *icmpv4Conn) SetMark(mark uint) error {
	return ErrMarkNotSupported
}

// SetMark sets the SO_MARK socket option on outgoing ICMP packets.
// Setting this option requires CAP_NET_ADMIN.
func (c *icmpV6Conn) SetMark(mark uint) error {
	return ErrMarkNotSupported
}

// SetDoNotFragment sets the do-not-fragment bit in the IP header of outgoing ICMP packets.
func (c *icmpConn) SetDoNotFragment() error {
	return ErrDFNotSupported
}

// SetDoNotFragment sets the do-not-fragment bit in the IP header of outgoing ICMP packets.
func (c *icmpv4Conn) SetDoNotFragment() error {
	return ErrDFNotSupported
}

// SetDoNotFragment sets the do-not-fragment bit in the IPv6 header of outgoing ICMPv6 packets.
func (c *icmpV6Conn) SetDoNotFragment() error {
	return ErrDFNotSupported
}

// No need for SetBroadcastFlag in non-linux OSes
func (c *icmpConn) SetBroadcastFlag() error {
	return nil
}

func (c *icmpv4Conn) SetBroadcastFlag() error {
	return nil
}

func (c *icmpV6Conn) SetBroadcastFlag() error {
	return nil
}
