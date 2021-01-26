package main

import (
	"FiascoExtension/ffmpeg"
	"FiascoExtension/fiasco"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"path/filepath"
)

const (
	ActionEncode = "encode"
	ActionDecode = "decode"

	EncodingTempFilename       = "out/frame"
	EncodingTempExtension      = "ppm"
	EncodingTempFiascoWildcard = "[001-%03d+1]"
	EncodingTempFFmpegWildcard = "%03d"
)

func main() {
	// Read program arguments
	parser := argparse.NewParser("FiascoExtension", "Extends the functionality of Fiasco")
	action := parser.Selector("a", "action", []string{ActionEncode, ActionDecode}, &argparse.Options{
		Required: true,
		Help:     "Action to run the coder with. One of: [" + ActionEncode + ", " + ActionDecode + "]",
	})
	input := parser.String("i", "input", &argparse.Options{
		Required: true,
		Help:     "Input file to encode/decode from",
	})
	output := parser.String("o", "output", &argparse.Options{
		Required: true,
		Help:     "Output file to encode/decode to",
	})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	switch *action {
	case ActionEncode:
		// Tile a given videos frames into .ppm files and store the number of produced files
		matches, err := ffmpeg.Encode(*input, EncodingTempFilename+EncodingTempFFmpegWildcard+"."+EncodingTempExtension)
		if err != nil {
			panic(err)
		}

		// Encode tiled files into 1 .fco file
		err = fiasco.Encode(fmt.Sprintf(EncodingTempFilename+EncodingTempFiascoWildcard+"."+EncodingTempExtension, matches), *output)
		cleanupCodingFiles()
	case ActionDecode:
		// Decode .fco compressed file into tiled .ppm files
		err = fiasco.Decode(*input, EncodingTempFilename+"."+EncodingTempExtension)

		// Fiasco puts out files in the format of '[filename without extension].[sequence number].[extension]'
		err = ffmpeg.Decode(fmt.Sprintf("%s.%%*.%s", EncodingTempFilename, EncodingTempExtension), *output)
		cleanupCodingFiles()
	}
}

// cleanupCodingFiles Removes all temporary files that are produced during encoding/decoding
func cleanupCodingFiles() {
	files, err := filepath.Glob(EncodingTempFilename + "*." + EncodingTempExtension)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			fmt.Println(err)
		}
	}
}
