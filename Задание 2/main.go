package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
)

type Menu struct {
    entries []MenuEntry
}

func (m *Menu) AddEntryWithAction(label string, next *Menu, action func()) {
    if next == nil { next = m }
    
    m.entries = append(m.entries, MenuEntry{
        Label: label,
        Next: next,
        Action: action,
    })
}

func (m *Menu) AddEntry(label string, next *Menu) {
    m.AddEntryWithAction(label, next, nil)
}

type MenuEntry struct {
    Label string
    Next *Menu
    Action func()
}

func (e MenuEntry) Use() *Menu {
    if e.Action != nil { e.Action() }
    
    return e.Next
}

func ChangeMenu(box *gtk.Box, m *Menu){
    box.GetChildren().Foreach(func(button any){
        button.(*gtk.Widget).Destroy()
    })

    box.Add(m.GtkWidget())
    box.ShowAll()
}

func (m *Menu) GtkWidget() *gtk.Widget {
    box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
    if err != nil {
        log.Panic(err)
    }

    for _, entrie := range m.entries{
        button, err := gtk.ButtonNewWithLabel(entrie.Label)
        if err != nil {
            log.Panic(err)
        }

        entriecopy := entrie
        button.Connect("clicked", func(){
                next := entriecopy.Use()
                if next != nil{ChangeMenu(box, next)}
        })

        box.Add(button)
    }

    return &box.Widget
}

func makeMainMenu(info *gtk.Label) *Menu {
    var mainMenu, newGameMenu, optionsMenu Menu
    
    var gameResult bool
    
    playGame := func() {
        if gameResult {
            info.SetText("Вы выиграли!")
        } else {
            info.SetText("Вы проиграли!")
        }
    }

    clearInfo := func() { info.SetText("") }
    
    mainMenu.AddEntryWithAction("Новая игра", &newGameMenu, playGame)
    
    mainMenu.AddEntry("Настройки", &optionsMenu)
    mainMenu.AddEntryWithAction("Выйти", nil, gtk.MainQuit)
    
    newGameMenu.AddEntryWithAction("Начать заново", nil, playGame)
    newGameMenu.AddEntryWithAction("Выйти в главное меню", &mainMenu, clearInfo)
    
    optionsMenu.AddEntryWithAction("Хочу всегда выигрывать", &mainMenu, func() {
        gameResult = true
    })
    
    optionsMenu.AddEntryWithAction("Хочу всегда проигрывать", &mainMenu, func() {
        gameResult = false
    })
    
    return &mainMenu
}

func main() {
    gtk.Init(nil)
    
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

    if err != nil {
        log.Fatal(err)
    }

    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
    if err != nil {
        log.Panic(err)
    }

    box.SetMarginTop(24)
    box.SetMarginBottom(24)
    box.SetMarginStart(24)
    box.SetMarginEnd(24)

    win.Add(box)

    label, err := gtk.LabelNew("")

    if err != nil {
        log.Panic(err)
    }
    
    box.Add(label)
    box.Add(makeMainMenu(label).GtkWidget())

    win.ShowAll()
    gtk.Main()
}