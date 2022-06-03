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

var applyMacButtonsScript = `
gsettings set org.gnome.desktop.wm.preferences button-layout "close,minimize,maximize:"
gsettings set org.pantheon.desktop.gala.appearance button-layout "close,minimize,maximize:"
gsettings set org.gnome.settings-daemon.plugins.xsettings overrides "{'Gtk/DialogsUseHeader': <0>, 'Gtk/ShellShowsAppMenu': <0>, 'Gtk/EnablePrimaryPaste': <1>, 'Gtk/DecorationLayout': <'close,minimize,maximize:menu'>}"
`

var applyWinButtonsScript = `
gsettings set org.gnome.desktop.wm.preferences button-layout ":minimize,maximize,close"
gsettings set org.pantheon.desktop.gala.appearance button-layout ":minimize,maximize,close"
gsettings set org.gnome.settings-daemon.plugins.xsettings overrides "{'Gtk/DialogsUseHeader': <0>, 'Gtk/ShellShowsAppMenu': <0>, 'Gtk/EnablePrimaryPaste': <1>, 'Gtk/DecorationLayout': <':menu,minimize,maximize,close'>}"
`

var applyWtfButtonsScript = `
gsettings set org.gnome.desktop.wm.preferences button-layout ""
gsettings set org.pantheon.desktop.gala.appearance button-layout ""
gsettings set org.gnome.settings-daemon.plugins.xsettings overrides "{'Gtk/DialogsUseHeader': <0>, 'Gtk/ShellShowsAppMenu': <0>, 'Gtk/EnablePrimaryPaste': <1>, 'Gtk/DecorationLayout': <':menue'>}"
`
var appTitle = `Elementary OS Buttons`

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
	w.SetSize(480, 480, webview.HintNone)

	w.Bind("checkButtonsStyle", checkButtonsStyle)
	w.Bind("addMinimizeButton", addMinimizeButton)
	w.Bind("restoreButtons", restoreButtons)
	w.Bind("applyMacButtons", applyMacButtons)
	w.Bind("applyWinButtons", applyWinButtons)
	w.Bind("applyWtfButtons", applyWtfButtons)

	w.Navigate("http://localhost:" + fmt.Sprint(port) + "/public/")
	w.Run()
}

func checkButtonsStyle() (r string) {
    out, err := exec.Command("gsettings", "get", "org.gnome.desktop.wm.preferences", "button-layout").Output()
    r = "unknown"

    if err != nil {
    	notification(appTitle, "Unable to execute gsettings command")
        panic(err)
    }

    eos := strings.Contains(string(out), "close:maximize")
    if eos {
    	r = "eos"
    }

    eosmin := strings.Contains(string(out), "close:minimize,maximize")
    if eosmin {
    	r = "eos+min"
    }

    mac := strings.Contains(string(out), "close,minimize,maximize:")
    if mac {
    	r = "mac"
    }

    win := strings.Contains(string(out), ":minimize,maximize,close")
    if win {
    	r = "win"
    }

    if len(strings.TrimSpace(string(out))) == 2 {
    	r = "wtf"
    }

	return
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
        notification(appTitle, "Unable to restore default style")
        return false
    }

	return true
}

func applyMacButtons() (bool) {
    c := exec.Command("bash")
    c.Stdin = strings.NewReader(applyMacButtonsScript)

    e := c.Run()
    if e != nil {
        fmt.Println(e)
        notification(appTitle, "Unable to apply Mac style")
        return false
    }

	notification(appTitle, "Succesfully apply Mac style")

	return true
}

func applyWinButtons() (bool) {
    c := exec.Command("bash")
    c.Stdin = strings.NewReader(applyWinButtonsScript)

    e := c.Run()
    if e != nil {
        fmt.Println(e)
        notification(appTitle, "Unable to apply Windows style")
        return false
    }

	notification(appTitle, "Succesfully apply Windows style")

	return true
}

func applyWtfButtons() (bool) {
    c := exec.Command("bash")
    c.Stdin = strings.NewReader(applyWtfButtonsScript)

    e := c.Run()
    if e != nil {
        fmt.Println(e)
        notification(appTitle, "Unable to apply Ninja style")
        return false
    }

	notification("Yoda", "Try Not. Do or do not, there is no try.")

	return true
}

func notification(title string, message string) {
	beeep.Notify(title, message, "")
}
