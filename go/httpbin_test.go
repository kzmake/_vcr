package vcr

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/assert"
)

func TestHttpbin(t *testing.T) {
	t.Run("GET /status/200", func(t *testing.T) {
		r, _ := recorder.New("testdata/fixtures/cassettes/httpbin_status_200")
		defer r.Stop()

		recClient := &http.Client{
			Transport: r,
		}

		res, err := recClient.Get("https://httpbin.org/status/200")
		assert.Nil(t, err)
		assert.Equal(t, 200, res.StatusCode)
	})
}

func TestHttpbin_存在しないAPI(t *testing.T) {
	t.Run("GET /fakestatus/200", func(t *testing.T) {
		r, _ := recorder.New("testdata/fixtures/cassettes/httpbin_fakestatus_200")
		defer r.Stop()

		recClient := &http.Client{
			Transport: r,
		}

		res, err := recClient.Get("https://httpbin.org/fakestatus/200")
		assert.Nil(t, err)
		assert.Equal(t, 200, res.StatusCode)
	})
}
