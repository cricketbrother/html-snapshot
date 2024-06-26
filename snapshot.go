package htmlsnapshot

import (
	"os/exec"
	"strconv"
)

// Converter is a struct that contains the necessary information to convert an HTML file to a PNG file.
//
// - NodePath is the path to the Node.js executable.
//
// - Src is the path to the HTML file.
//
// - Element is the CSS selector of the element to be converted to a PNG file. It is used in the MakeScreenshotElement function.
//
// - Dst is the path to the image file or the PDF file.
//
// - Width is the width of the PNG file. It is used in the MakeScreenshotPage and MakeScreenshotElement functions.
//
// - Height is the height of the PNG file. It is used in the MakeScreenshotPage and MakeScreenshotElement functions.
type Converter struct {
	NodePath string
	Src      string
	Element  string
	Dst      string
	Width    int
	Height   int
}

// MakeScreenshotPage is a function that converts an HTML file to an image file.
//
// - If the c.Width is not specified, the default width is 800. If the c.Height is not specified, the default height is 600.
//
// - If the c.NodePath is not specified, the default path is "node".
//
// - If the c.Src or c.Dst is not specified, the function will return an error.
func (c *Converter) MakeScreenshotPage() error {
	if c.NodePath == "" {
		c.NodePath = "node"
	}

	cmd := exec.Command(c.NodePath, "pptr/screenshotPage.mjs", c.Src, c.Dst, strconv.Itoa(c.Width), strconv.Itoa(c.Height))
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// MakeScreenshotElement is a function that converts an HTML file to an image file with a specific element.
//
// - If the c.Width is not specified, the default width is 800. If the c.Height is not specified, the default height is 600.
//
// - If the c.NodePath is not specified, the default path is "node".
//
// - If the c.Src or c.Dst is not specified, the function will return an error.
func (c *Converter) MakeScreenshotElement() error {
	if c.NodePath == "" {
		c.NodePath = "node"
	}

	cmd := exec.Command(c.NodePath, "pptr/screenshotElement.mjs", c.Src, c.Element, c.Dst, strconv.Itoa(c.Width), strconv.Itoa(c.Height))
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// MakePdfPage is a function that converts an HTML file to a PDF file.
//
// - If the c.NodePath is not specified, the default path is "node".
//
// - If the c.Src or c.Dst is not specified, the function will return an error.
func (c *Converter) MakePdfPage() error {
	if c.NodePath == "" {
		c.NodePath = "node"
	}

	cmd := exec.Command(c.NodePath, "pptr/pdfPage.mjs", c.Src, c.Dst)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
