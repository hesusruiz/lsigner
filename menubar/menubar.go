package menubar

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func fileMenu(a fyne.App, w fyne.Window) *fyne.Menu {

	mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	mailItem.Icon = theme.MailComposeIcon()

	// *********************************
	// Prepare the New item and its submenus

	fileItem := fyne.NewMenuItem("File", func() {
		fmt.Println("Menu New->File")
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			fmt.Println("URI:", uri)
		}, w)
	})
	fileItem.Icon = theme.FileIcon()

	dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	dirItem.Icon = theme.FolderIcon()

	otherItem := fyne.NewMenuItem("Other", nil)
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		mailItem,
	)

	newItem := fyne.NewMenuItem("New", nil)
	newItem.ChildMenu = fyne.NewMenu("",
		fileItem,
		dirItem,
		otherItem,
	)
	// end of New item and submenus
	// **************************************

	fileOpenItem := fyne.NewMenuItem("Open...", func() {
		fmt.Println("Menu New->File")
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			fmt.Println("URI:", uri)
		}, w)
	})

	// ***************************************
	// The top-level File menu item and sub-items
	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", fileOpenItem, newItem)

	// If it is not mobile, add a Settings menu item
	if device := fyne.CurrentDevice(); !device.IsMobile() && !device.IsBrowser() {
		openSettings := func() {
			w := a.NewWindow("Fyne Settings")
			w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
			w.Resize(fyne.NewSize(440, 520))
			w.Show()
		}
		settingsItem := fyne.NewMenuItem("Settings", openSettings)
		settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
		settingsItem.Shortcut = settingsShortcut
		w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
			openSettings()
		})

		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}

	return file

}

func editMenu(a fyne.App, w fyne.Window) *fyne.Menu {

	// Cut
	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut

	// Copy
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut

	// Paste
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut

	// Find
	performFind := func() { fmt.Println("Menu Find") }
	findItem := fyne.NewMenuItem("Find", performFind)

	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierControl}

	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	// ***************************************
	// The top-level File menu item and sub-items
	// a quit item will be appended to our first (File) menu
	edit := fyne.NewMenu("Edit",
		cutItem,
		copyItem,
		pasteItem,
		fyne.NewMenuItemSeparator(),
		findItem)

	return edit
}

func helpMenu(a fyne.App, w fyne.Window) *fyne.Menu {
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://dome-marketplace.eu")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://dome-marketplace.eu")
			_ = a.OpenURL(u)
		}),
	)

	return helpMenu
}

// makeMenu creates the main menu bar of the application
func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {

	// The mainMenuBar menu bar of the application
	mainMenuBar := fyne.NewMainMenu(
		fileMenu(a, w),
		editMenu(a, w),
		helpMenu(a, w),
	)
	return mainMenuBar
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
