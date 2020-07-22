package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"runtime"
	"time"
)

type ParserServer struct {
	UnimplementedParserIOServer
}

func (s *ParserServer) Parse(_ context.Context, input *Input) (*Output, error) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	alloc := mem.TotalAlloc
	start := time.Now()

	data, err := base64.StdEncoding.DecodeString(input.EncodedHTML)
	if err != nil {
		return nil, err
	}

	detector := chardet.NewTextDetector()
	char, err := detector.DetectBest(data)

	if err != nil {
		return nil, err
	}

	r, err := charset.NewReader(bytes.NewReader(data), char.Charset)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	result := &Output{
		Title:     []string{},
		Canonical: []string{},
	}

	doc.Find("title").Each(func(_ int, s *goquery.Selection) {
		result.Title = append(result.Title, s.Text())
	})

	doc.Find("meta[name=title]").Each(func(_ int, s *goquery.Selection) {
		val, exists := s.Attr("content")
		if exists {
			result.Title = append(result.Title, val)
		}
	})

	doc.Find("link[rel=canonical]").Each(func(_ int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists {
			result.Canonical = append(result.Canonical, val)
		}
	})

	runtime.ReadMemStats(&mem)
	duration := time.Since(start)
	var totalAlloc = float64(mem.TotalAlloc-alloc) / 1024 / 1024
	fmt.Printf("duration: %s; memory alloc: %.2f Mb\n", duration.String(), totalAlloc)

	return result, nil
}
