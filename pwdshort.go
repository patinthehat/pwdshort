/*
 * pwdshort -- abbreviates the CWD in your bash $PS1 prompt
 * v1.1.0 2015-07-15_01
 *
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"path"
)

type PathReplacement struct {
	Pathname string
	Replacement string
}

type maxDisplayed struct {
	Prefix int
	Suffix int
}

type Configuration struct {
	MaxDisplayedParts maxDisplayed
	Replacements []PathReplacement
}


func main() {
	thisApp := path.Base(os.Args[0])
	thisPath := path.Dir(os.Args[0])

	file, _ := os.Open(thisPath+"/"+thisApp+".json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("error:", err)
	} 

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
				switch value {
				//replace /home/user with "~"
				case "home":
					//make sure it's the /home directory
					if i == 1 {
						replacement = "~"
						bInHomeDir = true
						bSkipNext = true
					}
				default:
					replacement = value
				}
			
			value := replacement
			displayedParts++

			if displayedParts > configuration.MaxDisplayedParts.Prefix && i < totalParts-configuration.MaxDisplayedParts.Suffix {
				if !displayedElipsis {
					newPathStrs = append(newPathStrs, "...")
					displayedElipsis = true
				}
			} else {
				newPathStrs = append(newPathStrs, value)	
			}
		}
	}
	
	//read pathname replacement values from config and do any required replacements
	for i := 0; i < len(newPathStrs); i++ {
		for j := 0; j < len(configuration.Replacements); j++ {
			if newPathStrs[i] == configuration.Replacements[j].Pathname {
				newPathStrs[i] = configuration.Replacements[j].Replacement 
			}
		}
	}	

	sPrefixSlash := "/"
	if bInHomeDir {
		sPrefixSlash = ""
	}

	fmt.Printf("%s%s", sPrefixSlash, strings.TrimSpace(strings.Join(newPathStrs, "/")))
}
