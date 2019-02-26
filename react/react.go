// package react
//
// const testVersion = 4
//
// /* reactor */
//
// type reactor struct {
// 	cells []*compuCell
// }
//
// func New() Reactor {
// 	return &reactor{}
// }
//
// func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
// 	old := f(c.Value())
// 	cc := new(compuCell)
// 	cc.cb = make(map[CallbackHandle]func(int))
// 	cc.eval = func() int {
// 		v := f(c.Value())
// 		if v != old {
// 			for _, cb := range cc.cb {
// 				cb(v)
// 			}
// 			old = v
// 		}
// 		return v
// 	}
// 	r.cells = append(r.cells, cc)
// 	return cc
// }
//
// func (r *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
// 	old := f(c1.Value(), c2.Value())
// 	cc := new(compuCell)
// 	cc.cb = make(map[CallbackHandle]func(int))
// 	cc.eval = func() int {
// 		v := f(c1.Value(), c2.Value())
// 		if v != old {
// 			for _, cb := range cc.cb {
// 				cb(v)
// 			}
// 			old = v
// 		}
// 		return v
// 	}
// 	r.cells = append(r.cells, cc)
// 	return cc
// }
//
// func (r *reactor) CreateInput(i int) InputCell {
// 	return &inputCell{
// 		value:   i,
// 		reactor: r,
// 	}
// }
//
// /* input cell */
//
// type inputCell struct {
// 	value int
// 	*reactor
// }
//
// func (c *inputCell) SetValue(i int) {
// 	if c.value != i {
// 		c.value = i
// 		for _, cc := range c.cells {
// 			cc.eval()
// 		}
// 	}
// }
//
// func (c *inputCell) Value() int {
// 	return c.value
// }
//
// /* compute cell */
//
// type compuCell struct {
// 	eval func() int
// 	cb   map[CallbackHandle]func(int)
// }
//
// func (c *compuCell) AddCallback(f func(int)) CallbackHandle {
// 	c.cb[&f] = f
// 	return &f // guaranteed to be uniq
// }
//
// func (c *compuCell) RemoveCallback(h CallbackHandle) {
// 	delete(c.cb, h)
// }
//
// func (c *compuCell) Value() int {
// 	return c.eval()
// }
