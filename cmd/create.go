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
	// Misalnya, jika file berada di internal/domain, maka cari template di templates/domain
	dirName := filepath.Dir(currentPath)

	// Gunakan strings.TrimPrefix untuk mendapatkan path relatif setelah "internal/"
	relativePath := strings.TrimPrefix(dirName, "internal/")
	// Bangun path template dengan folder relatif
	templatePath := filepath.Join("cmd/templates", relativePath, fileNameWithoutExt+".tmpl")
	return templatePath
}

func createFolderStructure(pathFolder string, structure map[string]interface{}) error {
	for name, content := range structure {
		currentPath := filepath.Join(pathFolder, name)
		switch content := content.(type) {
		case nil: // Create a file
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(currentPath), os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory for file %s: %w", currentPath, err)
			}
			// Create the file
			if _, err := os.Create(currentPath); err != nil {
				return fmt.Errorf("failed to create file %s: %w", currentPath, err)
			}
		case string:
			if content == "template" {
				// Jika konten adalah "template", gunakan template untuk membuat file
				templatePath := generateTemplatePath(currentPath, name)
				if err := createFileFromTemplate(currentPath, templatePath, TemplateData{basePath}); err != nil {
					return err
				}
			} else {
				// Jika konten adalah nil, buat file kosong
				if err := os.MkdirAll(filepath.Dir(currentPath), os.ModePerm); err != nil {
					return fmt.Errorf("failed to create directory for file %s: %w", currentPath, err)
				}
				// Buat file kosong
				if err := os.WriteFile(currentPath, []byte{}, 0644); err != nil {
					return fmt.Errorf("failed to create file %s: %w", currentPath, err)
				}
			}
		case map[string]interface{}: // Create a folder
			// Create the folder
			if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create folder %s: %w", currentPath, err)
			}
			// Recursively process substructure
			if err := createFolderStructure(currentPath, content); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid structure for %s", name)
		}
	}

	return nil
}

func initializeGoMod(basePath string) error {
	// Navigasi ke basePath
	err := os.Chdir(basePath)
	if err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	// Jalankan perintah `go mod init`
	cmd := exec.Command("go", "mod", "init", basePath) // Ganti dengan path modul Anda
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

	// Add flags comand
	createCmd.Flags().StringVarP(&basePath, "path", "p", "./new-project", "Base path for the project")
	fmt.Println("Creating folder structure...")
}

var basePath string
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Algo is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure the base path exists
		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create base directory %s: %w", basePath, err)
		}

		// Create the folder structure
		if err := createFolderStructure(basePath, flags.Structure); err != nil {
			return err
		}

		if err := initializeGoMod(basePath); err != nil {
			return fmt.Errorf("error initializing go.mod: %v", err)
		}

		fmt.Println("Folder structure generated successfully at", basePath)
		return nil
	},
}
