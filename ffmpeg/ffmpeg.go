package ffmpeg

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Encode Splits a video file into a number of tiled .ppm image files containing the frames
func Encode(input string, output string) (matches int, err error) {
	// TODO: save tiling layout and framerate in some way
	cmd := exec.Command("ffmpeg", "-i", input, "-y", "-vf", "tile=layout=4x1", output)

	// Verbosity
	cmd.Stderr = os.Stdout

	err = cmd.Start()
	if err != nil {
		return 0, err
	}

	err = cmd.Wait()
	if err != nil {
		return 0, err
	}

	// Count the number of files produced
	matchesSlice, err := filepath.Glob("out/frame*.ppm")
	if err != nil {
		return 0, err
	}

	return len(matchesSlice), err
}

// Decode Combines a number of tiled .ppm images containing frames into a video file
func Decode(input string, output string) error {
	// TODO: read tiling layout and target framerate from input files
	// if '-f image2' isn't specified before the input file, ffmpeg fails to use wildcards correctly
	// start_number is set to 0, because Fiasco starts its output files at 0
	cmd := exec.Command("ffmpeg", "-f", "image2", "-i", input, "-y", "-vf", "untile=4x1,setpts=N/(25*TB)", "-start_number", "0", output)

	// Verbosity
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}