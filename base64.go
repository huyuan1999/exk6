// @Time    : 2023/11/30 7:40 PM
// @Author  : HuYuan
// @File    : base64.go
// @Email   : hy2803660215@163.com

package exk6

import (
    "encoding/base64"
    "go.k6.io/k6/js/modules"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

type Base64 struct {
    timeout time.Duration
}

func init() {
    modules.Register("k6/x/base64", new(Base64))
}

func (b *Base64) LocalFileEncode(path string) string {
    img, err := os.ReadFile(path)
    if err != nil {
        log.Fatalf("Open %s error: %v", path, err)
    }
    return base64.StdEncoding.EncodeToString(img)
}

func (b *Base64) SetHttpTimeout(timeout time.Duration) {
    b.timeout = timeout
}

func (b *Base64) HttpFileEncode(url string) string {
    if b.timeout == 0 {
        b.timeout = 30
    }

    client := &http.Client{
        Timeout: b.timeout,
    }

    resp, err := client.Get(url)
    if err != nil {
        log.Fatalf("HTTP GET %s error: %v", url, err)
    }
    defer func() { _ = resp.Body.Close() }()

    if resp.StatusCode != http.StatusOK {
        log.Fatalf("HTTP status code exception: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Read response body error: %v", err)
    }
    return base64.StdEncoding.EncodeToString(body)
}
