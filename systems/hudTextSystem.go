package systems

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	// TextTITLE = タイトル
	TextTITLE = 0
	// TextSTART = 開始
	TextSTART = 1
	// TextGOAL = ゴール
	TextGOAL = 2
	// TextEND = 終了
	TextEND = 3
	// TextRETRY = リトライ
	TextRETRY = 4
	// TextBACKGROUND = 背景画像設定
	TextBACKGROUND = 5
)

// Text is an entity containing text printed to the screen
type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

// HUDTextMessage updates the HUD text based on messages sent from other systems
type HUDTextMessage struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.MouseComponent
	Line1, Line2, Line3, Line4 string
}

// HUDTextEntity is an entity for the text system. This keeps track of the position
// size and text associated with that position.
type HUDTextEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.MouseComponent
	Line1, Line2, Line3, Line4 string
}

// HUDTextSystem prints the text to our HUD based on the current state of the game
type HUDTextSystem struct {
	world      *ecs.World
	TextEntity []*Text
	entities   []HUDTextEntity
}

// New is
func (h *HUDTextSystem) New(w *ecs.World) {
	h.world = w
	AddText(w, h, TextTITLE)
	AddText(w, h, TextSTART)
}

// Update is
func (h *HUDTextSystem) Update(dt float32) {

}

// Remove takes an Entity out of the RenderSystem.
func (h *HUDTextSystem) Remove(basic ecs.BasicEntity) {
	for _, system := range h.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, text := range h.TextEntity {
				sys.Remove(text.BasicEntity)
			}
		}
	}
}

// TextRemove takes a TextEntity out of the HUDTextSystem.
func (h *HUDTextSystem) TextRemove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range h.TextEntity {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		h.TextEntity = append(h.TextEntity[:delete], h.TextEntity[delete+1:]...)
	}
}

// AddText adds an entity to the system
func AddText(w *ecs.World, h *HUDTextSystem, textNo int) {
	h.world = w
	// Entitiy作成
	text := &Text{BasicEntity: ecs.NewBasic()}
	// Compenent設定
	SetTextComponent(text, textNo)
	// Entitiy追加
	h.TextEntity = append(h.TextEntity, text)

	// SystemにEntity追加
	for _, system := range h.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&text.BasicEntity, &text.RenderComponent, &text.SpaceComponent)
		}
	}
}

// SetTextComponent set SpaceComponent and RenderComonent and MouseComponent.
func SetTextComponent(text *Text, textNo int) {
	// 初期化
	// SpaceComponent
	TextPositionX := (float32)(0)
	TextPositionY := engo.WindowHeight() - 220
	size := float64(40)
	// RenderComponent
	textDisplay := ""

	// テキストNo毎の処理
	switch textNo {
	case TextTITLE:
		textDisplay = "        GO GAME!"
	case TextSTART:
		textDisplay = "click"
		TextPositionX = engo.WindowWidth()/2 - 12
		TextPositionY = engo.WindowHeight() - 120
		size = 12
	case TextGOAL:
		textDisplay = "          GOAL!!"
	case TextEND:
		textDisplay = "       GAME OVER"
	case TextRETRY:
		textDisplay = "retry"
		TextPositionX = engo.WindowWidth()/2 - 13
		TextPositionY = engo.WindowHeight() - 120
		size = 12
	case TextBACKGROUND:
		textDisplay = "        SETTING"
	}

	// SpaceComponent
	text.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: TextPositionX, Y: TextPositionY},
	}

	// RenderComponent
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: size,
	}
	fnt.CreatePreloaded()

	text.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: textDisplay,
	}
	text.SetShader(common.TextHUDShader)
	text.RenderComponent.SetZIndex(1001)

}
