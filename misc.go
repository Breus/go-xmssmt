package xmssmt

import (
	"encoding/binary"
	"fmt"
)

// Encodes the given uint64 into the buffer out in Big Endian
func encodeUint64Into(x uint64, out []byte) {
	if len(out)%8 == 0 {
		binary.BigEndian.PutUint64(out[len(out)-8:], x)
		for i := 0; i < len(out)-8; i += 8 {
			binary.BigEndian.PutUint64(out[i:i+8], 0)
		}
	} else {
		for i := len(out) - 1; i >= 0; i-- {
			out[i] = byte(x)
			x >>= 8
		}
	}
}

// Encodes the given uint64 as [outLen]byte in Big Endian.
func encodeUint64(x uint64, outLen int) []byte {
	ret := make([]byte, outLen)
	encodeUint64Into(x, ret)
	return ret
}

// Interpret []byte as Big Endian int.
func decodeUint64(in []byte) (ret uint64) {
	// TODO should we use binary.BigEndian?
	for i := 0; i < len(in); i++ {
		ret |= uint64(in[i]) << uint64(8*(len(in)-1-i))
	}
	return
}

type errorImpl struct {
	msg    string
	locked bool
	inner  error
}

func (err *errorImpl) Locked() bool { return err.locked }
func (err *errorImpl) Inner() error { return err.inner }

func (err *errorImpl) Error() string {
	if err.inner != nil {
		return fmt.Sprintf("%s: %s", err.msg, err.inner.Error())
	}
	return err.msg
}

// Formats a new Error
func errorf(format string, a ...interface{}) *errorImpl {
	return &errorImpl{msg: fmt.Sprintf(format, a...)}
}

// Formats a new Error that wraps another
func wrapErrorf(err error, format string, a ...interface{}) *errorImpl {
	return &errorImpl{msg: fmt.Sprintf(format, a...), inner: err}
}

type dummyLogger struct{}

func (logger *dummyLogger) Logf(format string, a ...interface{}) {}

var log Logger

type Logger interface {
	Logf(format string, a ...interface{})
}

// Enables logging
func SetLogger(logger Logger) {
	log = logger
}