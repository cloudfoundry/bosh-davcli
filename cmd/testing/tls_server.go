package testing

import (
	"bytes"
	"encoding/pem"
	"net/http/httptest"
)

func ExtractRootCa(server *httptest.Server) (string, error) {
	rootCa := new(bytes.Buffer)
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: server.Certificate().Raw,
	}

	err := pem.Encode(rootCa, block)
	if err != nil {
		return "", err
	}

	return rootCa.String(), nil
}
