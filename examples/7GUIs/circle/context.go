package main

type OpType int

const (
	OP_ADD OpType = iota
	OP_REMOVE
	OP_UPDATE
)

type Op struct {
	Type   OpType
	Circle *Circle
}

const (
	MAX_RADIUS = 100
	MIN_RADIUS = 10
	DEF_RADIUS = 30
)

type Circle struct {
	X, Y int
	R    int
	ID   int
}

func (c *Circle) Copy() *Circle {
	if c == nil {
		return nil
	}

	return &Circle{
		X:  c.X,
		Y:  c.Y,
		R:  c.R,
		ID: c.ID,
	}
}

type Context struct {
	ops         []*Op
	lastOpIndex int

	circles      []*Circle
	lastCircleID int
	pickedCircle *Circle
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.ops = make([]*Op, 0, 10)
	ctx.circles = make([]*Circle, 0, 10)
	ctx.lastOpIndex = -1
	return ctx
}

func (ctx *Context) NewCircle(x, y int) *Circle {
	c := &Circle{
		X:  x,
		Y:  y,
		R:  DEF_RADIUS,
		ID: ctx.lastCircleID,
	}
	ctx.lastCircleID++
	return c
}

func (ctx *Context) AddCircle(c *Circle) {
	ctx.circles = append(ctx.circles, c)
}

func (ctx *Context) RemoveCircle(c *Circle) {
	for i, cInPool := range ctx.circles {
		if cInPool.ID == c.ID {
			ctx.circles = append(ctx.circles[:i], ctx.circles[i+1:]...)
			if ctx.IsCirclePicked(c) {
				ctx.pickedCircle = nil
			}
			return
		}
	}
}

func (ctx *Context) UpdateCircle(c *Circle) {
	for _, cInPool := range ctx.circles {
		if cInPool.ID == c.ID {
			cInPool.R = c.R
		}
	}
}

func (ctx *Context) PickCircle(pickX, pickY int) *Circle {
	minDoubleR := MAX_RADIUS * MAX_RADIUS
	index := -1
	for i, cInPool := range ctx.circles {
		dx, dy := cInPool.X-pickX, cInPool.Y-pickY
		doubleR := dx*dx + dy*dy
		if doubleR <= cInPool.R*cInPool.R && doubleR < minDoubleR {
			minDoubleR = doubleR
			index = i
		}
	}

	if index < 0 {
		ctx.pickedCircle = nil
		return nil
	}

	ctx.pickedCircle = ctx.circles[index]

	return ctx.circles[index].Copy()
}

func (ctx *Context) UnPickCircle() {
	ctx.pickedCircle = nil
}

func (ctx *Context) IsCirclePicked(c *Circle) bool {
	return ctx.pickedCircle != nil && c.ID == ctx.pickedCircle.ID
}

func (ctx *Context) AddOp(t OpType, c *Circle) {
	if c == nil {
		return
	}

	op := &Op{
		Type:   t,
		Circle: c.Copy(),
	}

	ctx.lastOpIndex++
	ctx.ops = ctx.ops[:ctx.lastOpIndex]
	ctx.ops = append(ctx.ops, op)
}

func (ctx *Context) Undo() bool {
	if ctx.lastOpIndex < 0 {
		return false
	}

	op := ctx.ops[ctx.lastOpIndex]
	ctx.lastOpIndex--
	switch op.Type {
	case OP_ADD:
		ctx.RemoveCircle(op.Circle)
	case OP_REMOVE:
		ctx.AddCircle(op.Circle)
	case OP_UPDATE:
		undoCircle := ctx.ops[ctx.lastOpIndex].Circle
		ctx.lastOpIndex--
		ctx.UpdateCircle(undoCircle)
	}

	return true
}

func (ctx *Context) Redo() bool {
	if ctx.lastOpIndex+1 >= len(ctx.ops) {
		return false
	}

	ctx.lastOpIndex++
	op := ctx.ops[ctx.lastOpIndex]
	switch op.Type {
	case OP_ADD:
		ctx.AddCircle(op.Circle)
	case OP_REMOVE:
		ctx.RemoveCircle(op.Circle)
	case OP_UPDATE:
		ctx.lastOpIndex++
		redoCircle := ctx.ops[ctx.lastOpIndex].Circle
		ctx.UpdateCircle(redoCircle)
	}

	return true
}

func (ctx *Context) HasUndo() bool {
	return ctx.lastOpIndex >= 0
}

func (ctx *Context) HasRedo() bool {
	return ctx.lastOpIndex+1 < len(ctx.ops)
}

func (ctx *Context) Circles() []*Circle {
	return ctx.circles
}

func (ctx *Context) PickedCircle() *Circle {
	return ctx.pickedCircle.Copy()
}

func (ctx *Context) HasPickedCircle() bool {
	return ctx.pickedCircle != nil
}
