package main

import "flag"

func main() {
	flags()
}

func flags() {
	//flag.IntVar(&player.NumFfmpegThreads, "ffthreads", 1, "Number of threads for FFMPEG Conversions")
	//flag.BoolVar(&player.DisableFfmpeg, "ffdisable", false, "Disable FFMPEG Conversions")
	flag.Parse()
}
