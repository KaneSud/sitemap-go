package sitemap_go

import (
	"encoding/xml"
	"time"
)

type SitemapIndex struct {
	XMLName  xml.Name       `xml:"sitemapindex"`
	xmlns    string         `xml:"xmlns,attr"`
	Sitemaps []SitemapEntry `xml:"sitemap"`
}

type SitemapEntry struct {
	Loc     string     `xml:"loc"`
	LastMod *time.Time `xml:"lastmod,omitempty"`
}

func MakeSitemapIndex(entries []SitemapEntry) SitemapIndex {
	return SitemapIndex{
		xmlns:    "http://www.sitemaps.org/schemas/sitemap/0.9",
		Sitemaps: entries,
	}
}

func (si *SitemapIndex) Add(loc string, lastMod time.Time) {
	si.Sitemaps = append(si.Sitemaps, SitemapEntry{
		Loc:     loc,
		LastMod: &lastMod,
	})
}

func (si *SitemapIndex) GenerateXML() (string, error) {
	output, err := xml.MarshalIndent(si, "", "  ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(output), nil
}

func ParseXMLSitemapIndex(content string) (SitemapIndex, error) {
	var out SitemapIndex
	err := xml.Unmarshal([]byte(content), &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	XHTML   string   `xml:"xmlns:xhtml,attr,omitempty"`
	Image   string   `xml:"xmlns:image,attr,omitempty"`
	Video   string   `xml:"xmlns:video,attr,omitempty"`
	URLs    []*URL   `xml:"url"`
}

func MakeUrlSet() URLSet {
	return URLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		XHTML: "http://www.w3.org/1999/xhtml",
	}
}

func (u *URLSet) GenerateXML() (string, error) {
	output, err := xml.MarshalIndent(u, "", "  ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(output), nil
}

func ParseXMLUrlSet(content string) (URLSet, error) {
	var out URLSet
	err := xml.Unmarshal([]byte(content), &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (u *URLSet) Add(url *URL) {
	u.URLs = append(u.URLs, url)
}

type URL struct {
	Loc        string      `xml:"loc"`
	LastMod    *time.Time  `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq  `xml:"changefreq,omitempty"`
	Priority   *float64    `xml:"priority,omitempty"`
	Images     []Image     `xml:"image:image,omitempty"`
	Videos     []Video     `xml:"video:video,omitempty"`
	Alternate  []Alternate `xml:"xhtml:link,omitempty"`
}

type UrlOption func(*URL)

func WithLastMod(t time.Time) UrlOption {
	return func(u *URL) {
		u.LastMod = &t
	}
}

func WithChangeFreq(f ChangeFreq) UrlOption {
	return func(u *URL) {
		u.ChangeFreq = f
	}
}

func WithPriority(p float64) UrlOption {
	return func(u *URL) {
		u.Priority = &p
	}
}

func WithImages(img []Image) UrlOption {
	return func(u *URL) {
		u.Images = append(u.Images, img...)
	}
}

func WithVideosVideos(m []Video) UrlOption {
	return func(u *URL) {
		u.Videos = append(u.Videos, m...)
	}
}

func MakeUrl(loc string, options ...UrlOption) *URL {
	now := time.Now().UTC()
	priority := 0.5
	out := &URL{
		Loc:        loc,
		LastMod:    &now,
		ChangeFreq: ChangeFreqMonthly,
		Priority:   &priority,
	}
	for _, option := range options {
		option(out)
	}
	return out
}

type Image struct {
	Loc     string `xml:"image:loc"`
	Caption string `xml:"image:caption,omitempty"`
	Title   string `xml:"image:title,omitempty"`
}

type Video struct {
	Loc          string   `xml:"video:loc"`
	ThumbnailLoc string   `xml:"video:thumbnail_loc"`
	Title        string   `xml:"video:title"`
	Description  string   `xml:"video:description"`
	ContentLoc   string   `xml:"video:content_loc,omitempty"`
	Duration     int      `xml:"video:duration,omitempty"`
	Category     string   `xml:"video:category,omitempty"`
	Tags         []string `xml:"video:tag,omitempty"`
}

type Alternate struct {
	Rel      string `xml:"rel,attr"`
	HrefLang string `xml:"hreflang,attr"`
	Href     string `xml:"href,attr"`
}

type ChangeFreq string

const (
	ChangeFreqAlways  ChangeFreq = "always"
	ChangeFreqHourly  ChangeFreq = "hourly"
	ChangeFreqDaily   ChangeFreq = "daily"
	ChangeFreqWeekly  ChangeFreq = "weekly"
	ChangeFreqMonthly ChangeFreq = "monthly"
	ChangeFreqYearly  ChangeFreq = "yearly"
	ChangeFreqNever   ChangeFreq = "never"
)
