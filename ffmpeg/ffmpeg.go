package ffmpeg

import (
	"os"
	"os/exec"
)

func Encode(input string, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-y", "-vf", "tile=layout=4x1", output)

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func Decode(input string, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-y", "-vf", "untile=4x1,setpts=N/(25*TB)", output)

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}