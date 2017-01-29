package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var infile string
var toplevel string

var jasonfile = `./hls.json`

var vcodec = "-c:v libx264 -x264-params no-scenecut=1 "
var acodec = " -c:a aac"
var hls = "-hls_time 2 -hls_list_size 0 -hls_flags round_durations"

type Variant struct {
	Name      string  `json:"name"`
	Aspect    string  `json:"aspect"`
	Framerate float64 `json:"framerate"`
	Vbitrate  string  `json:"vbitrate"`
	Bufsize   string  `json:"bufsize"`
	Maxrate   string  `json:"maxrate"`
	Abitrate  string  `json:"abitrate"`
	Bandwidth int     `json:"bandwidth"`
	Stanza    string  `json:"Stanza"`
}

// This Variant method assembles the ffmpeg command
func (v *Variant) mkCmd() string {
	ffbase := fmt.Sprintf("ffmpeg -i %s -vf scale=%s ", infile, v.Aspect)
	ffvcodec := fmt.Sprintf("%v-g %v -r %v ", vcodec, v.Framerate, v.Framerate)
	ffvrate := fmt.Sprintf(" -b:v %s  -maxrate %s -bufsize %s  ", v.Vbitrate, v.Maxrate, v.Bufsize)
	fftail := fmt.Sprintf(" %s -b:a %s %s %s/%s/index.m3u8", acodec, v.Abitrate, hls, toplevel, v.Name)
	cmd := fmt.Sprintf("%s%s%s%s", ffbase, ffvcodec, ffvrate, fftail)
	return cmd
}

// This Variant method runs the ffmpeg command
func (v *Variant) runCmd(cmd string) {
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1:len(parts)]...).Output()
	chk(err)
	fmt.Printf("%s", out)
}

// Create variant's destination directory
func (v *Variant) mkDest() string {
	dest := fmt.Sprintf("%s/%s", toplevel, v.Name)
	os.MkdirAll(dest, 0755)
	return dest
}

// #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=7483000,RESOLUTION=1920:1080,
// CODECS="avc1.42e00a,mp4a.40.2" hd1920/index.m3u8
func (v *Variant) mkStanza() {
	v.Stanza = fmt.Sprintf(`#EXT-X-STREAM-INF:PROGRAM-ID=1, BANDWIDTH=%v, RESOLUTION=%v, CODECS="avc1.42e00a,mp4a.40.2"`, v.Bandwidth, v.Aspect)
}

// Start transcoding the variant
func (v *Variant) start() {
	dest := v.mkDest()
	fmt.Println("Starting ", dest)
	cmd := v.mkCmd()
	fmt.Println(cmd)
	v.runCmd(cmd)
	v.mkStanza()
}

// Read json file for variants
func dataToVariants() []Variant {
	var variants []Variant
	data, err := ioutil.ReadFile(jasonfile)
	chk(err)
	json.Unmarshal(data, &variants)
	return variants
}

// Generic catchall error checking
func chk(err error) {
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
}

func mkAll(variants []Variant) {
	os.MkdirAll(toplevel, 0755)
	var m3u8Master = fmt.Sprintf("%s/master.m3u8", toplevel)
	fp, err := os.Create(m3u8Master)
	chk(err)
	defer fp.Close()
	w := bufio.NewWriter(fp)
	w.WriteString("#EXTM3U \n")
	for _, v := range variants {
		v.start()
		w.WriteString(fmt.Sprintf("%s \n", v.Stanza))
		w.WriteString(fmt.Sprintf("%s/index.m3u8\n", v.Name))
	}
	w.Flush()
}

func main() {
	variants := dataToVariants()
	flag.StringVar(&infile, "i", "", "Video file to segment")
	flag.StringVar(&toplevel, "d", "", "top level directory for hls files")
	flag.Parse()
	if (infile != "") && (toplevel != "") {
		mkAll(variants)
	} else {
		flag.PrintDefaults()
	}
}
