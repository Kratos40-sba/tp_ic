package main

import (
	"encoding/csv"
	"fmt"
	knn2 "github.com/Kratos40-sba/tp_ic/knn"
	"io"
	"os"
	"strconv"
)

func main() {
	var irisMat [][]string
	iris, err := os.Open("iris.data")
	if err != nil {
		panic(err)
	}
	defer iris.Close()
	reader := csv.NewReader(iris)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		irisMat = append(irisMat, record)
	}
	var X [][]float64
	var Y []string
	for _, data := range irisMat {

		var temp []float64
		for _, i := range data[:4] {
			parsedValue, err := strconv.ParseFloat(i, 64)
			if err != nil {
				panic(err)
			}
			temp = append(temp, parsedValue)
		}

		X = append(X, temp)

		Y = append(Y, data[4])

	}
	var (
		trainX [][]float64
		trainY []string
		testX  [][]float64
		testY  []string
	)
	for i := range X {
		if i%2 == 0 {
			trainX = append(trainX, X[i])
			trainY = append(trainY, Y[i])
		} else {
			testX = append(testX, X[i])
			testY = append(testY, Y[i])
		}
	}
	knn := knn2.Knn{}
	knn.K = 4
	knn.Fit(trainX, trainY)
	prediction := knn.Predict(testX)

	correct := 0
	for i, _ := range prediction {
		if prediction[i] == testY[i] {
			correct += 1
		}
	}
	fmt.Println(correct)
	fmt.Println(len(prediction))
	fmt.Println("Accuracy", float64(correct)/float64(len(prediction)))

}
