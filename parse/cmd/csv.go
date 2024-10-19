package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var convertToJSON bool

// csvCmd represents the CSV subcommand
var csvCmd = &cobra.Command{
	Use:   "csv [file path]",
	Short: "Display CSV file data or convert to JSON",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		if convertToJSON {
			convertCSVToJSON(filePath)
		} else {
			readCSV(filePath)
		}
	},
}

func init() {
	// Adding flag for CSV to JSON conversion
	csvCmd.Flags().BoolVarP(&convertToJSON, "to-json", "j", false, "Convert CSV to JSON format")
	parseCmd.AddCommand(csvCmd)
}

func readCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}
}

func convertCSVToJSON(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	headers := records[0]
	data := make([]map[string]string, 0)

	for _, record := range records[1:] {
		entry := make(map[string]string)
		for i, header := range headers {
			entry[header] = record[i]
		}
		data = append(data, entry)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error converting CSV to JSON:", err)
		return
	}

	outputFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".json"
	err = os.WriteFile(outputFilePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error saving JSON file:", err)
		return
	}

	// Display the converted JSON data
	fmt.Printf("CSV file converted to JSON and saved at: %s\n", outputFilePath)
	fmt.Println("Converted JSON Data:")
	fmt.Println(string(jsonData)) // Print the converted JSON data to the console
}
