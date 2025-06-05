package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run fix-generated-code.go <file>")
		os.Exit(1)
	}

	filename := os.Args[1]
	
	dir := strings.TrimSuffix(filename, "/"+strings.Split(filename, "/")[len(strings.Split(filename, "/"))-1])
	cmd := exec.Command("go", "build", dir)
	output, _ := cmd.CombinedOutput()
	
	errorLines := strings.Split(string(output), "\n")
	lineTypeMappings := make(map[int]string)
	
	for _, line := range errorLines {
		errorPattern1 := regexp.MustCompile(`[^:]+:(\d+):\d+: cannot use.*as ([A-Z][a-zA-Z0-9_]*) value in assignment`)
		errorPattern2 := regexp.MustCompile(`[^:]+:(\d+):\d+: cannot use.*as \*([A-Z][a-zA-Z0-9_]*) value in assignment`)
		
		if match := errorPattern1.FindStringSubmatch(line); len(match) > 2 {
			lineNum, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}
			lineTypeMappings[lineNum] = match[2]
		} else if match := errorPattern2.FindStringSubmatch(line); len(match) > 2 {
			lineNum, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}
			lineTypeMappings[lineNum] = match[2]
		}
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	stringTypes := make(map[string]bool)
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if ident, ok := x.Type.(*ast.Ident); ok && ident.Name == "string" {
				stringTypes[x.Name.Name] = true
			}
		}
		return true
	})

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	fixCount := 0

	for i, line := range lines {
		lineNum := i + 1
		
		if correctType, hasError := lineTypeMappings[lineNum]; hasError {
			assignPattern := regexp.MustCompile(`([a-zA-Z_][a-zA-Z0-9_]*\.[A-Z][a-zA-Z0-9_]*)\s*=\s*"([^"]*)"`)
			if match := assignPattern.FindStringSubmatch(line); len(match) > 2 {
				assignment := match[1]
				value := match[2]
				
				replacement := assignment + ` = func() *` + correctType + ` { v := ` + correctType + `("` + value + `"); return &v }()`
				lines[i] = assignPattern.ReplaceAllString(line, replacement)
				fixCount++
			}
		}
	}

	fixed := strings.Join(lines, "\n")
	err = os.WriteFile(filename, []byte(fixed), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
	
	if fixCount > 0 {
		fmt.Printf("Applied %d fixes to %s\n", fixCount, filename)
	} else {
		fmt.Printf("No fixes needed for %s\n", filename)
	}
}