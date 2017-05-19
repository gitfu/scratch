package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const sitemap string = "http://golfchannel.com/sitemap_video.xml"

// UrlSet represents <urlset>
type UrlSet struct {
	VUrlList []Vurl `xml:"url"` //see Vurl struct
}
//method to show the details for all of the videos
func (u *UrlSet) showAll() {
	for i, vurl := range u.VUrlList {
		v := vurl.Video
		v.show(i)
	}
}

// Vurl represents  a single <url> element,
// you can pick and choose which child elements to include
type Vurl struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
	Video   Vid    `xml:"video"` // see Vid struct
}

// Vid represents a single <video> elememnt
// Title , Content, and Desc are child node values
type Vid struct {
	Title   string `xml:"title"`
	Content string `xml:"content_loc"`
	Desc    string `xml:"description"`
}

// Vid Struct method to print a Vid struct's values
func (v *Vid) show(i int) {
	fmt.Printf("video # %d", i)
	fmt.Printf("\nTitle:\t%s\nContent:\t%s\nDesc:\t%s\n\n",
		v.Title, v.Content, v.Desc)
}

// function to check errors
func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var u UrlSet
	res, err := http.Get(sitemap)
	chk(err)
	defer res.Body.Close()
	xmlBytes, err := ioutil.ReadAll(res.Body)
	chk(err)
	// decode the xml into u, a UrlSet struct
	xml.Unmarshal(xmlBytes, &u)
	u.showAll()

}
