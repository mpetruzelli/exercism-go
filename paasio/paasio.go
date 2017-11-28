package paasio

import (
	"io"
	"sync"
)

type Counter struct {
	n     int64
	nops  int
	rf    func(p []byte) (n int, err error)
	wf    func(p []byte) (n int, err error)
	mutex sync.Mutex
}

func NewReadCounter(r io.Reader) ReadCounter {
	var rc ReadCounter

	rc = &Counter{n: 0, nops: 0, rf: r.Read, wf: nil}

	return rc
}

// NewReadCounter returns a new instance of WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {

	var wc WriteCounter

	wc = &Counter{n: 0, nops: 0, rf: nil, wf: w.Write}

	return wc

}

// NewReadCounter returns a new instance of ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {

	var rwc ReadWriteCounter

	rwc = &Counter{n: 0, nops: 0, rf: rw.Read, wf: rw.Write}

	return rwc
}

// ReadCount returns the number of read bytes and the number of operations
func (r *Counter) ReadCount() (n int64, nops int) {
	return r.n, r.nops
}

// WriteCount returns the number of written bytes and the number of operations
func (w *Counter) WriteCount() (n int64, nops int) {
	return w.n, w.nops
}

// Read performs a read operation and counts the read bytes and the number of operations
func (r *Counter) Read(p []byte) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n, err := r.rf(p)
	n64 := int64(n)
	if err == nil {
		r.n += int64(n64)
		r.nops += 1
	}
	return n, err
}

// Read performs a write operation and counts the written bytes and the number of operations
func (w *Counter) Write(p []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	n, err := w.wf(p)
	n64 := int64(n)
	if err == nil {
		w.n += int64(n64)
		w.nops += 1
	}
	return n, err
}
