package gofx

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	data := GetFixtures(t, "sample.xml")
	var ofx OFX
	err := xml.Unmarshal(data, &ofx)
	assert.NoError(t, err, "must be able to unmarshal from xml")
	assert.Nil(t, ofx.EMAILMSGSRSV1, "email message response should be nil")
	assert.NotNil(t, ofx.SIGNONMSGSRSV1, "sign on message response should not be nil")
}

