package testutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func fixture(path string) string {
	filePath := filepath.Join(
		RootDir(),
		"/testdata/fixtures/",
		path,
	)
	b, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		panic(err)
	}
	return string(b)
}

func SetupMockAPI() (string, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/abuse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("abuses.json"))
	})

	mux.HandleFunc("/dedibox/plans", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("plans.json"))
	})

	mux.HandleFunc("/network/ddos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("ddos.json"))
	})

	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("servers.json"))
	})

	mux.HandleFunc("/server/backup/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.String(), "/") // /server/backup/5678
		sid, err := strconv.Atoi(s[len(s)-1])
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture(fmt.Sprintf("backup-%s.json", strconv.Itoa(sid))))
	})

	return server.URL, func() {
		server.Close()
	}
}
