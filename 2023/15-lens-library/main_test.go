package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	require.Equal(t, 52, hash("HASH"))
	require.Equal(t, 30, hash("rn=1"))
	require.Equal(t, 253, hash("cm-"))
	require.Equal(t, 0, hash("rn"))
	require.Equal(t, 0, hash("cm"))
	require.Equal(t, 1, hash("qp"))
	require.Equal(t, 3, hash("pc"))

	step := "cm-"
	label := step[:strings.Index(step, "-")]
	require.Equal(t, "cm", label)

	var p [256][]string = [256][]string{}
	steps := []string{"rn=1"}
	for _, step := range steps {
		parts := strings.Split(step, "=") // label=focal_length
		label = parts[0]
		focal := parts[1]
		box := hash(label)
		p[box] = append([]string{label + " " + focal}, p[box]...)
	}
	require.Equal(t, [256][]string{{"rn 1"}}, p)

}

func TestHashSeq(t *testing.T) {
	require.Equal(t, 1320, hashSeq("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
}
