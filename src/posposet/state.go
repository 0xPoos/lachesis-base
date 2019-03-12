package posposet

import (
	"github.com/Fantom-foundation/go-lachesis/src/common"
)

// TODO: make State internal

// State is a current poset state.
type State struct {
	LastFinishedFrameN uint64
	LastBlockN         uint64
	Genesis            common.Hash
	TotalCap           uint64
}

/*
 * Poset's methods:
 */

// State saves current state.
func (p *Poset) saveState() {
	p.store.SetState(p.state)
}

// bootstrap restores current state from store.
func (p *Poset) bootstrap() {
	// restore state
	p.state = p.store.GetState()
	if p.state == nil {
		panic("Apply genesis for store first")
	}
	// restore frames
	for n := p.state.LastFinishedFrameN; true; n++ {
		if f := p.store.GetFrame(n); f != nil {
			p.frames[n] = f
		} else if n > 0 {
			break
		}
	}
	// recalc in case there was a interrupted consensus
	p.reconsensusFromFrame(p.state.LastFinishedFrameN + 1)

}
