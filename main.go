package main

import (
	"time"
	"strings"	
	"github.com/gen2brain/beeep"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var value *walk.TextEdit
	var hotkeyLabel *walk.Label
	var slider *walk.Slider
	var mainWnd *walk.MainWindow
	var sendEnter *walk.CheckBox

	defaultBinding, err := SetNewBinding([]int{0, 1}, 10)
	if err != nil {
		walk.MsgBox(nil, "Failed to bind a default hotkey", err.Error()+"\nTry to change the hotkey to something another", walk.MsgBoxApplModal)
	}

	// go func() {
	// 	ListenForHotkey(func() {
	// 		time.Sleep(time.Millisecond * 150)

	// 		valueToWrite := value.Text()
	// 		if len(valueToWrite) == 0 {
	// 			return
	// 		}
	// 		beeep.Beep(beeep.DefaultFreq*2, beeep.DefaultDuration/4) // emulate scanner sound :)
			
	// 		EmulateTyping(valueToWrite, slider.Value(), sendEnter.CheckState() == walk.CheckChecked)
	// 	})
	// }() // listen in bg, non-blocking the main GUI thread
	
	go func() {
		ListenForHotkey(func() {
			// Give a brief pause before processing
			time.Sleep(150 * time.Millisecond)
	
			rawValue := value.Text()
			if len(rawValue) == 0 {
				return
			}
	
			// Split the input by semicolons so we can simulate typing for each segment
			segments := strings.Split(rawValue, ";")
			firstSegment := true
			for _, segment := range segments {
				// Trim each segment to remove any extra whitespace
				segment = strings.TrimSpace(segment)
				if len(segment) == 0 {
					continue
				}
				// Emit a beep to emulate the scanner sound for each segment
				if firstSegment {
					beeep.Beep(beeep.DefaultFreq*2, beeep.DefaultDuration/4) // Beep only for the first segment
					firstSegment = false // Set flag to false after the first beep
				}
				
				// Call EmulateTyping for the current segment. The slider value
				// and sendEnter state are applied consistently for each sequence.
				EmulateTyping(segment, slider.Value(), sendEnter.CheckState() == walk.CheckChecked)
			}
		})
	}() // Listen in background, non-blocking the main GUI thread

	ico, err := walk.NewIconFromResourceId(2) // icon id from manifest
	if err != nil {
		panic("could not load icon")
	}

	MainWindow{
		AssignTo: &mainWnd,
		Title:    "Barcode Reader Emulator",
		Size:     Size{Width: 350, Height: 135},
		Icon:     ico,
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Label{Text: "Value"},
					TextEdit{
						AssignTo:      &value,
						Text:          "test",
						MaxLength:     128,
						CompactHeight: true,
						VScroll:       true,
					},
					PushButton{
						Text:    "Scan from the screen",
						MaxSize: Size{Width: 40, Height: 20},
						OnClicked: func() {
							val, err := ScanScreenBarcode()
							if err != nil {
								walk.MsgBox(mainWnd, "Failed to scan", "Could not find or parse a barcode. Try to zoom in a bit or choose another one.\nParser error: "+err.Error(), walk.MsgBoxApplModal)
							} else {
								value.SetText(val)
							}
						},
					},
				},
			},
			HSplitter{
				Children: []Widget{
					Label{Text: "Hotkey"},
					Label{AssignTo: &hotkeyLabel, Text: defaultBinding},
					PushButton{
						Text:    "Change",
						MaxSize: Size{Width: 40, Height: 20},
						OnClicked: func() {
							var lb *walk.ListBox
							var cb *walk.ComboBox
							var dlg *walk.Dialog
							modLabels := make([]string, len(AllModifiers))
							for i := 0; i < len(AllModifiers); i++ {
								modLabels[i] = modifierToString(AllModifiers[i])
							}
							hotkeyLabels := make([]string, len(AllHotkeys))
							for i := 0; i < len(AllHotkeys); i++ {
								hotkeyLabels[i] = hotkeyToString(AllHotkeys[i])
							}

							Dialog{
								AssignTo: &dlg,
								Title:    "Select a new hotkey",
								Layout:   HBox{},
								MinSize:  Size{Width: 250, Height: 110},
								Children: []Widget{
									ListBox{
										AssignTo:       &lb,
										MultiSelection: true,
										Model:          modLabels,
									},
									ComboBox{
										Model:        hotkeyLabels,
										AssignTo:     &cb,
										CurrentIndex: 0,
									},
									PushButton{Text: "OK", OnClicked: func() {
										selModIdx := lb.SelectedIndexes()
										selHotkeyIdx := cb.CurrentIndex()
										bindingLabel, err := SetNewBinding(selModIdx, selHotkeyIdx)
										if err != nil {
											walk.MsgBox(mainWnd, "Failed", err.Error(), walk.MsgBoxApplModal)
										} else {
											hotkeyLabel.SetText(bindingLabel)
											dlg.Close(walk.DlgCmdOK)
										}
									}},
								},
							}.Run(mainWnd)
						},
					},
				},
			}, 
			HSplitter{
				Children: []Widget{
					Label{Text: "Input Key Delay"},
					Slider{AssignTo: &slider, MinValue: 10, MaxValue: 100},
					CheckBox{ 
						Text: "Send ENTER at the end", 
						AssignTo: &sendEnter,
						MinSize: Size{Width: 40, Height: 20},
					},		
				},
			},
		},
	}.Run()
}
