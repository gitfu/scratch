 
 Use Go to call ffmpeg to create multiple hls variants, directories for variants and a master.m3u8 file 
 
 Usage:
 
  -d string
    	top level directory for hls files
  -i string
    	Video file to segment
     
      go run manifesto.go -d myDir -i video.ts
