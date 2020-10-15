package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	// GoalPoint : GOALまでの距離
	GoalPoint = TileNum*CellWidth - 15*CellWidth
)

// Goal : Entitiy
type Goal struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// GoalSystem :
type GoalSystem struct {
	world      *ecs.World
	goalEntity *Goal
}

// Remove :
func (*GoalSystem) Remove(ecs.BasicEntity) {

}

// Update :
func (gs *GoalSystem) Update(dt float32) {

}

// New :
func (gs *GoalSystem) New(w *ecs.World) {
	gs.world = w
	// ゴールの作成
	goal := Goal{BasicEntity: ecs.NewBasic()}

	// 初期の配置
	positionX := GoalPoint
	positionY := int(engo.WindowHeight() - 112)
	goal.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: float32(positionX), Y: float32(positionY)},
		Width:    30,
		Height:   30,
	}

	// 画像の読み込み
	texture, err := common.LoadedSprite("Goal.png")
	if err != nil {
		fmt.Println("Unable to load texture: " + err.Error())
	}
	goal.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 0.05, Y: 0.05},
	}
	gs.goalEntity = &goal

	for _, system := range gs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&goal.BasicEntity, &goal.RenderComponent, &goal.SpaceComponent)
		}
	}

}
