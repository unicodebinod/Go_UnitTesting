package calculator

import "testing"

//Definieren eine Funktion die den Benchmarktest für die Addition durchführt
func BenchmarkAddition(b *testing.B) {
	//Führt den Schleifen b.N mal aus bis gewünschtes Ergebnis errreicht ist
    for i := 0; i < b.N; i++ {
        result := Add(2, 4)

		if result!= 6 {
			//Fehlerausgabe ,wenn Ergebnis falsch ist
            b.Error("Fehler: Das Ergebnis ist falsch")
        }
    }

    }
	//Definieren einer Funktion für die Subtraction
	func BenchmarkSubtraction (b * testing.B) {

		for i := 0; i < b.N; i++ {
			result := Sub(5, 2)
			if result!= 3 {
                //Fehlerausgabe,wenn Ergebnis falsch ist
                b.Error("Fehler: Das Ergebnis ist falsch")
            }
		}
}
//definieren eine Funktion die den Benchmarktest für die Multiplication dur
func BenchmarkMultiplication (b * testing.B) {

	for i := 0; i < b.N; i++ {
		result := Mul(5, 2)
		if result!= 10 {
			//Fehlerausgabe,wenn Ergebnis falsch ist
			b.Error("Fehler: Das Ergebnis ist falsch")
		}
	}
}
//Definieren eine Funktion die den Benchmarktest für die Division
func BenchmarkDivision (b * testing.B) {

	for i := 0; i < b.N; i++ {
		result := Div(10, 2)
		if result!= 5 {
			//Fehlerausgabe,wenn Ergebnis falsch ist
			b.Error("Fehler: Das Ergebnis ist falsch")
		}
	}
}