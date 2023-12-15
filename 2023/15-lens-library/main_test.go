package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	require.Equal(t, 52, hash("HASH"))
	require.Equal(t, 30, hash("rn=1"))
	require.Equal(t, 253, hash("cm-"))
}

func TestHashSeq(t *testing.T) {
	require.Equal(t, 1320, hashSeq("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
}
