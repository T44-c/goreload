
package main

import (
	"fmt"
	"io/fs"
	"os"
)

const outputPerm fs.FileMode = 0o644

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage : go run . <fichier_entrée> <fichier_sortie>")
		os.Exit(1)
	}

	if err := run(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, "erreur :", err)
		os.Exit(1)
	}
}

func run(inputPath, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("lecture de l'entrée : %w", err)
	}

	result := process(string(data))

	if err := os.WriteFile(outputPath, []byte(result), outputPerm); err != nil {
		return fmt.Errorf("écriture de la sortie : %w", err)
	}
	return nil
}
