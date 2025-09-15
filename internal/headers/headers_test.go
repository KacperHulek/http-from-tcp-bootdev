package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	host, _ := headers.Get("Host")
	assert.Equal(t, "localhost:42069", host)
	assert.Equal(t, 25, n)
	assert.True(t, done)

	// Test: valid multiple headers with the same name
	headers = NewHeaders()
	data = []byte("HoSt: localhost:42069\r\nHoST: localhost:42069\r\n")
	_, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	hostRandomCap, _ := headers.Get("hOsT")
	assert.Equal(t, "localhost:42069,localhost:42069", hostRandomCap)
	assert.False(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("H©st: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

}
