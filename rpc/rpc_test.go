package rpc_test

import (
	"lsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecoding(t *testing.T) {
	incommingMessage := "Content-Length: 16\r\n\r\n{\"Method\":\"hi!\"}"
	method, contentLength, err := rpc.DecodeMessage([]byte(incommingMessage))
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 16 {
		t.Fatalf("Expected: 16, Got: %d", contentLength)
	}
	if method != "hi!" {
		t.Fatalf("Expected: \"hi!\", Got: %s", method)
	}
}
