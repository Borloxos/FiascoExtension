package fiasco

import (
	"os"
	"os/exec"
)

// Encode Encodes several files matching with the input pattern in fiasco
func Encode(input string, output string) error {
	// Only encode to I-Frames for now, as the default pattern causes crashes while decoding
	cmd := exec.Command("cfiasco", "-V", "2", "-i", input, "-o", output, "--pattern=I")

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

// Decode Decodes a fiasco file
func Decode(input string, output string) error {
	cmd := exec.Command("dfiasco", "-o", output, input)

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
