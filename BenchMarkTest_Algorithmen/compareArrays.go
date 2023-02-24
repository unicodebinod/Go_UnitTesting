package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Funktion zum Generieren eines zufälligen Arrays von Ganzzahlen mit der Länge n
//It uses the "rand.Intn" function from the "math/rand" package to generate 
//random integers between 0 and 999 and stores them in an array.
func generateRandomArray(n int) []int {
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = rand.Intn(1000)
    }
    return arr
}

// Funktion zum Messen der Zeit, die ein Sortieralgorithmus benötigt, um ein Array zu sortieren
// It takes as arguments a sorting function and an array of integers to be sorted. 
//It uses the "time.Now" function from the "time" package to get the current time before and after sorting the array, 
//and calculates the difference between the two times using the "Sub" method of the "time.Duration" type.
func measureSortTime(sortFunc func([]int), arr []int) time.Duration {
    start := time.Now()
    sortFunc(arr)
    end := time.Now()
    return end.Sub(start)
}

// Bubblesort-Algorithmus
//it sorts an array of integers in ascending order by repeatedly swapping adjacent elements 
//that are in the wrong order.
func bubblesort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// Quicksort-Algorithmus
//This is the quicksort algorithm, which sorts an array of integers in ascending 
//order by recursively partitioning the array around a pivot element, and then sorting the two partitions. 
func quicksort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    pivot := arr[0]
    i, j := 1, len(arr)-1
    for i <= j {
        if arr[i] > pivot {
            arr[i], arr[j] = arr[j], arr[i]
            j--
        } else {
            i++
        }
    }
    arr[0], arr[i-1] = arr[i-1], arr[0]
    quicksort(arr[:i-1])
    quicksort(arr[i:])
}

func main() {
    // Array-Größe und Anzahl der Testläufe
    n := 10000
    numRuns := 10

    // Generieren des zufälligen Arrays
    arr := generateRandomArray(n)

    // Messen der Zeit für Bubblesort
    var bubbleSortTime time.Duration
    for i := 0; i < numRuns; i++ {
        arrCopy := make([]int, n)
        copy(arrCopy, arr)
        bubbleSortTime += measureSortTime(bubblesort, arrCopy)
    }
    bubbleSortTime /= time.Duration(numRuns)

    // Messen der Zeit für Quicksort
    var quickSortTime time.Duration
    for i := 0; i < numRuns; i++ {
        arrCopy := make([]int, n)
        copy(arrCopy, arr)
        quickSortTime += measureSortTime(quicksort, arrCopy)
    }
    quickSortTime /= time.Duration(numRuns)

    // Ausgabe der Ergebnisse
    fmt.Printf("Array-Größe: %d\n", n)
    fmt.Printf("Anzahl der Testläufe: %d\n", numRuns)
    fmt.Printf("Durchschnittliche Zeit für Bubblesort: %s\n", bubbleSortTime)
    fmt.Printf("Durchschnittliche Zeit für Quicksort: %s\n", quickSortTime)
}