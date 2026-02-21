package main

import (
	"fmt"
	"os"

	"xlang/xlang"
	"xlang/xlang/stdlib"
)

func main() {
	// Read the script from the file
	script, err := os.ReadFile("test.x")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading script: %s\n", err)
		os.Exit(1)
	}

	// Create a new xlang script
	xs := xlang.NewScript(script)

	// Create a module map and add the standard library modules.
	moduleMap := xlang.NewModuleMap()
	for name, module := range stdlib.BuiltinModules {
		moduleMap.AddBuiltinModule(name, module)
	}

	// Also add the source modules from the standard library.
	for name, src := range stdlib.SourceModules {
		moduleMap.AddSourceModule(name, []byte(src))
	}

	xs.SetImports(moduleMap)

	// Enable file imports for custom modules.
	xs.EnableFileImport(true)

	// Compile the script
	compiled, err := xs.Compile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling script: %s\n", err)
		os.Exit(1)
	}

	// Run the script
	if err := compiled.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running script: %s\n", err)
		os.Exit(1)
	}
}
