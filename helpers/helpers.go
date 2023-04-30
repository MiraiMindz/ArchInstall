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

	//"syscall"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Automates the process of running shell commands
func RunShellCommand(testMode, returnOutput bool, name string, args ...string) string {
	if testMode {
		x := []string{name}
		x = append(x, args...)
		fmt.Println(x)
		return ""
	} else {
		cmd := exec.Command(name, args...)
		cmd.Stdin = os.Stdin
		if returnOutput {
			out, err := cmd.Output()
			Check(err)
			return string(out)
		} else {
			cmd.Stdout = os.Stdout
			cmd.Run()
			return ""
		}
	}
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
	s := RunShellCommand(false, true, "stty", "size")
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
		Mask:  '*',
	}
	result, err := inputPrompt.Run()

	Check(err)
	return result
}

func ReadFileToStrSlice(file, split string) []string {
	readFile, err := os.ReadFile(file)
	Check(err)

	lines := strings.Split(string(readFile), split)
	return lines
}

func ReadLocaleFile(file, split string) []string {
	var lines []string
	readFile, err := os.ReadFile(file)
	Check(err)

	rawLines := strings.Split(string(readFile), split)
	for i, v := range rawLines {
		if i > 22 {
			strippedLines := strings.Trim(v, "#")
			if strings.Contains(strings.ToLower(strippedLines), "utf") || strings.Contains(strings.ToLower(strippedLines), "iso") {
				lines = append(lines, strippedLines)
			}
		}
	}
	return lines
}

// Find and replace a line in a file.
func ReplaceFileLine(file, line, replace string) interface{} {
	var lineFound = true

	lines := ReadFileToStrSlice(file, "\n")
	for i, currLine := range lines {
		if strings.Contains(currLine, line) {
			lines[i] = replace
		} else {
			lineFound = false
		}
	}
	output := strings.Join(lines, "\n")
	err := os.WriteFile(file, []byte(output), 0644)
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
func JsonAppender(file, attribute string, value interface{}) {
	var data []map[string]interface{}
	content, err := os.ReadFile(file)
	Check(err)
	json.Unmarshal(content, &data)
	new_data := &map[string]interface{}{
		attribute: value,
	}
	data = append(data, *new_data)
	dataBytes, err := json.Marshal(data)
	Check(err)

	err = os.WriteFile(file, dataBytes, 0644)
	Check(err)
}

// Updates data in a JSON file.
func JsonUpdater(file, attribute string, value interface{}, returnValue bool) interface{} {
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
func FindFiles(root, fileExtension, secondExtension string, onlyFileName bool) []string {
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
	promptui.IconSelect = promptui.Styler(promptui.FGBlack)(">")
	promptui.IconWarn = promptui.Styler(promptui.FGYellow)("!")
	promptui.IconBad = promptui.Styler(promptui.FGRed)("X")
	promptui.IconGood = promptui.Styler(promptui.FGGreen)("O")
	promptui.KeyNextDisplay = "v"
	promptui.KeyPrevDisplay = "^"
	promptui.KeyBackwardDisplay = "<"
	promptui.KeyForwardDisplay = ">"

	helpStr := promptui.Styler(promptui.FGBlack)("Use the <up-arrow> and <down-arrow> to navigate and <ENTER> to confirm.")

	templates := &promptui.SelectTemplates{
		Help:     helpStr, // Use the arrow keys to navigate: ↓ ↑ → ←
		Active:   fmt.Sprintf("%s {{ . }}", promptui.IconSelect),
		Inactive: "  {{ . }}",
		Selected: fmt.Sprintf("%s {{ . }}", promptui.IconSelect),
	}

	prompt := promptui.Select{
		Label: fmt.Sprintf("%s [%s/%s]", promptString, PrintGreen("Yes"), PrintRed("No")),
		Items: []string{PrintGreen("Yes"), PrintRed("No")},
		//HideHelp: true,
		Templates: templates,
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
	promptui.IconSelect = promptui.Styler(promptui.FGBlack)(">")
	promptui.IconWarn = promptui.Styler(promptui.FGYellow)("!")
	promptui.IconBad = promptui.Styler(promptui.FGRed)("X")
	promptui.IconGood = promptui.Styler(promptui.FGGreen)("O")
	promptui.KeyNextDisplay = "v"
	promptui.KeyPrevDisplay = "^"
	promptui.KeyBackwardDisplay = "<"
	promptui.KeyForwardDisplay = ">"
	helpStr := promptui.Styler(promptui.FGBlack)("Use the <up-arrow> and <down-arrow> to navigate and <ENTER> to confirm.")

	templates := &promptui.SelectTemplates{
		Help: helpStr, // Use the arrow keys to navigate: ↓ ↑ → ←
	}
	prompt := promptui.Select{
		Label: promptString,
		Items: itemsList,
		//HideHelp: true,
		Templates: templates,
	}
	numRes, strRes, err := prompt.Run()
	Check(err)

	return numRes, strRes
}

// Prompts for a menu selection and provides a info.
func PromptSelectInfo(promptString string, itemsList []ItemInfo) (int, string) {
	promptui.IconSelect = promptui.Styler(promptui.FGBlack)(">")
	promptui.IconWarn = promptui.Styler(promptui.FGYellow)("!")
	promptui.IconBad = promptui.Styler(promptui.FGRed)("X")
	promptui.IconGood = promptui.Styler(promptui.FGGreen)("O")
	promptui.KeyNextDisplay = "v"
	promptui.KeyPrevDisplay = "^"
	promptui.KeyBackwardDisplay = "<"
	promptui.KeyForwardDisplay = ">"
	helpStr := promptui.Styler(promptui.FGBlack)("Use the <up-arrow> and <down-arrow> to navigate and <ENTER> to confirm.")

	templates := &promptui.SelectTemplates{
		Help:     helpStr, // Use the arrow keys to navigate: ↓ ↑ → ←
		Details:  promptui.Styler(promptui.FGBlack)("{{ .Info }}"),
		Active:   fmt.Sprintf("%s {{ .Item }}", promptui.IconSelect),
		Inactive: "  {{ .Item | black }}",
		Selected: fmt.Sprintf("%s {{ .Item }}", promptui.IconSelect),
	}

	prompt := promptui.Select{
		Label:     promptString,
		Items:     itemsList,
		Templates: templates,
	}
	numRes, _, err := prompt.Run()
	Check(err)

	return numRes, itemsList[numRes].Item
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
		Label:   inputPrompt,
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

func PacmanInstallPackages(pkgs ...string) {
	_PacManCommands := []string{"-S", "--noconfirm", "--needed"}
	pkgs = append(pkgs, _PacManCommands...)
	RunShellCommand(COMMANDS_TEST_MODE, false, "pacman", pkgs...)
}

func GetOccurrences(file, key string) int {
	readFile, err := os.ReadFile(file)
	var occurrences int
	Check(err)
	lines := strings.Split(string(readFile), "\n")
	for _, currLine := range lines {
		if strings.Contains(currLine, key) {
			occurrences++
		}
	}
	return occurrences
}

func GrepCommandOut(cmdOut, key string) interface{} {
	lines := strings.Split(cmdOut, "\n")
	for _, currLine := range lines {
		if strings.Contains(currLine, key) {
			return currLine
		}
	}
	return nil
}

func BulkPrint(strs []string) {
	for _, v := range strs {
		fmt.Println(v)
	}
}

func AurNixInstall(installer string, pkgs []string) {
	if strings.ToLower(installer) == "nix" {

	}
}

func PromptMultiSelect(promptText string, opts []string) []string {
	selections := []string{}
	prompt := &survey.MultiSelect{
		Message: promptText,
		Options: opts,
	}
	survey.AskOne(prompt, &selections)
	return selections
}

func CheckExistsInStringSlice(key string, slice []string) bool {
	for _, v := range slice {
		if strings.ToLower(key) == strings.ToLower(v) {
			return true
		}
	}
	return false
}

func InstallAurPkgWithoutHelper(pkgName string) {
	RunShellCommand(COMMANDS_TEST_MODE, false, "git", "clone", fmt.Sprintf("https://aur.archlinux.org/%s.git", pkgName))
	err := os.Chdir(fmt.Sprintf("./%s", pkgName))
	Check(err)
	RunShellCommand(COMMANDS_TEST_MODE, false, "makepkg", "-si", "--noconfirm")
}

// difference returns the elements in `a` that aren't in `b`.
func DifferenceBetweenSlices(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func EmptyFile(filePath string) {
	err := os.Truncate(filePath, 0)
	Check(err)
}

func AppendToFile(filePath, textToAppend string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s\n", textToAppend))
	Check(err)
}

func SudoExecute(testMode, returnOutput bool, name string, args ...string) string {
	if testMode {
		x := []string{name}
		x = append(x, args...)
		fmt.Println(x)
		return ""
	} else {
		sudoArgs := []string{"sudo", name}
		sudoArgs = append(sudoArgs, args...)
		sudoCommand := strings.Join(sudoArgs, " ")
		cmd := exec.Command("/bin/sh", "-c", sudoCommand)

		cmd.Stdin = os.Stdin
		if returnOutput {
			out, err := cmd.Output()
			Check(err)
			return string(out)
		} else {
			cmd.Stdout = os.Stdout
			cmd.Run()
			return ""
		}
	}
}

func RunShellCommandStdIn(testMode, returnOutput bool, stdIn, name string, args ...string) string {
	var cmd *exec.Cmd
	if testMode {
		x := []string{name}
		x = append(x, args...)
		fmt.Println(x)
		return ""
	} else {
		if len(args) != 0 {
			cmd = exec.Command(name, args...)
		} else {
			cmd = exec.Command(name)
		}
		cmd.Stdin = strings.NewReader(stdIn)
		if returnOutput {
			out, err := cmd.Output()
			Check(err)
			return string(out)
		} else {
			cmd.Stdout = os.Stdout
			cmd.Run()
			return ""
		}
	}
}

// CopyDir copies the content of src to dst. src should be a full path. SOURCE: https://stackoverflow.com/a/72246196
func CopyDir(src, dst string) error {

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// copy to this path
		outpath := filepath.Join(dst, strings.TrimPrefix(path, src))

		if info.IsDir() {
			os.MkdirAll(outpath, info.Mode())
			return nil // means recursive
		}

		// handle irregular files
		if !info.Mode().IsRegular() {
			switch info.Mode().Type() & os.ModeType {
			case os.ModeSymlink:
				link, err := os.Readlink(path)
				if err != nil {
					return err
				}
				return os.Symlink(link, outpath)
			}
			return nil
		}

		// copy contents of regular file efficiently

		// open input
		in, _ := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		// create output
		fh, err := os.Create(outpath)
		if err != nil {
			return err
		}
		defer fh.Close()

		// make it the same
		fh.Chmod(info.Mode())

		// copy content
		writtenBytes, err := io.Copy(fh, in)
		fmt.Printf("Copied %d bytes from %s to %s\n", writtenBytes, in.Name(), fh.Name())
		return err
	})
}

func GetEnvironmentVariables(variable string) string {
	// cmd := exec.Command("sh", "-c", fmt.Sprintf("echo $%s", variable))
	// out, err := cmd.Output()
	// Check(err)
	// return string(out)
	return os.Getenv(variable)
}

func IsCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("command -v %s", name))
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func ExportVariable(name, value string) {
	err := os.Setenv(name, value)
	Check(err)
}

// Automates the process of running shell commands
func RunShellCommandDir(testMode, returnOutput bool, dir, name string, args ...string) string {
	if testMode {
		x := []string{name}
		x = append(x, args...)
		fmt.Println(x)
		return ""
	} else {
		cmd := exec.Command(name, args...)
		cmd.Dir = dir
		cmd.Stdin = os.Stdin
		if returnOutput {
			out, err := cmd.Output()
			Check(err)
			return string(out)
		} else {
			cmd.Stdout = os.Stdout
			cmd.Run()
			return ""
		}
	}
}

func SudoExecuteDir(testMode, returnOutput bool, dir, name string, args ...string) string {
	if testMode {
		x := []string{name}
		x = append(x, args...)
		fmt.Println(x)
		return ""
	} else {
		sudoArgs := []string{"sudo", name}
		sudoArgs = append(sudoArgs, args...)
		sudoCommand := strings.Join(sudoArgs, " ")
		cmd := exec.Command("/bin/sh", "-c", sudoCommand)
		cmd.Dir = dir
		cmd.Stdin = os.Stdin
		if returnOutput {
			out, err := cmd.Output()
			Check(err)
			return string(out)
		} else {
			cmd.Stdout = os.Stdout
			cmd.Run()
			return ""
		}
	}
}

// func GetMapKeys(cMap map[interface{}]interface{}) []interface{} {
// 	// Use make() to create the slice for better performance
// 	mapKeys := make([]interface{}, 0, len(cMap))

// 	// We only need the keys
// 	for key := range cMap {
// 		mapKeys = append(mapKeys, key)
// 	}

// 	return mapKeys
// }

// func GetMapKeys(cMap map[string]map[string]string) []interface{} {
// 	mapKeys := make([]interface{}, 0, len(cMap))
// 	for key := range cMap {
// 		mapKeys = append(mapKeys, key)
// 	}

// 	return mapKeys
// }

func GetMapKeys(m interface{}) []string {
	keys := []string{}
	switch val := m.(type) {
	case map[string]interface{}:
		for k := range val {
			keys = append(keys, k)
			// nestedKeys := GetMapKeys(val[k])
			// keys = append(keys, nestedKeys...)
		}
	case map[string]map[string]string:
		for k := range val {
			keys = append(keys, k)
			// nestedKeys := GetMapKeys(val[k])
			// keys = append(keys, nestedKeys...)
		}
	case map[string]string:
		for k := range val {
			keys = append(keys, k)
			// nestedKeys := GetMapKeys(val[k])
			// keys = append(keys, nestedKeys...)
		}
	case []string:
		for _, k := range val {
			keys = append(keys, k)
			// nestedKeys := GetMapKeys(val[k])
			// keys = append(keys, nestedKeys...)
		}
	}
	return keys
}

func GetKeysWithParents(m map[string]interface{}, parentKeys []string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range m {
		keys := make([]string, len(parentKeys)+1)
		copy(keys, parentKeys)
		keys[len(parentKeys)] = k
		switch v := v.(type) {
		case map[string]interface{}:
			result[k] = keys
			nestedMap := GetKeysWithParents(v, keys)
			for nestedK, nestedV := range nestedMap {
				result[nestedK] = nestedV
			}
		default:
			result[k] = keys
		}
	}
	return result
}
