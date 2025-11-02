package codec


type SliceWriter[E any] struct {
	Index int  // Next available cell.
	Buf []E
}

func (w SliceWriter[E]) Write(p []E) (n int, err error) {
	if w.Index >= len(w.Buf) {
		return
	}
	n = copy(w.Buf[w.Index:], p)
	w.Index += n
	return
}
