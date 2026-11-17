//go:build windows

// For-MAC: a small Windows utility that normalizes MAC addresses to
// AA:BB:CC:DD:EE:FF, one per line, no matter how they were typed or pasted.
package main

import (
	"bytes"
	_ "embed"
	"image/png"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"

	"forgejo.smartspacesprogramming.dev/yara/For-MAC/internal/macfmt"
)

const (
	appTitle   = "MAC Address Formatter - Version 1.2"
	licenseURL = "https://creativecommons.org/licenses/"
)

//go:embed assets/logo_small.png
var logoPNG []byte

func main() {
	var (
		mw      *walk.MainWindow
		inputTE *walk.TextEdit
		outTE   *walk.TextEdit
		logo    walk.Image
		link    *walk.Label
	)

	if img, err := png.Decode(bytes.NewReader(logoPNG)); err == nil {
		if bmp, err := walk.NewBitmapFromImageForDPI(img, 96); err == nil {
			logo = bmp
		}
	}

	mono := Font{Family: "Consolas", PointSize: 11}

	formatAll := func() {
		outTE.SetText(macfmt.FormatAll(inputTE.Text(), "\r\n"))
	}
	copyOutput := func() {
		if err := walk.Clipboard().SetText(outTE.Text()); err != nil {
			walk.MsgBox(mw, "Copy failed", err.Error(), walk.MsgBoxIconError)
		}
	}
	clearAll := func() {
		inputTE.SetText("")
		outTE.SetText("")
		inputTE.SetFocus()
	}

	err := MainWindow{
		AssignTo: &mw,
		Title:    appTitle,
		MinSize:  Size{Width: 640, Height: 480},
		Size:     Size{Width: 960, Height: 680},
		Layout:   VBox{Margins: Margins{Left: 16, Top: 12, Right: 16, Bottom: 12}, Spacing: 10},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true, Spacing: 18},
				Children: []Widget{
					ImageView{Image: logo, Mode: ImageViewModeIdeal},
					Composite{
						Layout: VBox{MarginsZero: true, Spacing: 2},
						Children: []Widget{
							VSpacer{},
							Label{Text: appTitle, Font: Font{Family: "Segoe UI", PointSize: 18, Bold: true}},
							Label{Text: "Internal formatting utility", TextColor: walk.RGB(0x55, 0x55, 0x55)},
							VSpacer{},
						},
					},
					HSpacer{},
				},
			},
			Composite{
				StretchFactor: 10,
				Layout:        HBox{MarginsZero: true, Spacing: 16},
				Children: []Widget{
					Composite{
						Layout: VBox{MarginsZero: true},
						Children: []Widget{
							Label{Text: "Input  (one value per line)", Font: Font{Family: "Segoe UI", PointSize: 9, Bold: true}},
							TextEdit{AssignTo: &inputTE, Font: mono, VScroll: true, MinSize: Size{Height: 300},
								OnKeyDown: func(key walk.Key) {
									// Ctrl+Enter formats without reaching for the mouse.
									if key == walk.KeyReturn && walk.ControlDown() {
										formatAll()
									}
								}},
						},
					},
					Composite{
						Layout: VBox{MarginsZero: true},
						Children: []Widget{
							Label{Text: "Output", Font: Font{Family: "Segoe UI", PointSize: 9, Bold: true}},
							TextEdit{AssignTo: &outTE, Font: mono, ReadOnly: true, VScroll: true, MinSize: Size{Height: 300}},
						},
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true, Spacing: 10},
				Children: []Widget{
					HSpacer{},
					PushButton{Text: "&Do", MinSize: Size{Width: 110}, OnClicked: formatAll, ToolTipText: "Format (Ctrl+Enter)"},
					PushButton{Text: "&Copy Output", MinSize: Size{Width: 110}, OnClicked: copyOutput},
					PushButton{Text: "C&lear", MinSize: Size{Width: 110}, OnClicked: clearAll},
					HSpacer{},
				},
			},
			Composite{
				Layout: VBox{MarginsZero: true, Spacing: 2},
				Children: []Widget{
					Label{Text: "Made by YZ for Smart Spaces WI LLC.", Alignment: AlignHCenterVCenter, TextColor: walk.RGB(0x66, 0x66, 0x66)},
					Composite{
						Layout: HBox{MarginsZero: true},
						Children: []Widget{
							HSpacer{},
							Label{Text: "CC BY 4.0: ", TextColor: walk.RGB(0x66, 0x66, 0x66)},
							Label{
								AssignTo:    &link,
								Text:        "creativecommons.org/licenses/",
								TextColor:   walk.RGB(0x1a, 0x5f, 0xb4),
								Font:        Font{Family: "Segoe UI", PointSize: 9, Underline: true},
								ToolTipText: licenseURL,
								OnMouseUp: func(x, y int, b walk.MouseButton) {
									if b == walk.LeftButton {
										win.ShellExecute(0, nil, syscallStr("open"), syscallStr(licenseURL), nil, win.SW_SHOWNORMAL)
									}
								},
							},
							HSpacer{},
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}

	link.SetCursor(walk.CursorHand())
	inputTE.SetFocus()
	mw.Run()
}
