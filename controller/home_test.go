package controller

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginTemplate(t *testing.T) {
	h := new(home)
	expected := "login template"
	h.loginTemplate, _ = template.New("").Parse(expected)

	r := httptest.NewRequest(http.MethodGet, "/loginURL", nil)
	w := httptest.NewRecorder()

	h.loginHandler(w, r)
	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed to execute login template")
	}
}
