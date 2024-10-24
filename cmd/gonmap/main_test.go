package main

import (
	"encoding/xml"
	"testing"

	"github.com/Ullaakut/nmap/v3"
	"github.com/stretchr/testify/require"
)

func TestConvertTestdata(t *testing.T) {
	var in *nmap.Run
	err := xml.Unmarshal([]byte(example), &in)
	require.NoError(t, err)
	require.Equal(t, "nmap", in.Scanner, "proper input data")
	got, err := convert(in)
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestConvertNil(t *testing.T) {
	got, err := convert(nil)
	require.NoError(t, err)
	require.Empty(t, got)
}
