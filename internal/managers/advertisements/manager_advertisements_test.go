package advertisements

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAdvertisementsList(t *testing.T) {
	const search = "Тест & пример"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/advertisements" ||
			r.URL.Query().Get("search") != search || r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("unexpected request: method=%s path=%s search=%q authorization=%q", r.Method, r.URL.Path, r.URL.Query().Get("search"), r.Header.Get("Authorization"))
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"items":[{"id":"1","title":"Тест & пример"}],"total":1,"has_next":true,"next_cursor":"next"}`)
	}))
	defer server.Close()
	t.Setenv("API_URL", server.URL)

	list := GetAdvertisementsList(t, "test-token", search)
	require.Equal(t, 1, list.Total)
	require.True(t, list.HasNext)
	require.Equal(t, "next", list.NextCursor)
	require.Len(t, list.Items, 1)
	require.Equal(t, "1", list.Items[0].ID)
	require.Equal(t, search, list.Items[0].Title)
}
