package main

import "testing"

//The test function signature is TestXXX(t *testing.T) 
//with XXX as the name of the function we are testing.

//TestAdd testet die Add-Methode des Calculator-Types
func TestAdd(t *testing.T){
c:= Calculator{} //Erstellt ein neues Calculator-Objects
result  := c.Add(2,3) //Ruft die Mehtode auf
expected := 5   //Das erwartete Ergebnis der Berechnung

if result != expected {
	t.Errorf("Add(2,3)= %d; want %d", result,expected)
}
}

//TestSub testet die Subtraction-Methode des Calculator-Types
func TestSub(t *testing.T){
	c := Calculator{} 
	result := c.Sub(5,3)
	expected :=2
	if result!= expected {
        t.Errorf("Sub(5,3)= %d; want %de", result,expected)
    }
}

//TestMultiply  testet die Multiply-Methode des Calculator-Types
func TestMuliply(t *testing.T){
	c := Calculator{} 
	result := c.Multiply(5,3)
	expected :=15
	if result!= expected {
        t.Errorf("Multiply(5,3)= %d; want %d", result,expected)
    }
}
//TestDivide testet die Division-Methode des Calculator-Types
func TestDivide(t *testing.T){
	c := Calculator{} 
	result,err := c.Divide(6,3)
	expected :=2
	if err!= nil {
        t.Errorf("Divide(6,3) returned an error: %s", err.Error())
    }
	if result != expected {
		t.Errorf("Divide(6,3) = %d; want %d", result, expected)
	}
}

//calculator definiert eien einfachen Taschenrechner
type Calculator struct{}
func (c Calculator) Add(a, b int) int {
return a + b 
}
func (c Calculator) Sub(a, b int) int {
	return a - b 
}
func (c Calculator) Multiply(a, b int) int {
	return a * b 
}
func (c Calculator) Divide(a, b int) (int,error) {
	if b == 0 {
        return 0, &DivideByZeroError{

		}
    }
    return a/b,nil
	
}
//DivideByZero definiert einen benutzerdefinierten fehler
//der zurückgegeb wird, wenn die Division durch Null auftritt.
type DivideByZeroError struct{}

//DieseFunction gibt eine Fehlermelödung zurück wenn die Division 
//durch Nullauftritt
func (e *DivideByZeroError) Error() string {
return "Division by Null"
}




