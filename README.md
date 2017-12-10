# scratch [![Go Report Card](https://goreportcard.com/badge/gitfu/scratch)](https://goreportcard.com/report/gitfu/scratch)

short programs on scratch paper

#### ```Install go```
      https://golang.org/doc/install

#### ```Set your Environment```
```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

#### ``` Fetch scratch```
```go
go get github.com/gitfu/scratch
```


```go
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

var Blank = ""
var Manifest string
var UriPrefix string
var SubGroup string
var x264Profiles = map[string]string{"Baseline": "42", "Main": "4d", "High": "64"}
var AudioProfiles = map[string]string{"HE-AACv2": "mp4a.40.5", "LC": "mp4a.40.2"}

type Format struct {
	FormatName string `json:"format_name"`
	BitRate    string `json:"bit_rate"`
}

type Stream struct {
	CodecType string  `json:"codec_type"`
	CodecName string  `json:"codec_name"`
	Profile   string  `json:"profile"`
	Level     float64 `json:"level"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
}

type Container struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

type Stanza struct {
	Bandwidth  string
	Resolution string
	Level      float64
	Profile    string
	AProfile   string
}

// Generic catchall error checking
func chk(err error, mesg string) {
	if err != nil {
		fmt.Printf("%s\n", mesg)
		syscall.Exit(-1)
	}
}

func Probe(segment string) []byte {
	one := "ffprobe -hide_banner  -show_entries format=bit_rate -show_entries "
	two := "stream=codec_type,codec_name,height,width,profile,level -of json -i "
	cmd := fmt.Sprintf("%s%s%s", one, two, segment)
	parts := strings.Fields(cmd)
	data, err := exec.Command(parts[0], parts[1:]...).Output()
	chk(err, fmt.Sprintf("Error running \n %s \n %v", cmd, string(data)))
	return data
}

func fixPrefix(manifest string, uriprefix string) string {
	if uriprefix != Blank {
		if !(strings.HasSuffix(uriprefix, "/")) {
			if !(strings.HasPrefix(manifest, "/")) {
				uriprefix += "/"
			}
		}
	}	
	return uriprefix
}

// create a subtitle stanza for use in the  master.m3u8
func mkSubStanza(manifest string, uriprefix string, subgroup string) string {
	if subgroup ==Blank {
		subgroup="WebVtt"
	}	
	one := fmt.Sprintf("#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID=\"%s\",",subgroup)
	two := "NAME=\"English\",DEFAULT=YES,AUTOSELECT=YES,FORCED=NO,"
	three := fmt.Sprintf("LANGUAGE=\"en\",URI=\"%s%s\"\n", uriprefix, manifest)
	return one + two + three
}

func findSegment(manifest string) string {
	file, err := os.Open(manifest)
	chk(err, "trouble reading manifest")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !(strings.HasPrefix(line, "#")) {
			segment := strings.Replace(manifest, path.Base(manifest), line, 1)
	
			return segment
		}
	}
	return Blank
}

func showStanza(stanza string, mpath string) {
	fmt.Println("")
	fmt.Println(stanza)
	fmt.Println(mpath)
}

func mkStanza(manifest string, segment string, subgroup string, uriprefix string) {
	var st Stanza
	var f Container
	jason := Probe(segment)
	err := json.Unmarshal(jason, &f)
	chk(err, "bad data while probing file")
	st.Bandwidth = f.Format.BitRate
	uriprefix = fixPrefix(manifest, uriprefix)
	for _, i := range f.Streams {
		if i.CodecType == "subtitle" {
			substanza := mkSubStanza(manifest, uriprefix,subgroup)
			showStanza(substanza,Blank)
			return
		}
		if i.CodecType == "video" {
			st.Resolution = fmt.Sprintf("%vx%v", i.Width, i.Height)
			st.Profile = x264Profiles[i.Profile]
			st.Level = i.Level
		}
		if i.CodecType == "audio" {
			st.AProfile = "," + AudioProfiles[i.Profile]
		}

	}
	m3u8Stanza := fmt.Sprintf("#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=%v,RESOLUTION=%s,CODECS=\"avc1.%v00%x%v\"", st.Bandwidth, st.Resolution, st.Profile, int(st.Level), st.AProfile)
	if subgroup != Blank {
		m3u8Stanza = fmt.Sprintf("%s,SUBTITLES=\"%s\"", m3u8Stanza, subgroup)
	}
	mpath := fmt.Sprintf("%s%s\n", uriprefix, manifest)
	showStanza(m3u8Stanza ,mpath)
}

func mkFlags() {
	flag.StringVar(&Manifest, "i", Blank, "manifest file (required)")
	flag.StringVar(&SubGroup, "s", Blank, "add subtitle group i.e. SUBTITLES= (optional)")
	flag.StringVar(&UriPrefix, "u", Blank, "url prefix to add to index.m3u8 path (optional)")
	flag.Parse()
}

func do(manifest string,subgroup string,uriprefix string) {
	segment := findSegment(Manifest)
	mkStanza(manifest, segment, subgroup, uriprefix)
}

func main() {
	mkFlags()
	if Manifest != Blank {
		do(Manifest,SubGroup, UriPrefix)
	}else {
		flag.PrintDefaults()
	}
}

```
