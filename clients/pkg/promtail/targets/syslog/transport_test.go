package syslog

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIdleTimeoutConnPipe(t *testing.T) {
	addrs, _ := net.InterfaceAddrs()
	timeout := 100 * time.Millisecond

	p := NewIdleTimeoutConnPipe(addrs[0], timeout)
	// upon creation, the idle timeout is set
	require.False(t, p.IsIdle(time.Now()))
	time.Sleep(50 * time.Millisecond)
	require.False(t, p.IsIdle(time.Now()))

	// When reading or writing, the deadline is extended
	go func() {
		buf := make([]byte, 0, 1024)
		p.Read(buf)
	}()
	p.Write([]byte{})
	time.Sleep(50 * time.Millisecond)
	require.False(t, p.IsIdle(time.Now()))
	time.Sleep(80 * time.Millisecond)
	require.True(t, p.IsIdle(time.Now()))
}
