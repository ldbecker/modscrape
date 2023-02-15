package modscrape

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func NewCookieClient(xurl string, xcookiestr string) (*http.Client, error) {
	fmt.Println("hi")
	cjar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	xcookies := make([]*http.Cookie, 0)

	for _, cx := range strings.Split(xcookiestr, ";") {
		cxs := strings.Split(strings.TrimSpace(cx), "=")
		xcookies = append(xcookies, &http.Cookie{
			Name:  cxs[0],
			Value: cxs[1],
		})
	}
	uuu, err := url.Parse(xurl)
	if err != nil {
		return nil, err
	}
	cjar.SetCookies(uuu, xcookies)
	cookieclient := &http.Client{
		Jar: cjar,
	}
	return cookieclient, nil
}
