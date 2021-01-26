package main

import (
	"FiascoExtension/ffmpeg"
	"FiascoExtension/fiasco"
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"path/filepath"
	"regexp"
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
		Help:     "Action to run the coder with. One of: [" + ActionEncode + ", " + ActionDecode + "].",
	})
	input := parser.String("i", "input", &argparse.Options{
		Required: true,
		Help:     "Input file to encode/decode from.",
	})
	output := parser.String("o", "output", &argparse.Options{
		Required: true,
		Help:     "Output file to encode/decode to.",
	})
	layout := parser.String("l", "layout", &argparse.Options{
		Required: false,
		Validate: validateLayout,
		Help: "Layout to tile the picture groups in. Specified in the format '4x1', where the first number is the " +
			"width and the second number is the height of the tiling.",
		Default: "4x1",
	})
	fps := parser.Int("f", "fps", &argparse.Options{
		Required: false,
		Help:     "Target fps for the decoded video.",
		Default:  25,
	})
	ffmpegArgs := parser.String("", "ffmpegArgs", &argparse.Options{
		Required: false,
		Help:     "Additional arguments to append to the ffmpeg command.",
		Default:  "",
	})
	fiascoArgs := parser.String("", "fiascoArgs", &argparse.Options{
		Required: false,
		Help:     "Additional arguments to append to the fiasco command.",
		Default:  "",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	switch *action {
	case ActionEncode:
		// Tile a given videos frames into .ppm files and store the number of produced files
		matches, err := ffmpeg.Encode(*input, EncodingTempFilename+EncodingTempFFmpegWildcard+"."+EncodingTempExtension, *layout, *ffmpegArgs)
		if err != nil {
			panic(err)
		}

		// Encode tiled files into 1 .fco file
		err = fiasco.Encode(fmt.Sprintf(EncodingTempFilename+EncodingTempFiascoWildcard+"."+EncodingTempExtension, matches), *output, *fiascoArgs)
		cleanupCodingFiles()
	case ActionDecode:
		// Decode .fco compressed file into tiled .ppm files
		err = fiasco.Decode(*input, EncodingTempFilename+"."+EncodingTempExtension, *fiascoArgs)

		// Fiasco puts out files in the format of '[filename without extension].[sequence number].[extension]'
		err = ffmpeg.Decode(fmt.Sprintf("%s.%%*.%s", EncodingTempFilename, EncodingTempExtension), *output, *layout, *fps, *ffmpegArgs)
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

func validateLayout(args []string) error {
	if len(args) <= 0 {
		return errors.New("no layout parameter specified")
	}
	if len(args) > 1 {
		return errors.New("too many layout parameters specified")
	}

	matched, err := regexp.Match(`^\d+x\d+$`, []byte(args[0]))
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("invalid layout specified. Use the format '4x1', where the first number is the width " +
			"and the second number is the height of the tiling")
	}

	return nil
}