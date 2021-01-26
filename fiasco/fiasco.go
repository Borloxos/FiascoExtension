package fiasco

import (
	"os"
	"os/exec"
)

// Encode Encodes several files matching with the input pattern in fiasco
func Encode(input string, output string, args string) error {
	// Only encode to I-Frames for now, as the default pattern causes crashes while decoding
	var cmd *exec.Cmd
	if args == "" {
		cmd = exec.Command("cfiasco", "-V", "2", "-i", input, "-o", output, "--pattern=I")
	} else {
		cmd = exec.Command("cfiasco", "-V", "2", "-i", input, "-o", output, "--pattern=I", args)
	}

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

// Decode Decodes a fiasco file
func Decode(input string, output string, args string) error {
	var cmd *exec.Cmd
	if args == "" {
		cmd = exec.Command("dfiasco", "-o", output, input)
	} else {
		cmd = exec.Command("dfiasco", "-o", output, input, args)
	}

	// Verbose
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
