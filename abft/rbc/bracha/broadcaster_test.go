package bracha

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBroadcaster_Broadcast(t *testing.T) {
	b := NewBroadcaster()
	v := b.Broadcast(1, []byte{1, 2, 3, 4})
	require.EqualValues(t, []byte{1, 2, 3, 4}, v)
}

func TestBroadcaster_FMaliciousBroadcast(t *testing.T) {
	b := NewLessThanFMaliciousBroadcaster()
	v := b.Broadcast(1, []byte{1, 2, 3, 4})
	require.EqualValues(t, []byte{1, 2, 3, 4}, v)
}

func TestBroadcaster_MoreThanFMaliciousBroadcast(t *testing.T) {
	b := NewMoreThanFMaliciousBroadcaster()
	v := b.Broadcast(1, []byte{1, 2, 3, 4})
	require.Nil(t, v)
}
