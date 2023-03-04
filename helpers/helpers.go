package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/manifoldco/promptui"
)
// Function that prints a [DEBUG] statement
func DebugPrint(variable interface{}) {
	fmt.Printf("[DEBUG] %s\t=\t%s\n", variable, variable)
}

// Used to show the current step in the installation
func StepPrint(step string, stpCnt int8) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	stpStr := ""
	if stpCnt < 10 {
		stpStr = fmt.Sprintf("0%v", stpCnt)
	} else {
		stpStr = fmt.Sprint(stpCnt)
	}
	base_step_string := "[" + stpStr + "]"
	colored_step_string := magenta(base_step_string)

	final_string := colored_step_string + " " + step
	fmt.Println(final_string)
}

// Returns a string with the text centralized.
func CenterSprint(text, ghostText string) string {
	_, termW := GetTerminalSize()
	empty_space := 0
	if ghostText != "" {
		empty_space = ((termW - 4) - len(ghostText)) / 2
	} else {
		empty_space = ((termW - 4) - len(text)) / 2
	}

	empty_string := ""
	for i := 0; i < empty_space; i++ {
		empty_string += " "
	}
	final_string := empty_string + text + empty_string
	return final_string
}

// Uses the 'stty' UNIX command to get terminal size
func GetTerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	height, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return height, width
}

// Checks a error and panics
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Creates an empty file
func CreateEmptyFile(name string) {
	d := []byte("")
	Check(os.WriteFile(name, d, 0644))
}

// Returns the path of the current directory
func GetCurrDirPath() string {
	dir, err := os.Getwd()
	Check(err)
	return dir
}

// Checks if a file exists
func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

// Reads a password
func PromptReadPassword(prompt string) string {
	inputPrompt := promptui.Prompt{
		Label: prompt,
		Mask: '*',
	}
	result, err := inputPrompt.Run()

	if err != nil {
		panic(err)
	}
	return result
}

// Find and replace a line in a file.
func ReplaceFileLine(file, line, replace string) interface{} {
	var lineFound = true
	readFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(readFile), "\n")
	for i, currLine := range lines {
		if strings.Contains(currLine, line) {
			lines[i] = replace
		} else {
			lineFound = false
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(file, []byte(output), 0644)
	if err != nil {
		panic(err)
	}
	if !lineFound {
		return lineFound
	} else {
		return nil
	}
}

// Append data to a JSON file.
func JsonAppender(file string, attribute string, value interface{}) {
	var data []map[string]interface{}
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, &data)
	new_data := &map[string]interface{} {
		attribute: value,
	}
	data = append(data, *new_data)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(file, dataBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// Updates data in a JSON file.
func JsonUpdater(file string, attribute string, value interface{}, returnValue bool) interface{} {
	file_stats, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	val := gjson.Get(string(content), attribute)
	if val.Exists() {
		sjson.Delete(string(content), attribute)
	}

	_val, err := sjson.Set(string(content), attribute, value)
	if err != nil {
		panic(err)
	}

	if file_stats.Size() == 0 {
		os.WriteFile(file, []byte(_val), 0644)
	} else {
		lineFound := ReplaceFileLine(file, attribute, _val)
		if lineFound != nil && !lineFound.(bool) {
			JsonAppender(file, attribute, value)
			os.WriteFile(file, []byte(_val), 0644)
		}
	}

	if returnValue {
		return _val
	} else {
		return nil
	}
}

// CURL implementation.
func CurlResponse(URL string) string {
	req, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := http.DefaultClient.Do(req.Request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

// Find any files recursively and tries with 2 distinct extensions.
func FindFiles(root, fileExtension string, secondExtension string, onlyFileName bool) []string {
	var _tempList []string
	filepath.WalkDir(root, func(path string, data fs.DirEntry, err error) error {
		if err != nil {return err}
		if filepath.Ext(data.Name()) == fileExtension {
			if !onlyFileName {
				_tempList = append(_tempList, path)
			} else {
				_tempList = append(_tempList, strings.TrimSuffix(data.Name(), secondExtension))
			}
		}
		return nil
	})
	return _tempList
}

// Pretty byte formatter
func ByteSizeConverter(b uint64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}

// Checks ifs a string is a numeric value.
func IsNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

// Prompts for a yes/no decision.
func YesNo(promptString string) bool {
    prompt := promptui.Select{
        Label: fmt.Sprintf("%s [Yes/No]", promptString),
        Items: []string{"Yes", "No"},
		HideHelp: true,
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }
    return result == "Yes"
}

// Clears the console in UNIX like systems.
func ClearConsole() {
	fmt.Printf("\033c")
}

// Prompts for a menu selection.
func PromptSelect(promptString string, itemsList []string) (int, string) {
	prompt := promptui.Select{
		Label: promptString,
		Items: itemsList,
		HideHelp: true,
	}
	numRes, strRes, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return numRes, strRes
}

// Prompts for a text input.
func InputPrompt(inputPrompt string) string {
	prompt := promptui.Prompt{
		Label: inputPrompt,
	}

	result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}

// Prints the header seen through the script steps.
func PrintHeader(sectionText, stepText string) {
	_, termW := GetTerminalSize()
	var sectionSize int
	var stepSize int
	hr_line := ""

	for i := 0; i < len(sectionText); i++ {
		sectionSize++
	}

	for i := 0; i < len(stepText); i++ {
		stepSize++
	}

	for i := 0; i < (sectionSize + 1); i++ {
		hr_line += BoxHorizontalChar()
	}

	hr_line += BoxHorizontalUpChar()

	for i := 0; i < ((termW - sectionSize) - 5); i++ {
		hr_line += BoxHorizontalChar()
	}

	fmt.Printf("%s %s %s\n", sectionText, BoxVerticalChar(), stepText)
	fmt.Println(hr_line)
}
