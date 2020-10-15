package main

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/yamazaki-ko/game/systems"
	"golang.org/x/image/font/gofont/gosmallcaps"
)

type myScene struct{}

func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
	engo.Files.Load("gopher.png")
	engo.Files.Load("Hero.png")
	engo.Files.Load("Goal.png")
	engo.Files.Load("./Tilesets/GroundTiles.png")
	engo.Files.Load("./Tilesets/GameObjectsTiles.png")
	engo.Files.Load("./Tilesets/DecorationTiles.png")
	engo.Files.Load("./Free Hero/LongHair_BlueTunic_Shield COMPRESSED.png")
	common.SetBackground(color.RGBA{120, 226, 250, 3})
}

func (*myScene) Setup(u engo.Updater) {
	// キーボード設定
	engo.Input.RegisterButton("MoveRight", engo.KeyD, engo.KeyArrowRight)
	engo.Input.RegisterButton("MoveLeft", engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton("Jump", engo.KeySpace)
	engo.Input.RegisterButton("スタート", engo.KeyOne)
	engo.Input.RegisterButton("背景設定", engo.KeyTwo)
	// World設定
	world, _ := u.(*ecs.World)
	// Systemの追加
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&systems.TileSystem{})
	world.AddSystem(&systems.ButtonSystem{})
	world.AddSystem(&systems.HUDTextSystem{})
	world.AddSystem(&systems.GoalSystem{})
}

func main() {
	opts := engo.RunOptions{
		Title:          "GoGame",
		Width:          400,
		Height:         300,
		StandardInputs: true,
		NotResizable:   true,
	}
	fmt.Println("GoGame Start")
	engo.Run(opts, &myScene{})
}

func (*myScene) Exit() {
	engo.Exit()
}
