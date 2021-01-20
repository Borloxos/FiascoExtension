package main

import (
	"FiascoExtension/ffmpeg"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

const (
	ActionEncode = "encode"
	ActionDecode = "decode"
)

func main() {
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
		err = ffmpeg.Encode(*input, *output)
	case ActionDecode:
		err = ffmpeg.Decode(*input, *output)
	}
}
