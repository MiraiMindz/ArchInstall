package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Automates the process of running shell commands
func RunShellCommand( testMode bool, name string, args ...string) string {
	if testMode {
		x := []string{name}
		x = append(x, args...)
		fmt.Println(x)
		return ""
	} else {
		cmd := exec.Command(name, args...)
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		Check(err)
		return string(out)
	}
}

// Function that prints a [DEBUG] statement
// func DebugPrint(variable interface{}) {
// 	fmt.Printf("[DEBUG] %s\t=\t%s\n", variable, variable)
// }

// Used to show the current step in the installation
// func StepPrint(step string, stpCnt int8) {
// 	magenta := color.New(color.FgMagenta).SprintFunc()
// 	stpStr := ""
// 	if stpCnt < 10 {
// 		stpStr = fmt.Sprintf("0%v", stpCnt)
// 	} else {
// 		stpStr = fmt.Sprint(stpCnt)
// 	}
// 	base_step_string := "[" + stpStr + "]"
// 	colored_step_string := magenta(base_step_string)

// 	final_string := colored_step_string + " " + step
// 	fmt.Println(final_string)
// }

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
	s := RunShellCommand(false, "stty", "size")
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	height, err := strconv.Atoi(sArr[0])
	Check(err)

	width, err := strconv.Atoi(sArr[1])
	Check(err)
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

	Check(err)
	return result
}

// Find and replace a line in a file.
func ReplaceFileLine(file, line, replace string) interface{} {
	var lineFound = true
	readFile, err := os.ReadFile(file)
	Check(err)

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
	Check(err)
	if !lineFound {
		return lineFound
	} else {
		return nil
	}
}

func GetLine(file, key string) interface{} {
	readFile, err := os.ReadFile(file)
	Check(err)
	lines := strings.Split(string(readFile), "\n")
	for _, currLine := range lines {
		if strings.Contains(currLine, key) {
			return string(currLine)
		}
	}
	return nil
}

// Append data to a JSON file.
func JsonAppender(file string, attribute string, value interface{}) {
	var data []map[string]interface{}
	content, err := os.ReadFile(file)
	Check(err)
	json.Unmarshal(content, &data)
	new_data := &map[string]interface{} {
		attribute: value,
	}
	data = append(data, *new_data)
	dataBytes, err := json.Marshal(data)
	Check(err)

	err = os.WriteFile(file, dataBytes, 0644)
	Check(err)
}

// Updates data in a JSON file.
func JsonUpdater(file string, attribute string, value interface{}, returnValue bool) interface{} {
	file_stats, err := os.Stat(file)
	Check(err)

	content, err := os.ReadFile(file)
	Check(err)
	val := gjson.Get(string(content), attribute)
	if val.Exists() {
		sjson.Delete(string(content), attribute)
	}

	_val, err := sjson.Set(string(content), attribute, value)
	Check(err)

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

func JsonGetter(file, path string) string {
	content, err := os.ReadFile(file)
	Check(err)
	res := gjson.Get(string(content), path)
	return res.Str
}

// CURL implementation.
func CurlResponse(URL string) string {
	req, err := http.Get(URL)
	Check(err)
	defer req.Body.Close()
	res, err := http.DefaultClient.Do(req.Request)
	Check(err)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	Check(err)

	return string(body)
}

// Find any files recursively and tries with 2 distinct extensions.
func FindFiles(root, fileExtension string, secondExtension string, onlyFileName bool) []string {
	var _tempList []string
	filepath.WalkDir(root, func(path string, data fs.DirEntry, err error) error {
		Check(err)
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
        Label: fmt.Sprintf("%s [%s/%s]", promptString, PrintGreen("Yes"),  PrintRed("No")),
        Items: []string{PrintGreen("Yes"), PrintRed("No")},
		HideHelp: true,
    }
    _, result, err := prompt.Run()
    Check(err)
    return result == PrintGreen("Yes")
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
	Check(err)

	return numRes, strRes
}

// Prompts for a text input.
func InputPrompt(inputPrompt string) string {
	prompt := promptui.Prompt{
		Label: inputPrompt,
	}

	result, err := prompt.Run()
	Check(err)

	return result
}

func InputDefaultPrompt(inputPrompt, defaultValue string) string {
	prompt := promptui.Prompt{
		Label: inputPrompt,
		Default: defaultValue,
	}

	result, err := prompt.Run()
	Check(err)

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

	fmt.Printf("%s %s %s\n", PrintBlue(sectionText), PrintHiBlack(BoxVerticalChar()), stepText)
	fmt.Println(PrintHiBlack(hr_line))
}

func CopyFile(sourceFile, destination string) {
	sourceFileStatus, err := os.Stat(sourceFile)
	Check(err)
	if !sourceFileStatus.Mode().IsRegular() {
		panic(fmt.Errorf("%s is not a regular file", sourceFile))
	}
	source, err := os.Open(sourceFile)
	Check(err)
	defer source.Close()
	dest, err := os.Create(destination)
	Check(err)
	defer dest.Close()
	nBytes, err := io.Copy(dest, source)
	Check(err)
	fmt.Printf("Copied %s from %s to %s\n", ByteSizeConverter(uint64(nBytes)), sourceFile, destination)
}


func ExtractNumbers(text string) []int {
	re := regexp.MustCompile("[0-9]+")
	res := re.FindAllString(text, -1)
	var resList []int
	for _, i := range res {
		x, err := strconv.Atoi(i)
		Check(err)
		resList = append(resList, x)
	}
	return resList
}

func ExtractLetters(text string) []string {
	re := regexp.MustCompile("[a-z]+")
	res := re.FindAllString(strings.ToLower(text), -1)
	return res
}

func ParseSizeString(text string) string {
	var resString string
	numSize := ExtractNumbers(text)[0]
	typeSize := ExtractLetters(text)[0]

	// Edited because I'm not sure if binary notation works on sgdisk
	// So I'm using the example notation in the manpage.
	switch strings.ToLower(typeSize) {
	case "k", "ki", "kb", "kib":
		//resString = fmt.Sprintf("%dKiB", numSize)
		resString = fmt.Sprintf("%dM", numSize)
	case "m", "mi", "mb", "mib":
		//resString = fmt.Sprintf("%dMiB", numSize)
		resString = fmt.Sprintf("%dM", numSize)
	case "g", "gi", "gb", "gib":
		//resString = fmt.Sprintf("%dGiB", numSize)
		resString = fmt.Sprintf("%dG", numSize)
	}
	return resString
}

func ConvertToByte(text string) int {
	var resSize int
	numSize := ExtractNumbers(text)[0]
	typeSize := ExtractLetters(text)[0]

	switch strings.ToLower(typeSize) {
	case "k", "ki", "kb", "kib":
		resSize = numSize * KiB
	case "m", "mi", "mb", "mib":
		resSize = numSize * MiB
	case "g", "gi", "gb", "gib":
		resSize = numSize * GiB
	}

	return resSize
}

func RoundMultiple(number, multiple float64) float64 {
	return multiple * math.Round(number/multiple)
}


func IsValidHostname(hostname string) bool {
	if net.ParseIP(hostname) != nil {
		return false
	}

	if len(hostname) < 1 || len(hostname) > 255 {
		return false
	}

	if hostname[len(hostname)-1] == '.' {
		hostname = hostname[:len(hostname)-1]
		return false
	}

	parts := strings.Split(hostname, ".")
	for _, part := range parts {
		if len(part) < 1 || len(part) > 63 {
			return false
		}

		if part[0] == '-' || part[len(part)-1] == '-' {
			return false
		}

		for i := 0; i < len(part); i++ {
			if !((part[i] >= 'a' && part[i] <= 'z') || (part[i] >= 'A' && part[i] <= 'Z') || (part[i] >= '0' && part[i] <= '9') || part[i] == '-') {
				return false
			}
		}
	}

	return true
}


func IsValidLinuxUsername(username string) bool {
	if net.ParseIP(username) != nil {
		return false
	}

	if len(username) < 1 || len(username) > 32 {
		return false
	}

	if username[0] == '-' || username[0] == '.' {
		return false
	}

	for i := 0; i < len(username); i++ {
		if !((username[i] >= 'a' && username[i] <= 'z') || username[i] == '_' || (i > 0 && ((username[i] >= '0' && username[i] <= '9') || username[i] == '-' || username[i] == '_'))) {
			return false
		}
	}

	if username[len(username)-1] == '$' {
		return len(username) <= 31
	}

	return true
}

func CountDown(seconds int, prompt string) {
	for {
		if seconds <= 0 {
			break
		} else {
			if seconds > 1 {
				fmt.Println(PrintHiBlack(fmt.Sprintf("%s in %d Seconds...", prompt, seconds)))
			} else {
				fmt.Println(PrintHiBlack(fmt.Sprintf("%s in %d Second...", prompt, seconds)))
			}
			time.Sleep(1 * time.Second)
			seconds--
		}
	}
}

func WriteToFile(filePath, content string, permissions fs.FileMode) {
	err := os.WriteFile(filePath, []byte(content), permissions)
	Check(err)
	fmt.Printf("Written %d bytes to %s\n", len(content), filePath)
}
