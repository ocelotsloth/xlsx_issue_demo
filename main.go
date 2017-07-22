package main

import (
	"fmt"
	"os"

	"github.com/ocelotsloth/xlsx"
)

// Attempts to open an excel file and locate the header. The test fails if the
//  program pauses where indicated. Your computer's fans will likely start
//  spinning up if this happens. Personally, I just watch htop when I run the
//  program.
//
// NOTE: If you are using a remotely up to date version fo xlsx, you will have
//  swap out the statements at
//
// The commit which breaks this program is:
//    https://github.com/tealeg/xlsx/commit/e07ef57987388ad2510dac5fa2b5dde5945b1b31
func main() {
	file := "exampleFall2017.xlsx"

	fmt.Printf("Opening excel file: %s\n", file)
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Printf("File \"%s\" failed to open!\n", file)
		fmt.Println(err)
		os.Exit(1)
	}

	// This is where excecution essentially halts, never reaching the print.

	fmt.Printf("%s opened successfully.\n", file)

	// If you got here, all is well.
	fmt.Printf("Hey it worked, here's proof: %s is the name of sheet 0.\n\n", xlFile.Sheets[0].Name)
}
