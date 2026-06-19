package main

import (
	"fmt"
	"path"
)

var carriers = []string{"Allianz", "Banner", "Corebridge", "Equitable", "Pacific Life (First Heartland)", "John Hancock", "Lincoln", "MassMutual", "Mutual of Omaha", "National Life", "Nationwide", "New York Life", "Principal", "Protective", "Prudential", "Securian", "Symetra"}

func pdfImport(path string) {}
func importFile(fpath string) {
	incomMsg := `Unknown file type.  Please convert to:
	PDF`
	fileTypeI := path.Ext(fpath)
	fmt.Println(fileTypeI)
	switch fileTypeI {
	case ".pdf":
		pdfImport(fpath)
	default:
		fmt.Println(incomMsg)
	}
}
func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	
	importFile("C:\\Users\\Nick\\Downloads\\John Hancock 06.13.2026 $5,381.43.pdf")
}
