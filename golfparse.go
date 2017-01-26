package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var SITEMAP = "http://golfchannel.com/sitemap_video.xml"

type UrlSet struct {
	VUrlList []Vurl `xml:"url"`
}

type Vurl struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
	Video   Vid    `xml:"video"`
}

type Vid struct {
	Title   string `xml:"title"`
	Content string `xml:"content_loc"`
	Desc    string `xml:"description"`
}

func (v Vid) toXMLString() {
	bytes, err := xml.Marshal(v)
	chk(err)
	fmt.Printf("\n\n%s\n", string(bytes))
}

func (v Vid) show() {
	fmt.Printf("\ntitle:\t%s\nContent location:\t%s\nDescription:\t%s\n\n", v.Title, v.Content, v.Desc)
}

func chk(err error) {
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
}

func main() {
	res, err := http.Get(SITEMAP)
	chk(err)
	defer res.Body.Close()
	xmlFile, err := ioutil.ReadAll(res.Body)
	chk(err)
	var u UrlSet
	xml.Unmarshal(xmlFile, &u)
	var count = 0
	for _, vurl := range u.VUrlList {
		v := vurl.Video
		v.toXMLString()
		v.show()
		fmt.Printf(" video # %d", count)
		count++
	}
}
