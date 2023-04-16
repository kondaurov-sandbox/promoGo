package gen

import (
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"os"
	"time"
)

func generatePromotionCSV() string {

	id := uuid.NewV4().String()
	value := rand.Float64() * 100
	format := "2006-01-02 15:04:05 -0700 MST"
	timestamp := time.Now().Add(-time.Duration(rand.Intn(365*24)) * time.Hour).Format(format)

	return fmt.Sprintf("%s,%.6f,%s", id, value, timestamp)

}

func GeneratePromotionsCsvFile(filename string, n int) {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < n; i++ {
		line := generatePromotionCSV()
		fmt.Fprintln(writer, line)
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	fmt.Println("File written successfully.")
}
