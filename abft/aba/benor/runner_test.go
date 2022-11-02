package benor

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunner_Run(t *testing.T) {
	r := New()
	decided, value := r.Run(map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true})
	require.EqualValues(t, true, decided)
	require.EqualValues(t, value, value)
	fmt.Printf("decided on round %d with value %t\n", r.Nodes[1].Round, value)
}

func TestRunner_different_values(t *testing.T) {
	r := New()
	decided, value := r.Run(map[int]bool{1: false, 2: false, 3: false, 4: true, 5: true})
	require.EqualValues(t, true, decided)
	require.EqualValues(t, value, value)
	fmt.Printf("decided on round %d with value %t\n", r.Nodes[1].Round, value)
}

func TestRunner_different_values2(t *testing.T) {
	r := NewLessThanFMaliciousBroadcaster()
	decided, value := r.Run(map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true})
	require.EqualValues(t, true, decided)
	require.EqualValues(t, value, value)
	fmt.Printf("decided on round %d with value %t\n", r.Nodes[1].Round, value)
}

func TestRunner_malicious_different_values(t *testing.T) {
	r := NewMoreThanFMaliciousBroadcaster()
	decided, value := r.Run(map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true})
	require.EqualValues(t, true, decided)
	require.EqualValues(t, value, value)
	fmt.Printf("decided on round %d with value %t\n", r.Nodes[1].Round, value)
}
