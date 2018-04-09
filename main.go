package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

func handleParseError(err error) bool {
	if err == io.EOF {
		return false
	} else if err, ok := err.(*csv.ParseError); ok && err.Err == csv.ErrFieldCount {
		return true
	} else if err != nil {
		return false
	}
	return true
}

func findPosition(key string, inArray []string) int {
	for i, val := range inArray {
		if val == key {
			return i
		}
	}
	return -1
}

func getFilePaths(dirPath string) ([]string, error) {
	filePaths := make([]string, 0)
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Println("Not enough arguments supplied", args)
		os.Exit(1)
	}
	log.Println("reading csv files")

	// get input csv files
	inputDir := args[1]
	inputFiles, err := getFilePaths(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	// setup writer
	outputFileName := args[2]
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// write default keys to output csv
	defaultKeys := []string{"linkedinProfile", "description", "imgUrl", "firstName", "lastName", "fullName", "subscribers", "company", "companyUrl", "jobTitle", "jobTenure", "jobDescription", "location", "mail", "phoneNumber", "twitter", "skill1", "skill2", "skill3", "school-name-1", "school-url-1", "school-degree-1", "school-degreeSpec-1", "school-dateRange-1", "school-name-2", "school-url-2", "school-degree-2", "school-degreeSpec-2", "school-dateRange-2", "school-name-3", "school-url-3", "school-degree-3", "school-degreeSpec-3", "school-dateRange-3", "school-name-4", "school-url-4", "school-degree-4", "school-degreeSpec-4", "school-dateRange-4", "school-name-5", "school-url-5", "school-degree-5", "school-degreeSpec-5", "school-dateRange-5", "school-name-6", "school-url-6", "school-degree-6", "school-degreeSpec-6", "school-dateRange-6", "previousJob-companyName-1", "previousJob-companyUrl-1", "previousJob-title-1", "previousJob-dateRange-1", "previousJob-location-1", "previousJob-description-1", "previousJob-companyName-2", "previousJob-companyUrl-2", "previousJob-title-2", "previousJob-dateRange-2", "previousJob-location-2", "previousJob-description-2", "previousJob-companyName-3", "previousJob-companyUrl-3", "previousJob-title-3", "previousJob-dateRange-3", "previousJob-location-3", "previousJob-description-3", "previousJob-companyName-4", "previousJob-companyUrl-4", "previousJob-title-4", "previousJob-dateRange-4", "previousJob-location-4", "previousJob-description-4", "previousJob-companyName-5", "previousJob-companyUrl-5", "previousJob-title-5", "previousJob-dateRange-5", "previousJob-location-5", "previousJob-description-5", "previousJob-companyName-6", "previousJob-companyUrl-6", "previousJob-title-6", "previousJob-dateRange-6", "previousJob-location-6", "previousJob-description-6", "previousJob-companyName-7", "previousJob-companyUrl-7", "previousJob-title-7", "previousJob-dateRange-7", "previousJob-location-7", "previousJob-description-7", "previousJob-companyName-8", "previousJob-companyUrl-8", "previousJob-title-8", "previousJob-dateRange-8", "previousJob-location-8", "previousJob-description-8", "previousJob-companyName-9", "previousJob-companyUrl-9", "previousJob-title-9", "previousJob-dateRange-9", "previousJob-location-9", "previousJob-description-9", "previousJob-companyName-10", "previousJob-companyUrl-10", "previousJob-title-10", "previousJob-dateRange-10", "previousJob-location-10", "previousJob-description-10", "previousJob-companyName-11", "previousJob-companyUrl-11", "previousJob-title-11", "previousJob-dateRange-11", "previousJob-location-11", "previousJob-description-11", "previousJob-companyName-12", "previousJob-companyUrl-12", "previousJob-title-12", "previousJob-dateRange-12", "previousJob-location-12", "previousJob-description-12", "previousJob-companyName-13", "previousJob-companyUrl-13", "previousJob-title-13", "previousJob-dateRange-13", "previousJob-location-13", "previousJob-description-13", "previousJob-companyName-14", "previousJob-companyUrl-14", "previousJob-title-14", "previousJob-dateRange-14", "previousJob-location-14", "previousJob-description-14", "previousJob-companyName-15", "previousJob-companyUrl-15", "previousJob-title-15", "previousJob-dateRange-15", "previousJob-location-15", "previousJob-description-15"}
	err = writer.Write(defaultKeys)
	if err != nil {
		log.Fatal(err)
	}

	for _, inputFile := range inputFiles {
		csvFile, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		reader := csv.NewReader(csvFile)
		var currentKeys []string
		count := 0
		for {
			if count == 0 {
				currentKeys, err = reader.Read()
				if err == io.EOF {
					break
				}
				toRun := handleParseError(err)
				if !toRun {
					log.Fatal(err)
				}
				count++
				continue
			}
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			toRun := handleParseError(err)
			if !toRun {
				log.Fatal(err)
			}
			formattedLine := make([]string, 0)
			for _, key := range defaultKeys {
				position := findPosition(key, currentKeys)
				if position == -1 {
					formattedLine = append(formattedLine, " ")
					continue
				}
				var value string
				if position > len(line)-1 {
					value = " "
				} else {
					value = line[position]
				}
				formattedLine = append(formattedLine, value)
			}
			err = writer.Write(formattedLine)
			if err != nil {
				log.Fatal(err)
			}
			count++
		}
		csvFile.Close()
	}
	log.Println("parsed/merged successfully, file saved in", outputFile.Name())
}
