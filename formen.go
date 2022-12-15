package main

import "fmt"

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print("*")
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
