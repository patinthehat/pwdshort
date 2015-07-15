package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//number of path parts to display at the front of the output
	var maxDisplayedPartsPrefix = 2
	//number of path parts to display at the end of the output
	var maxDisplayedPartsSuffix = 1

	dst := string(os.Getenv("PWD"))

	var sa []string
	sa = strings.Split(dst, "/")

	var bSkipNext = false
	var bInHomeDir = false
	var totalParts = len(sa)
	var displayedParts = 0
	var displayedElipsis = false
	var newPathStrs []string

	for i, value := range sa {
		var replacement string
		if bSkipNext {
			bSkipNext = false
			continue
		}

		if len(value) > 0 {
			//custom path-part replacements
			switch value {
			case "home":
				//make sure it's the /home directory
				if i == 1 {
					replacement = "~"
					bInHomeDir = true
					bSkipNext = true
				}
			case "Development":
				replacement = "Dev"
			case "projects":
				replacement = "proj"
			default:
				replacement = value
			}
			value := replacement

			displayedParts++

			if displayedParts > maxDisplayedPartsPrefix && i < totalParts-maxDisplayedPartsSuffix {
				if !displayedElipsis {
					newPathStrs = append(newPathStrs, "...")
					displayedElipsis = true
				}
			} else {
				newPathStrs = append(newPathStrs, value)
			}
		}
	}

	sPrefixSlash := "/"
	if bInHomeDir {
		sPrefixSlash = ""
	}

	fmt.Printf("%s%s", sPrefixSlash, strings.TrimSpace(strings.Join(newPathStrs, "/")))
}
