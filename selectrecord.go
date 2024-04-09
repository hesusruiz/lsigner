package main

import (
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/imroc/req/v3"
)

func makeTableOfCredentials() (fyne.CanvasObject, error) {

	var records []map[string]any
	client := req.C()
	response := client.Get("https://issuer.mycredential.eu/apiuser/retrievecredentials").Do()
	if response.Err != nil {
		return nil, fmt.Errorf("error: %s", response.Err)
	}
	if response.IsErrorState() {
		return nil, fmt.Errorf("HTTP status error %s", http.StatusText(response.StatusCode))
	}
	err := response.Into(&records)
	if err != nil {
		return nil, err
	}

	t := widget.NewTableWithHeaders(
		func() (int, int) { return len(records) + 1, 4 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row))
			default:
				label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
			}
		})
	t.SetColumnWidth(0, 102)
	t.SetRowHeight(2, 50)

	return t, nil

}
