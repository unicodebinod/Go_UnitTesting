package main 

import "testing"

// Benchmark-Test für den Bubblesort-Algorithmus
func BenchmarkBubblesort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        arr := generateRandomArray(1000) // Array mit 1000 Elementen generieren
        bubblesort(arr)
    }
}

// Benchmark-Test für den Quicksort-Algorithmus
func BenchmarkQuicksort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        arr := generateRandomArray(1000) // Array mit 1000 Elementen generieren
        quicksort(arr)
    }
}