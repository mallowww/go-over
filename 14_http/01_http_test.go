package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id":1, "name":"okie", "info":"well"}`))
}

func TestMakeHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	// fmt.Println(server.URL)
	// time.Sleep(20 * time.Second)

	want := &Response{
		ID:   1,
		Name: "okie",
		Info: "well",
	}

	t.Run("ok server resp", func(t *testing.T) {
		resp, err := MakeHTTPCall(server.URL)

		// เทียบ struct
		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		// เช็คว่าไม่มี err ด้วย
		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})
}
