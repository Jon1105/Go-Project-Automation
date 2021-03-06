package common

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var home, err = os.UserHomeDir()

// Documents path to documents folder where projects lie
var Documents string = filepath.Join(home, "OneDrive", "Documents")

// GoPath Path to Go projects to find current directory
var GoPath string = filepath.Join(Documents, "Programming", "Go", "src")

// CurrentPath Path representing location of this files parent's parent's
var CurrentPath string = filepath.Join(GoPath, "ProjectAutomation")

// Language struct to hold all variable info regarding each language with projects folder
type Language struct {
	Name, Path string
}

func w(path string) string {
	return filepath.Join(Documents, path)
}

// Classify used to determine the language from human inputted strings
func Classify(strLang string) (Language, error) {
	var lang Language
	var err error
	var low string = strings.ToLower(strLang)
	if low == "py" {
		lang = Language{"Python", w("Programming/Python/Projects")}

	} else if low == "cpp" || low == "c++" {
		lang = Language{"C++", w("Programming/C++/Projects")}
		
	} else if low == "node" || low == "js" {
		lang = Language{"Node.js", w("Programming/Node.js/Projects")}

	} else if low == "flutter" {
		lang = Language{"Flutter", w("Programming/Flutter/Projects")}

	} else if low == "go" {
		lang = Language{"Go", GoPath}

	} else if low == "ino" || low == "arduino" {
		lang = Language{"Arduino", w("Electronics/Arduino/Sketches")}

	} else if low == "rust" {
		lang = Language{"Rust", w("Programming/Rust/Projects")}

	} else if low == "workspace" {
		lang = Language{"Workspace", filepath.Join(home, ".vscode", "Workspaces")}
		
	} else {
		err = fmt.Errorf("common: string %q does not represent a valid language", strLang)

	}
	return lang, err
}

// Exists function to check whether path exists on the machine
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// OpenWithCode wrapper to open a folder with vscode
func OpenWithCode(path string) error {
	var cmd *exec.Cmd = exec.Command("code", path)
	return cmd.Run()
}

// Check wrapper for checking for errors and reporting them immediately
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
