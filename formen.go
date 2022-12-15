package main

import "fmt"

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	PrintCustomSquare(n, "*", "*")
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	PrintCustomSquare(n, "*", " ")
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print(border)
			} else {
				fmt.Print(fill)
			}
		}
		fmt.Println()
	}
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	PrintCustomRectangle(m, n, "*", "*")
}

// Hilfsfunktion: Diese macht die eigentliche Arbeit für die Square- und Rectangle-Funktionen.
func PrintCustomRectangle(m, n int, border, fill string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				fmt.Print(border)
			} else {
				fmt.Print(fill)
			}
		}
		fmt.Println()
	}
}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Main-Funktion mit Benutzereingabe.
// Wir verwenden Hilfsfunktionen (weiter unten definiert), damit es übersichtlich bleibt.
func main() {
	shape := GetShape()

	width := GetDimension("Breite")
	if shape == 3 {
		PrintTriangle(width)
		return
	}

	fill := GetSymbol("die Füllung")
	border := GetSymbol("den Rand")
	if shape == 2 {
		height := GetDimension("Höhe")
		PrintCustomRectangle(width, height, border, fill)
		return
	}
	PrintCustomSquare(width, border, fill)
}

// Fragt den Benutzer nach der Form und liefert eine entsprechende Zahl:
// 1: Quadrat
// 2: Rechteck
// 3: Dreieck
func GetShape() int {
	fmt.Println("Was soll gezeichnet werden?")
	fmt.Println("1: Quadrat")
	fmt.Println("2: Rechteck")
	fmt.Println("3: Dreieck")
	fmt.Print("Bitte geben Sie die entsprechende Zahl ein: ")

	var input int
	fmt.Scanln(&input)

	if input >= 1 && input <= 3 {
		return input
	}
	fmt.Println("Eingabe ungültig, bitte wiederholen.")
	fmt.Println("")
	return GetShape()
}

// Fragt den Benutzer nach eine Länge und liefert diese.
// Der Name der gewünschten Länge wird als Parameter erwartet.
func GetDimension(name string) int {
	fmt.Printf("Welche %s soll die Form haben?\n", name)
	fmt.Print("Bitte eine Zahl > 0 eingeben: ")

	var input int
	fmt.Scanln(&input)

	if input > 0 {
		return input
	}
	fmt.Println("Eingabe ungültig, bitte wiederholen.")
	fmt.Println("")
	return GetDimension(name)
}

// Fragt den Benutzer nach einem Zeichen.
// Die Bezeichnung wird als Parameter erwartet.
func GetSymbol(name string) string {
	fmt.Printf("Welches Zeichen soll für %s verwendet werden?\n", name)
	fmt.Print("Bitte einen String eingeben: ")

	var input string
	fmt.Scanln(&input)

	return input
}
