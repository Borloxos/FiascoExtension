package fiasco

import (
	"os"
	"os/exec"
)

// Encode Encodes several files matching with the input pattern in fiasco
func Encode(input string, output string, path string, customArgs string) error {
	// Only encode to I-Frames for now, as the default pattern causes crashes while decoding
	args := []string{"-V", "2", "-i", input, "-o", output, "--pattern=I"}
	if customArgs != "" {
		args = append([]string{customArgs}, args...)
	}

	var cmd *exec.Cmd
	cmd = exec.Command(path, args...)

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

// Decode Decodes a fiasco file
func Decode(input string, output string, path string, customArgs string) error {
	args := []string{"-o", output, input}
	if customArgs != "" {
		args = append([]string{customArgs}, args...)
	}

	var cmd *exec.Cmd
	cmd = exec.Command(path, args...)

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
