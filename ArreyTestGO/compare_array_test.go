package main_test

import (
    "testing"
)
//TestCompareArrays testet die Funktion CompareArrays, um zu überprüfen,
//ob die beiden Arrays identisch sind


func TestCompareArrays(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}
 //Hier wird getestet ob 2 Arrays identisch sind
    if !CompareArrays(arr1, arr2){
		t.Errorf("Test  failed : die Arrays sollten identisch sein")
	}
    //verändern eines Elements im zweiten Arrays
	arr2[2] = 10

	//testen ob die beiden Arrays als unterschiedlich erkennen werden 
	if CompareArrays(arr1,arr2){
		t.Errorf("Test failed : die Arrays sollten unterschiedlich sein")
	}
}
//CompareArrays vergleicht zwei Arrays miteinander und gibt zurück,
//ob sie identisch sind
func CompareArrays(a,b []int) bool{
	if len(a)!= len(b) {
        return false
    }
	//vergleich der Elements beider Arrays
    for i := 0; i < len(a); i++ {
        if a[i]!= b[i] {
            return false
        }
    }
	//Wenn alle Elemente überstimmen sind die Arrays identisch
    return true
}





