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

var convertToCSV bool

// jsonCmd represents the JSON subcommand
var jsonCmd = &cobra.Command{
	Use:   "json [file path]",
	Short: "Display JSON file data or convert to CSV",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		if convertToCSV {
			convertJSONToCSV(filePath)
		} else {
			readJSON(filePath)
		}
	},
}

func init() {
	// Adding flag for JSON to CSV conversion
	jsonCmd.Flags().BoolVarP(&convertToCSV, "to-csv", "c", false, "Convert JSON to CSV format")
	parseCmd.AddCommand(jsonCmd)
}

func readJSON(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var data interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON file:", err)
		return
	}

	fmt.Printf("Data: %+v\n", data)
}

func convertJSONToCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var data []map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON file:", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("Empty JSON data")
		return
	}

	// Extract CSV headers from JSON keys
	headers := make([]string, 0, len(data[0]))
	for key := range data[0] {
		headers = append(headers, key)
	}

	// Prepare CSV file for writing
	outputFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".csv"
	csvFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		fmt.Println("Error writing CSV headers:", err)
		return
	}

	// Write records
	var csvData [][]string
	for _, record := range data {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = fmt.Sprintf("%v", record[header])
		}
		csvData = append(csvData, row)
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing CSV record:", err)
			return
		}
	}

	// Display the converted CSV data
	fmt.Printf("JSON file converted to CSV and saved at: %s\n", outputFilePath)
	fmt.Println("Converted CSV Data:")
	for _, row := range csvData {
		fmt.Println(row)
	}
}
