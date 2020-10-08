package algo

type RingBuffer struct {
	inputChan  <-chan int
	outputChan chan int
}

func NewRingBuffer(in <-chan int, out chan int) *RingBuffer {
	return &RingBuffer{
		inputChan:  in,
		outputChan: out,
	}
}

func (r *RingBuffer) Run() {
	for v := range r.inputChan {
		select {
		case r.outputChan <- v:
		default:
			<-r.outputChan
			r.outputChan <- v
		}
	}
	close(r.outputChan)
}
