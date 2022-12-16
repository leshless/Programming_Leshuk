package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
    "fmt"
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

func main() {
    gtk.Init(nil)
    
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal(err)
    }

    win.Connect("destroy", func() {
        win.Destroy()
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

    label, err := gtk.LabelNew("Хотите ли купить слона?")
    if err != nil {
        log.Panic(err)
    }
    

    var initialMenu, repeatMenu, endMenu Menu
    var closelist []*gtk.Window
    closelist = append(closelist, win)

    initialMenu.AddEntryWithAction("Да", nil, func(){
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            log.Fatal(err)
        }

        box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
        if err != nil {
            log.Panic(err)
        }

        label, err := gtk.LabelNew("К сожалению, сейчас у нас нет слона, так что мы не сможем Вам его продать.")
        if err != nil {
            log.Panic(err)
        }

        box.SetMarginTop(24)
        box.SetMarginBottom(24)
        box.SetMarginStart(24)
        box.SetMarginEnd(24)

        for _, window := range closelist{
            window.Destroy()
        }

        win.Add(box)
        box.Add(label)
        box.Add(endMenu.GtkWidget())

        win.ShowAll()
    })

    initialMenu.AddEntryWithAction("Нет", nil, func(){
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            log.Fatal(err)
        }

        box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
        if err != nil {
            log.Panic(err)
        }

        label, err := gtk.LabelNew("Может, Вы всё-таки хотите ли купить слона?")
        if err != nil {
            log.Panic(err)
        }

        box.SetMarginTop(24)
        box.SetMarginBottom(24)
        box.SetMarginStart(24)
        box.SetMarginEnd(24)

        win.Connect("destroy", func() {
            win.Destroy()
        })

        win.Add(box)
        box.Add(label)
        box.Add(repeatMenu.GtkWidget())

        closelist = append(closelist, win)
        fmt.Println(closelist)

        win.ShowAll()
    })

    repeatMenu.AddEntryWithAction("Да", nil, func(){
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            log.Fatal(err)
        }

        box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
        if err != nil {
            log.Panic(err)
        }

        label, err := gtk.LabelNew("К сожалению, сейчас у нас нет слона, так что мы не сможем Вам его продать.")
        if err != nil {
            log.Panic(err)
        }

        box.SetMarginTop(24)
        box.SetMarginBottom(24)
        box.SetMarginStart(24)
        box.SetMarginEnd(24)

        for _, window := range closelist{
            window.Destroy()
        }

        win.Add(box)
        box.Add(label)
        box.Add(endMenu.GtkWidget())

        win.ShowAll()
    })

    repeatMenu.AddEntryWithAction("Нет", nil, func(){
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            log.Fatal(err)
        }

        box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
        if err != nil {
            log.Panic(err)
        }

        label, err := gtk.LabelNew("Может, Вы всё-таки хотите ли купить слона?")
        if err != nil {
            log.Panic(err)
        }

        box.SetMarginTop(24)
        box.SetMarginBottom(24)
        box.SetMarginStart(24)
        box.SetMarginEnd(24)

        win.Connect("destroy", func() {
            win.Destroy()
        })

        win.Add(box)
        box.Add(label)
        box.Add(repeatMenu.GtkWidget())

        closelist = append(closelist, win)
        fmt.Println(closelist)

        win.ShowAll()
    })

    endMenu.AddEntryWithAction("Жаль!", nil, gtk.MainQuit)

    box.Add(label)
    box.Add(initialMenu.GtkWidget())

    win.ShowAll()
    gtk.Main()
}