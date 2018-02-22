package gorouter_test

import (
	"fmt"
	"github.com/xujiajun/gorouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"log"
)

var (
	errorFormat, expected string
)

func init() {
	expected = "hi,gorouter"
	errorFormat = "handler returned unexpected body: got %v want %v"
}

func TestRouter_GET(t *testing.T) {
	router := gorouter.New()

	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.GET("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestRouter_POST(t *testing.T) {
	router := gorouter.New()
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.POST("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestRouter_DELETE(t *testing.T) {
	router := gorouter.New()
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodDelete, "/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.DELETE("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestRouter_PATCH(t *testing.T) {
	router := gorouter.New()
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPatch, "/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.PATCH("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestRouter_PUT(t *testing.T) {
	router := gorouter.New()
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPut, "/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.PUT("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestRouter_Group(t *testing.T) {
	router := gorouter.New()

	rr := httptest.NewRecorder()

	prefix := "/api"

	req, err := http.NewRequest(http.MethodGet, prefix+"/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.Group(prefix).GET("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestRouter_HandleNotFound(t *testing.T) {
	router := gorouter.New()

	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/xxx", nil)

	if err != nil {
		t.Fatal(err)
	}

	customNotFoundStr := "404 page !"
	router.NotFoundFunc(func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, customNotFoundStr)
	})

	router.GET("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != customNotFoundStr {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}

func TestGetParam(t *testing.T) {
	router := gorouter.New()

	rr := httptest.NewRecorder()

	param := "1"
	req, err := http.NewRequest(http.MethodGet, "/test/"+param, nil)

	if err != nil {
		t.Fatal(err)
	}

	router.GET("/test/:id", func(w http.ResponseWriter, r *http.Request) {
		id := gorouter.GetParam(r, "id")
		if id != param {
			t.Fatal("TestGetParam test fail")
		}
	})
	router.ServeHTTP(rr, req)
}

func TestGetAllParams(t *testing.T) {
	router := gorouter.New()

	rr := httptest.NewRecorder()

	param1 := "1"
	param2 := "2"
	req, err := http.NewRequest(http.MethodGet, "/param1/"+param1+"/param2/"+param2, nil)

	if err != nil {
		t.Fatal(err)
	}

	router.GET("/param1/:param1/param2/:param2", func(w http.ResponseWriter, r *http.Request) {
		id1 := gorouter.GetParam(r, "param1")
		if id1 != param1 {
			t.Fatal("TestGetAllParams test fail")
		}

		id2 := gorouter.GetParam(r, "param2")
		if id2 != param2 {
			t.Fatal("TestGetAllParams test fail")
		}
	})
	router.ServeHTTP(rr, req)
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func TestRouter_Use(t *testing.T) {
	router := gorouter.New()

	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/hi", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.Use(withLogging)
	router.GET("/hi", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, expected)
	})
	router.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf(errorFormat,
			rr.Body.String(), expected)
	}
}
