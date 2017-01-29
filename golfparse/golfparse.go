package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 35 MB  xml file, so big it will choke your browser
const sitemap string = "http://golfchannel.com/sitemap_video.xml"

type UrlSet struct {
	// This represents <urlset>
	VUrlList []Vurl `xml:"url"` //see Vurl struct
}

type Vurl struct {
	// This represents  a single <url> element,
	// you can pick and choose which child elements to include
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
	Video   Vid    `xml:"video"` // see Vid struct
}

type Vid struct {
	// this represents a single <video> elememnt
	// Title , Content, and Desc are child node values
	Title   string `xml:"title"`
	Content string `xml:"content_loc"`
	Desc    string `xml:"description"`
}

func (v Vid) show() {
	// Vid Struct method to print a Vid struct's values
	fmt.Printf("\nTitle:\t%s\nContent:\t%s\nDesc:\t%s\n\n",
		v.Title, v.Content, v.Desc)
}

func chk(err error) {
	// function to check errors
	if err != nil {
		panic(err)
	}
}

func main() {
	res, err := http.Get(sitemap)
	chk(err)
	// keep file open so we can read it
	defer res.Body.Close()
	xmlFile, err := ioutil.ReadAll(res.Body)
	chk(err)
	var u UrlSet
	// decode the xml into u, a UrlSet struct
	xml.Unmarshal(xmlFile, &u)
	var count = 0
	for _, vurl := range u.VUrlList {
		v := vurl.Video
		v.show()
		fmt.Printf("video # %d", count)
		count++
	}
}
