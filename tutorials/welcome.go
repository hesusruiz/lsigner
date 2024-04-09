package tutorials

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/hesusruiz/lsigner/logos"
	"github.com/imroc/req/v3"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(logos.ResourceDomeBluePng)
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(256, 256))
	}
	return makeTableOfCredentials()
	// return container.NewCenter(container.NewStack(
	// widget.NewLabelWithStyle("Welcome to the DOME Verifiable Credential Signer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	// logo,
	// return makeTableOfCredentials()
	// &widget.Button{
	// 	Text:       "Retrieve credentials to sign",
	// 	Importance: widget.HighImportance,
	// 	OnTapped:   func() { fmt.Println("high importance button") },
	// },

	// container.NewHBox(
	// 	widget.NewHyperlink("DOME-Marketplace.eu", parseURL("https://dome-marketplace.eu/")),
	// 	widget.NewLabel("-"),
	// 	widget.NewHyperlink("documentation", parseURL("https://developer.fyne.io/")),
	// ),
	// widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	// 	))
}

func makeTableOfCredentials() fyne.CanvasObject {

	var records []map[string]string
	client := req.C()
	response := client.Get("https://issuer.mycredential.eu/apiuser/retrievecredentials").Do()
	if response.Err != nil {
		return nil
	}
	if response.IsErrorState() {
		return nil
	}
	err := response.Into(&records)
	if err != nil {
		return nil
	}

	fmt.Println("Number of records", len(records))
	t := widget.NewTable(
		func() (int, int) {
			totalLength := len(records)
			return totalLength, 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			fmt.Println("populate", id.Row, id.Col)
			label := cell.(*widget.Label)
			switch id.Row {
			case 0:
				fmt.Println("case 0", id.Row, id.Col)
				switch id.Col {
				case 0:
					label.SetText("ID")
				case 1:
					label.SetText("Status")
				case 2:
					label.SetText("Email")
				case 3:
					label.SetText("Creator")
				}
			default:
				fmt.Println("default", id.Row, id.Col)
				switch id.Col {
				case 0:
					label.SetText(fmt.Sprintf("%d", id.Row))
				case 1:
					fmt.Println("STATUS for", id.Row, id.Col)
					label.SetText(records[id.Row]["status"])
				case 2:
					label.SetText(records[id.Row]["email"])
				case 3:
					label.SetText(records[id.Row]["creator_email"])
				}
			}

		})
	t.SetColumnWidth(0, 102)

	return t

}
