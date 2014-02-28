package pubsDomain

import (
	"bytes"
	"github.com/m4tty/pubcloud/web/resources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"
)

type MockedTransformer struct {
	mock.Mock
}
type MockedRawGetter struct {
	mock.Mock
}

func (mockedTransformer MockedTransformer) PubDataTransform(r interface{}) ([]resources.Pubs, error) {
	args := mockedTransformer.Mock.Called(r)
	return args.Get(0).([]resources.Pubs), args.Error(1)
}

func (mockedRawGetter MockedRawGetter) RawDataGet(path string) (io.ReadCloser, error) {
	args := mockedRawGetter.Mock.Called(path)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func TestODCDataRetriever(t *testing.T) {
	mockTransformer := new(MockedTransformer)
	mockedRawGetter := new(MockedRawGetter)
	//TODO: set up expectations
	//reader := bufio.NewReader(nil)
	var path = "http://data.denvergov.org/download/gis/liquor_licenses/csv/liquor_licenses.csv"
	var reader = fakeReader()
	mockedRawGetter.On("RawDataGet", path).Return(reader, nil)
	mockTransformer.On("PubDataTransform", reader).Return([]resources.Pubs{}, nil)

	dataRetriever := NewODCPubDataRetriever(mockedRawGetter, mockTransformer)
	results, err := dataRetriever.GetPubData()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	// assert.Equal(t, "someTitle", results[0].Name)
	// assert.Equal(t, 1, len(results))
	mockTransformer.Mock.AssertExpectations(t)

}

type BytesReadCloser struct {
	*bytes.Reader
}

func (brc *BytesReadCloser) Close() (err error) {
	return
}

func fakeReader() io.ReadCloser {
	csvBytes := []byte("this,is,a,test,so,blah\n1234,1234,2314,someTitle,qwer,asdf")
	reader := bytes.NewReader(csvBytes)
	brc := &BytesReadCloser{reader}
	return brc
}

//TODO: need to add a interface... something like GetData() io.Reader, so I can control the mock data.  Pass this into Retriever.
