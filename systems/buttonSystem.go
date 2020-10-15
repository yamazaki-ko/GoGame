package systems

import (
	"fmt"
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	// BtnInit : 未使用
	BtnInit = 0
	// BtnSTART : 開始
	BtnSTART = 1
	// BtnEND : 終了
	BtnEND = 2
	// BtnRETRY : リトライ
	BtnRETRY = 3
	// BtnBACKGROUND : 背景画像設定
	BtnBACKGROUND = 4
)

// Button :
type Button struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
	ButtonNo int
	enable   bool
}

// ButtonSystem :
type ButtonSystem struct {
	world        *ecs.World
	ButtonEntity []*Button
}

// Remove :
func (bs *ButtonSystem) Remove(entity ecs.BasicEntity) {
}

// Update :
func (bs *ButtonSystem) Update(dt float32) {
	for _, item := range bs.ButtonEntity {
		// 有効ボタンがクリック
		if item.enable == true && item.MouseComponent.GetMouseComponent().Clicked {
			switch item.ButtonNo {
			case BtnSTART:
				fmt.Println("スタートボタンクリック")
				GameStart(bs)
			case BtnEND:
				fmt.Println("終了ボタンクリック")
				engo.Exit()
			case BtnRETRY:
				fmt.Println("リトライボタンクリック")
				GameStart(bs)
			case BtnBACKGROUND:
				fmt.Println("背景設定ボタンクリック")
			default:
				fmt.Println("それ以外のボタンクリック")
			}
		}
	}
}

// New is
func (bs *ButtonSystem) New(w *ecs.World) {
	AddButton(w, bs, BtnSTART)
	AddButton(w, bs, BtnBACKGROUND)
}

// GameStart is remove buttons and make player.
func GameStart(bs *ButtonSystem) {
	for _, item := range bs.ButtonEntity {
		switch item.ButtonNo {
		case BtnSTART:
			item.enable = false
			fmt.Println("スタートボタン無効")
		case BtnEND:
			item.enable = false
			fmt.Println("終了ボタン無効")
		case BtnRETRY:
			item.enable = false
			fmt.Println("リトライボタン無効")
		case BtnBACKGROUND:
			item.enable = false
			fmt.Println("背景ボタン無効")
		default:
			fmt.Println("それ以外")
		}
	}
	for _, system := range bs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			bs.world.AddSystem(&PlayerSystem{})
			for _, btn := range bs.ButtonEntity {
				sys.Remove(btn.BasicEntity)
			}
		case *HUDTextSystem:
			for _, text := range sys.TextEntity {
				sys.Remove(text.BasicEntity)
				sys.TextRemove(text.BasicEntity)
			}
		}
	}
}

// AddButton adds an entity to the system
func AddButton(w *ecs.World, bs *ButtonSystem, btnNo int) {
	bs.world = w
	// Entitiy作成
	btn := &Button{BasicEntity: ecs.NewBasic()}
	// ボタン有効
	btn.enable = true
	// Compenent設定
	SetButtonComponent(btn, btnNo)
	// Entitiy追加
	bs.ButtonEntity = append(bs.ButtonEntity, btn)

	// SystemにEntity追加
	for _, system := range bs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&btn.BasicEntity, &btn.RenderComponent, &btn.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&btn.BasicEntity, &btn.MouseComponent, &btn.SpaceComponent, nil)
		}
	}
}

// SetButtonComponent set SpaceComponent and RenderComonent and MouseComponent.
func SetButtonComponent(btn *Button, btnNo int) {
	// 初期化
	// SpaceComponent
	BsPositionX := (float32)(0)
	BsPositionY := (float32)(0)
	width := (float32)(0)
	height := (float32)(0)
	// RenderComponent
	url := "Hero.png"
	scaleX := (float32)(0)
	scaleY := (float32)(0)
	color := color.White

	// ボタン毎の処理
	switch btnNo {
	case BtnSTART:
		btn.ButtonNo = BtnSTART
		BsPositionX = engo.WindowWidth()/2 - 20
		BsPositionY = engo.WindowHeight()/2 - 20
		width = (float32)(50)
		height = (float32)(65)
		scaleX = (float32)(1.2)
		scaleY = (float32)(1.2)
		//fmt.Printf(strconv.FormatFloat(float64(engo.WindowWidth()/2), 'f', 2, 64))
		fmt.Println("ButtonEntity(START) NEW")
	case BtnEND:
		btn.ButtonNo = BtnEND
		url = "gopher.png"
		fmt.Println("ButtonEntity(END) NEW")
	case BtnRETRY:
		btn.ButtonNo = BtnRETRY
		BsPositionX = engo.WindowWidth()/2 - 20
		BsPositionY = engo.WindowHeight()/2 - 20
		width = (float32)(50)
		height = (float32)(65)
		scaleX = (float32)(1.2)
		scaleY = (float32)(1.2)
		fmt.Println("ButtonEntity(RETRY) NEW")
		// カメラを移動する
		engo.Mailbox.Dispatch(common.CameraMessage{
			Axis:        common.XAxis,
			Value:       engo.WindowWidth() / 2,
			Incremental: false,
		})
	case BtnBACKGROUND:
		btn.ButtonNo = BtnBACKGROUND
		BsPositionX = float32(-16)
		BsPositionY = engo.WindowHeight() - 100
		width = (float32)(64)
		height = (float32)(32)
		url = "gopher.png"
		scaleX = (float32)(0.3)
		scaleY = (float32)(0.3)
		fmt.Println("ButtonEntity(Hidden) NEW")
	default:
		fmt.Println("ButtonEntity(default) NEW")
	}

	// SpaceComponent
	btn.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: BsPositionX, Y: BsPositionY},
		Width:    width,
		Height:   height,
	}

	// RenderComponent
	texture, err := common.LoadedSprite(url)
	if err != nil {
		fmt.Println("Unable to load texture: " + err.Error())
	}
	btn.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: scaleX, Y: scaleY},
		Color:    color,
	}

	// MouseComponent
	btn.MouseComponent = common.MouseComponent{Track: false}
}
