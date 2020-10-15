package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

// Gopher is charactor
type Gopher struct {
	ecs.BasicEntity
	// RenderComponent:Entityの見た目に関する情報
	common.RenderComponent
	// SpaceComponent:位置に関する情報
	common.SpaceComponent
}

// GopherSystem is system
type GopherSystem struct {
	world *ecs.World
	// x軸座標
	positionX int
	// y軸座標
	positionY    int
	gopherEntity []*Gopher
	texture      *common.Texture
}

// Remove is
func (*GopherSystem) Remove(ecs.BasicEntity) {}

// Update is
func (ts *GopherSystem) Update(dt float32) {
}

// New is
func (ts *GopherSystem) New(w *ecs.World) {
}
