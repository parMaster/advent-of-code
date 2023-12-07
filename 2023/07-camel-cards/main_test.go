package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Hands(t *testing.T) {

	hh := hands{{[]rune("AAAAQ"), 400}, {[]rune("KAAAA"), 92}}

	require.False(t, hh.Less(0, 1))

}
