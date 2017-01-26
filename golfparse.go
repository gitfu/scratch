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

func (v Vid) toXMLString() {
	// Vid struct method to print struct values as xml 
	bytes, err := xml.Marshal(v)
	chk(err)
	fmt.Printf("\n\n%s\n", string(bytes))
}

func (v Vid) show() {
	// Vid Struct method to print a Vid struct's values
	fmt.Printf("\ntitle:\t%s\nContent location:\t%s\nDescription:\t%s\n\n", v.Title, v.Content, v.Desc)
}

func chk(err error) {
	// function to check errors
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
}

func main() {
	 // http get the xml file
	res, err := http.Get(sitemap)
	// check for errors
	chk(err)
	// keep file open so we can read it
	defer res.Body.Close()
	// read in the xml
	xmlFile, err := ioutil.ReadAll(res.Body)
	// check for errors
	chk(err)
	// declare a var of type UrlSet
	var u UrlSet
	// decode the xml into u, a UrlSet struct
	xml.Unmarshal(xmlFile, &u)
	// so we know how many videos we process
	var count = 0
	// go only does for loops
	// u is a struct, that has an array of vurl structs
	for _, vurl := range u.VUrlList {
		// for each vurl struct in the array
		// v is the vurl's Video struct
		v := vurl.Video
		// call the  Video struct's methods 
		v.toXMLString()
		v.show()
		//print the count
		fmt.Printf(" video # %d", count)
		// increment count
		count++
	}
}
