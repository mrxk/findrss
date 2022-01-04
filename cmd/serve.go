package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/spf13/cobra"
)

func Serve(cmd *cobra.Command, args []string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/*", http.FileServer(http.Dir(".")))
	r.Post("/find", lookupRssFeed)
	return http.ListenAndServe(":80", r)
}

func lookupRssFeed(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		_, _ = w.Write(formatError(err))
		return
	}
	site := string(data)
	parsedURL, err := url.Parse(site)
	if err != nil {
		_, _ = w.Write(formatError(err))
		return
	}
	collector := colly.NewCollector(colly.AllowedDomains(parsedURL.Host))
	collector.SetDebugger(&debug.LogDebugger{})
	linkCount := 0
	collector.OnHTML("link", func(e *colly.HTMLElement) {
		linkCount += 1
		w.Write(formatLink(parsedURL, e))
	})
	w.Write([]byte("<ul>"))
	err = collector.Visit(site)
	if linkCount == 0 {
		w.Write([]byte("No feeds detected"))
	}
	w.Write([]byte("</ul>"))
	if err != nil {
		_, _ = w.Write(formatError(err))
		return
	}
}

func formatLink(root *url.URL, e *colly.HTMLElement) []byte {
	title := e.Attr("title")
	href := e.Attr("href")
	if href == "" {
		return nil
	}
	rel := e.Attr("rel")
	if title == "" {
		title = rel
	} else {
		title += " (" + rel + ")"
	}
	parsedURL, err := url.Parse(href)
	if err != nil {
		return formatError(err)
	}
	if parsedURL.Host == "" {
		parsedURL.Host = root.Host
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = root.Scheme
	}
	href = parsedURL.String()
	return []byte(fmt.Sprintf("<li>%s: <a href=\"%s\">%s</a></li>", title, href, href))
}

func formatError(err error) []byte {
	return []byte("<pre>" + err.Error() + "</pre>")
}
