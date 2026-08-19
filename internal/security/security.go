package security

import (
	"os"
)

func IsSafeFile(path string, maxSize int64) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if info.Size() > maxSize {
		return false
	}
	// Detekce binárních souborů: zkusíme otevřít a podívat se na první bajty
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return false
	}
	// Pokud obsahuje nulové bajty, považujeme za binární
	for i := 0; i < n; i++ {
		if buf[i] == 0 {
			return false
		}
	}
	return true
}
