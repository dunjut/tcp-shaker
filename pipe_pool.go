package tcp

import "sync"

type pipePool struct {
	pool sync.Pool
}

func newPipePool() pipePool {
	return pipePool{sync.Pool{
		New: func() interface{} {
			return make(chan error, 1)
		}},
	}
}

func (p *pipePool) getPipe() chan error {
	return p.pool.Get().(chan error)
}

func (p *pipePool) putBackPipe(pipe chan error) {
	p.cleanPipe(pipe)
	p.pool.Put(pipe)
}

func (p *pipePool) cleanPipe(pipe chan error) {
	select {
	case <-pipe:
	default:
	}
}
