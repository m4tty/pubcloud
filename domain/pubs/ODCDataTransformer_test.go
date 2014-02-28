package pubsDomain

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestODCDataFormatter(t *testing.T) {
	dataTransformer := ODCDataTransformer{}
	csvBytes := []byte("this,is,a,test,so,blah\n1234,1234,2314,someTitle,qwer,asdf")
	reader := bytes.NewReader(csvBytes)
	results, err := dataTransformer.PubDataTransform(reader)
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, "someTitle", results[0].Name)
	assert.Equal(t, 1, len(results))
}
