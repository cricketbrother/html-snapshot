package html2png

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	// Check if Node.js is installed
	checkNodeJS()
	// Check if NPM is installed
	checkNPM()
	// Check if Puppeteer is installed
	checkPuppeteer()
	// Check if screenshotPage.mjs exists
	checkScreenshotPage()
	// Check if screenshotElement.mjs exists
	checkScreenshotElement()
	// Check if pdfPage.mjs exists
	checkPdfPage()
}

// checkNodeJS checks if Node.js is installed.
func checkNodeJS() {
	_, err := exec.LookPath("node")
	if err != nil {
		fmt.Println("Node.js not found. Please install Node.js and configure the PATH environment variable.")
		os.Exit(1)
	}
}

// checkNPM checks if NPM is installed.
func checkNPM() {
	_, err := exec.LookPath("npm")
	if err != nil {
		fmt.Println("NPM not found. Please install Node.js and configure the PATH environment variable.")
		os.Exit(1)
	}
}

// checkPuppeteer checks if Puppeteer is installed.
func checkPuppeteer() {
	_, err := os.Stat("pptr/node_modules/puppeteer/lib/esm/puppeteer/puppeteer.js")
	if err != nil {
		err = os.MkdirAll("pptr", 0755)
		if err != nil {
			fmt.Println("Failed to create pptr directory.")
			os.Exit(1)
		}

		cmd := exec.Command("npm", "install", "puppeteer")
		cmd.Dir = "pptr"
		err = cmd.Run()
		if err != nil {
			fmt.Println("Failed to install Puppeteer.")
			os.Exit(1)
		}
	}
}

// checkScreenshotPage checks if screenshotPage.mjs exists.
func checkScreenshotPage() {
	_, err := os.Stat("pptr/screenshotPage.mjs")
	if err != nil {
		f, err := os.Create("pptr/screenshotPage.mjs")
		if err != nil {
			fmt.Println("Failed to create screenshotPage.mjs.")
			os.Exit(1)
		}

		_, err = f.WriteString(`const argv = process.argv;
let src, dst, width, height;
if (argv.length === 6) {
	src = argv[2];
	dst = argv[3];
	width = argv[4];
	height = argv[5];
} else {
	console.log('Usage: \n> node screenshotPage.mjs src dst width height\n');
	process.exit(1);
}

import puppeteer from './node_modules/puppeteer/lib/esm/puppeteer/puppeteer.js';
const browser = await puppeteer.launch();
const page = await browser.newPage();
await page.setViewport({ width: parseInt(width), height: parseInt(height) });
await page.goto(src);
await page.screenshot({ path: dst });
await browser.close();`)
		if err != nil {
			fmt.Println("Failed to write screenshotPage.mjs.")
			os.Exit(1)
		}
	}
}

// checkScreenshotElement checks if screenshotElement.mjs exists.
func checkScreenshotElement() {
	_, err := os.Stat("pptr/screenshotElement.mjs")
	if err != nil {
		f, err := os.Create("pptr/screenshotElement.mjs")
		if err != nil {
			fmt.Println("Failed to create screenshotElement.mjs.")
			os.Exit(1)
		}

		_, err = f.WriteString(`const argv = process.argv;
let src, element, dst, width, height;
if (argv.length === 7) {
	src = argv[2];
	element = argv[3];
	dst = argv[4];
	width = argv[5];
	height = argv[6];
} else {
	console.log('Usage: \n> node screenshotElement.mjs src element dst width height\n');
	process.exit(1);
}

import puppeteer from './node_modules/puppeteer/lib/esm/puppeteer/puppeteer.js';
const browser = await puppeteer.launch();
const page = await browser.newPage();
await page.setViewport({ width: parseInt(width), height: parseInt(height) });
await page.goto(src);
const elementHandle = await page.waitForSelector(element);
await elementHandle.screenshot({ path: dst });
await browser.close();`)
		if err != nil {
			fmt.Println("Failed to write screenshotElement.mjs.")
			os.Exit(1)
		}
	}
}

// checkPdfPage checks if pdfPage.mjs exists.
func checkPdfPage() {
	_, err := os.Stat("pptr/pdfPage.mjs")
	if err != nil {
		f, err := os.Create("pptr/pdfPage.mjs")
		if err != nil {
			fmt.Println("Failed to create pdfPage.mjs.")
			os.Exit(1)
		}

		_, err = f.WriteString(`const argv = process.argv;
let src, dst, width, height;
if (argv.length === 4) {
	src = argv[2];
	dst = argv[3];
} else {
	console.log('Usage: \n> node pdfPage.mjs src dst\n');
	process.exit(1);
}

import puppeteer from './node_modules/puppeteer/lib/esm/puppeteer/puppeteer.js';
const browser = await puppeteer.launch();
const page = await browser.newPage();
await page.goto(src);
await page.pdf({ path: dst });
await browser.close();`)
		if err != nil {
			fmt.Println("Failed to write pdfPage.mjs.")
			os.Exit(1)
		}
	}
}
