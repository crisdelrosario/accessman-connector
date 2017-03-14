package accessman

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

// Hash represents a Hashable string
type Hash struct {
	s string
}

// MD5 calculates the MD5 hash of a string
func (h Hash) MD5() string {
	m := md5.New()
	io.WriteString(m, h.s)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// SHA1 calculates the SHA1 hash of a string
func (h Hash) SHA1() string {
	m := sha1.New()
	io.WriteString(m, h.s)
	return fmt.Sprintf("%x", m.Sum(nil))
}
