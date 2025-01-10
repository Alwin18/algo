package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Alwin18/algo/cmd/flags"
	"github.com/spf13/cobra"
)

type TemplateData struct {
	BasePath string
}

func createFileFromTemplate(filePath, templatePath string, data TemplateData) error {
	updatedPath := strings.ReplaceAll(templatePath, basePath+"/", "")
	// Parse template dari file
	tmpl, err := template.ParseFiles(updatedPath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Buat file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Eksekusi template dan tulis ke file
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func generateTemplatePath(currentPath, name string) string {
	fileName := filepath.Base(name)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	// Tentukan folder template berdasarkan lokasi file
	dirName := filepath.Dir(currentPath)

	// Path relatif dari `internal/` ke `cmd/templates/`
	relativePath := strings.TrimPrefix(dirName, "internal/")
	relativePath = filepath.ToSlash(relativePath) // Normalize path to forward slashes

	// Gabungkan path template
	templatePath := filepath.Join("cmd/templates", relativePath, fileNameWithoutExt+".tmpl")
	return templatePath
}

func createFolderStructure(pathFolder string, structure map[string]interface{}) error {
	for name, content := range structure {
		currentPath := filepath.Join(pathFolder, name)
		switch content := content.(type) {
		case nil: // File kosong
			// Buat direktori induk jika belum ada
			if err := os.MkdirAll(filepath.Dir(currentPath), os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory for file %s: %w", currentPath, err)
			}
			// Buat file kosong
			if _, err := os.Create(currentPath); err != nil {
				return fmt.Errorf("failed to create file %s: %w", currentPath, err)
			}
		case string:
			if content == "template" {
				// Gunakan template untuk membuat file
				templatePath := generateTemplatePath(currentPath, name)
				if err := createFileFromTemplate(currentPath, templatePath, TemplateData{BasePath: basePath}); err != nil {
					return err
				}
			} else {
				// Buat file kosong
				if err := os.WriteFile(currentPath, []byte{}, 0644); err != nil {
					return fmt.Errorf("failed to create file %s: %w", currentPath, err)
				}
			}
		case map[string]interface{}: // Folder
			// Buat folder
			if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create folder %s: %w", currentPath, err)
			}
			// Rekursif untuk substruktur
			if err := createFolderStructure(currentPath, content); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid structure for %s", name)
		}
	}

	return nil
}

// Inisialisasi Go module di basePath
func initializeGoMod(basePath string) error {
	// Navigasi ke basePath
	err := os.Chdir(basePath)
	if err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	// Jalankan perintah `go mod init`
	cmd := exec.Command("go", "mod", "init", basePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize go.mod: %v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Add flags command
	createCmd.Flags().StringVarP(&basePath, "path", "p", "./new-project", "Base path for the project")
}

var basePath string
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Algo is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Pastikan base path ada
		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create base directory %s: %w", basePath, err)
		}

		// Buat struktur folder
		if err := createFolderStructure(basePath, flags.Structure); err != nil {
			return err
		}

		// Inisialisasi go.mod
		if err := initializeGoMod(basePath); err != nil {
			return fmt.Errorf("error initializing go.mod: %v", err)
		}

		fmt.Println("Folder structure generated successfully at", basePath)
		return nil
	},
}
