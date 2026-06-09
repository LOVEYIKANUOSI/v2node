package rate

import (
	"time"

	"github.com/xtls/xray-core/common/buf"
)

type Reader struct {
	reader        buf.Reader
	timeoutReader buf.TimeoutReader
	interrupter   interface{ Interrupt() }
	limiter       *DynamicBucket
}

func NewRateLimitReader(reader buf.Reader, limiter *DynamicBucket) buf.Reader {
	r := &Reader{
		reader:  reader,
		limiter: limiter,
	}
	if timeoutReader, ok := reader.(buf.TimeoutReader); ok {
		r.timeoutReader = timeoutReader
	}
	if interrupter, ok := reader.(interface{ Interrupt() }); ok {
		r.interrupter = interrupter
	}
	return r
}

func (r *Reader) wait(mb buf.MultiBuffer) {
	limiter := r.limiter.Get()
	if limiter != nil {
		limiter.Wait(int64(mb.Len()))
	}
}

func (r *Reader) ReadMultiBuffer() (buf.MultiBuffer, error) {
	mb, err := r.reader.ReadMultiBuffer()
	if !mb.IsEmpty() {
		r.wait(mb)
	}
	return mb, err
}

func (r *Reader) ReadMultiBufferTimeout(timeout time.Duration) (buf.MultiBuffer, error) {
	if r.timeoutReader == nil {
		return r.ReadMultiBuffer()
	}
	mb, err := r.timeoutReader.ReadMultiBufferTimeout(timeout)
	if !mb.IsEmpty() {
		r.wait(mb)
	}
	return mb, err
}

func (r *Reader) Interrupt() {
	if r.interrupter != nil {
		r.interrupter.Interrupt()
	}
}
