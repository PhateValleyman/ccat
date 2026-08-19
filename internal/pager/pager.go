package pager

import (
	"os"
	"os/exec"
)

type Pager struct {
	Stdin  *os.File
	cmd    *exec.Cmd
}

func StartPager() (*Pager, error) {
	less := exec.Command("less", "-R")
	less.Stdout = os.Stdout
	less.Stderr = os.Stderr
	in, err := less.StdinPipe()
	if err != nil {
		return nil, err
	}
	if err := less.Start(); err != nil {
		return nil, err
	}
	// Přesměrujeme standardní výstup do pipe
	return &Pager{Stdin: in.(*os.File), cmd: less}, nil
}

func (p *Pager) Close() error {
	if err := p.Stdin.Close(); err != nil {
		return err
	}
	return p.cmd.Wait()
}
