package handles

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteToHomeWithSuccess(t *testing.T) {
	r := Router()
	ts := httptest.NewServer(r)

	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for / is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestRouteToGetArticleSuccess(t *testing.T) {
	r := Router()
	ts := httptest.NewServer(r)

	defer ts.Close()

	res, err := http.Get(ts.URL + "/article")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /article is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestRouteToPostArticleSuccess(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"title":   "",
		"desc":    "",
		"content": "",
	})

	r := Router()
	ts := httptest.NewServer(r)

	defer ts.Close()

	res, err := http.Post(ts.URL+"/article", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /article is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestRouteToPostArticleFail(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"xxxx":    "",
		"desc":    "",
		"content": "",
	})

	r := Router()
	ts := httptest.NewServer(r)

	defer ts.Close()

	res, err := http.Post(ts.URL+"/article", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code for /article is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}
