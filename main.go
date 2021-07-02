package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Arr struct {
	Arr []int
}

func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
func handler(w http.ResponseWriter, r *http.Request) {
	var arr Arr
	_ = json.NewDecoder(r.Body).Decode(&arr)
	sortd_arr := bubbleSort(arr.Arr)

	fmt.Fprintf(w, "Sorted-v2 %v", sortd_arr)
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
