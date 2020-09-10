package conf

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const commentChar = "#"

var props map[string]string

var once sync.Once

var DefaultSource Source = source{}

type Source interface {
	Read() (io.ReadCloser, error)
}

type source struct {
}

func (s source) Read() (io.ReadCloser, error) {
	f, err := os.Open(filepath.Join(os.Getenv("CONF_BASEDIR"), "app.conf"))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func load() {
	props = make(map[string]string)

	r, err := DefaultSource.Read()
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 && !strings.HasPrefix(trimmedLine, commentChar) {
			kv := strings.Split(trimmedLine, "=")
			props[strings.TrimSpace(kv[0])] =
				strings.TrimSpace(strings.Split(kv[1], commentChar)[0]) // take value part before comment char
		}
	}
}

func Get(key string, defaultValue ...string) string {
	once.Do(load)

	v, found := props[key]
	if !found || len(v) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return v
}

func GetBool(key string, defaultValue ...bool) bool {
	once.Do(load)

	v, found := props[key]
	if !found {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	b := v == "true" || v == "yes"
	return b
}

func GetInt(key string, defaultValue ...int) int {

	defFunc := func() int {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	once.Do(load)

	v, found := props[key]
	if !found {
		return defFunc()
	}

	d, err := strconv.Atoi(v)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return defFunc()
	}
	return d
}

func GetDuration(key string, defaultValue ...time.Duration) time.Duration {
	once.Do(load)

	defFunc := func() time.Duration {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return time.Duration(0)
	}

	v, found := props[key]
	if !found {
		return defFunc()
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return defFunc()
	}
	return d
}
