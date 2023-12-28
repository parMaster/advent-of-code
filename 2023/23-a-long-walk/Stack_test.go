package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_StringStack(t *testing.T) {

	s := NewStack("string")

	s.Push("a")
	s.Push("b")
	s.Push("c")

	require.Equal(t, "c", *s.Pop())
	assert.Equal(t, []string{"a", "b"}, s.items)
	assert.Equal(t, "b", *s.Pop())
	assert.Equal(t, "a", *s.Pop())

	s.Push("x")
	s.Push("y")
	s.Push("z")

	assert.Equal(t, "z", *s.Peek())
	assert.Equal(t, "z", *s.Peek())
	assert.Equal(t, "z", *s.Peek())

	assert.Equal(t, "x", *s.PopFirst())
	assert.Equal(t, "z", *s.Peek())
	assert.Equal(t, "y", *s.PopFirst())
	assert.Equal(t, "z", *s.PopFirst())

	assert.Equal(t, true, s.IsEmpty())

	assert.Nil(t, s.PopFirst())
	assert.Nil(t, s.Pop())
	assert.True(t, s.Pop() == nil)

}
