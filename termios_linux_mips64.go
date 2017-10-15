// +build linux

package serial

import "syscall"

func getTermios(Iflag uint32, Cflag uint32, Cc [32]uint8, rate  uint32) syscall.Termios {
	return syscall.Termios{
		Iflag:  Iflag,
		Cflag:  Cflag,
		Cc:     Cc,
	}
}