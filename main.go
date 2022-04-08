package main

import (
	"embed"
	"fmt"
	"strings"
	"net"
	"net/http"
    "os/exec"
	"github.com/webview/webview"
	"github.com/gen2brain/beeep"
)

//go:embed public/*
var files embed.FS

var minimizeButtonScript = `
gsettings set org.gnome.desktop.wm.preferences button-layout "close:minimize,maximize"
gsettings set org.pantheon.desktop.gala.appearance button-layout "close:minimize,maximize"
gsettings set org.gnome.settings-daemon.plugins.xsettings overrides "{'Gtk/DialogsUseHeader': <0>, 'Gtk/ShellShowsAppMenu': <0>, 'Gtk/EnablePrimaryPaste': <1>, 'Gtk/DecorationLayout': <'close:minimize,maximize,menu'>}"
`

var restoreButtonsScript = `
gsettings set org.gnome.desktop.wm.preferences button-layout "close:maximize"
gsettings set org.pantheon.desktop.gala.appearance button-layout "close:maximize"
gsettings set org.gnome.settings-daemon.plugins.xsettings overrides "{'Gtk/DialogsUseHeader': <0>, 'Gtk/ShellShowsAppMenu': <0>, 'Gtk/EnablePrimaryPaste': <1>, 'Gtk/DecorationLayout': <'close:maximize,menu'>}"
`
var appTitle = `Elementary OS Minimize Button`

func main() {
	// Serveur à partir du système de fichiers intégré
	listener, err := net.Listen("tcp", ":0") // Ouvrir sur n'importe quel port libre
	if err != nil {
		panic(err)
	}

	port := listener.Addr().(*net.TCPAddr).Port

	go func() {
		err := http.Serve(listener, http.FileServer(http.FS(files)))
		if err != nil {
			panic(err)
		}
	}()

	// Initialiser la vue web
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle(appTitle)
	w.SetSize(480, 160, webview.HintNone)

	w.Bind("checkCloseButtonPosition", checkCloseButtonPosition)
	w.Bind("getButtonsLayout", getButtonsLayout)
	w.Bind("addMinimizeButton", addMinimizeButton)
	w.Bind("restoreButtons", restoreButtons)

	w.Navigate("http://localhost:" + fmt.Sprint(port) + "/public/")
	w.Run()
}

func checkCloseButtonPosition() (bool) {
    out, err := exec.Command("gsettings", "get", "org.gnome.desktop.wm.preferences", "button-layout").Output()

    if err != nil {
    	notification(appTitle, "Unable to execute gsettings command")
        panic(err)
    }

    res := strings.Contains(string(out), "close:")
    if res {
    	return true
    }

	return false
}

func getButtonsLayout() (bool) {
    out, err := exec.Command("gsettings", "get", "org.gnome.desktop.wm.preferences", "button-layout").Output()

    if err != nil {
    	notification(appTitle, "Unable to execute gsettings command")
        panic(err)
    }

    res := strings.Contains(string(out), "minimize")
    if res {
    	return true
    }

	return false
}

func addMinimizeButton() (bool) {
    c := exec.Command("bash")
    c.Stdin = strings.NewReader(minimizeButtonScript)

    e := c.Run()
    if e != nil {
        fmt.Println(e)
        notification(appTitle, "Unable to add button")
        return false
    }

	notification(appTitle, "Succesfully added minimize button")

	return true
}

func restoreButtons() (bool) {
    c := exec.Command("bash")
    c.Stdin = strings.NewReader(restoreButtonsScript)

    e := c.Run()
    if e != nil {
        fmt.Println(e)
        notification(appTitle, "Unable to remove button")
        return false
    }

	notification(appTitle, "Succesfully removed minimize button")

	return true
}

func notification(title string, message string) {
	beeep.Notify(title, message, "")
}
