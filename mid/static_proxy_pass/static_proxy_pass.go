package static_proxy_pass

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/wzshiming/cache"
)

// StaticProxyPass is a middleware handler that serves static files in the given
// directory/filesystem. If the file does not exist on the filesystem, it
// passes along to the next middleware in the chain. If you desire "fileserver"
// type behavior where it returns a 404 for unfound files, you should consider
// using http.FileServer from the Go stdlib.
type StaticProxyPass struct {
	// Dir is the directory to serve static files from
	Url *url.URL
	// Prefix is the optional prefix used to serve the static directory content
	Prefix string
	// Cache time
	CacheTime time.Duration
	// Cache
	Cache cache.Cache
}

// NewStatic returns a new instance of Static
func NewStaticProxyPass(u *url.URL) *StaticProxyPass {
	return &StaticProxyPass{
		Url:    u,
		Prefix: "",
		Cache:  cache.NewMemory(),
	}
}

func (s *StaticProxyPass) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method != "GET" && r.Method != "HEAD" {
		next(rw, r)
		return
	}

	file := r.URL.Path

	if s.CacheTime > 0 {
		data := s.Cache.Get(file)
		if d, ok := data.(*bytes.Reader); ok && d != nil {
			http.ServeContent(rw, r, file, time.Now(), d)
			return
		}
	}

	if s.Prefix != "" {
		if !strings.HasPrefix(file, s.Prefix) {
			next(rw, r)
			return
		}
		file = file[len(s.Prefix):]
		if file != "" && file[0] != '/' {
			next(rw, r)
			return
		}
	}

	u, err := s.Url.Parse(path.Join(s.Url.Path, file))
	if err != nil {
		next(rw, r)
		return
	}

	cli := http.Client{}
	req, err := http.NewRequest(r.Method, u.String(), nil)
	if err != nil {
		next(rw, r)
		return
	}

	resp, err := cli.Do(req)
	if err != nil {
		next(rw, r)
		return
	}
	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		next(rw, r)
		return
	}

	data := bytes.NewReader(d)
	if s.CacheTime > 0 {
		s.Cache.Put(file, data, s.CacheTime)
	}

	http.ServeContent(rw, r, file, time.Now(), data)
}
