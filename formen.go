package main

import "fmt"

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	// Hinweis: Es ist bereits eine Schleife vorgegeben.
	// Diese läuft für jede Zeile des Quadrats einmal, das gezeichnet werden soll.
	// Bisher schreibt sie nur einen Zeilenumbruch.
	// Schreiben Sie nun eine weitere Schleife, die vor dem abschließenden fmt.Println()
	// die entsprechende Anzahl an '*' ausgibt. Verwenden Sie dafür fmt.Print().
	for i := 0; i < n; i++ {
		// TODO: Innere Schleife schreiben.
		fmt.Println()
	}
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	// Hinweis: Kopieren Sie sich die Funktion PrintSquare() und erweitern Sie sie um
	// eine Prüfung (mittels if), mit der Sie herausfinden, ob Sie gerade am Rand sind.
	// Am Rand befinden Sie sich, falls i oder j gleich 0 oder gleich n-1 sind.
	for i := 0; i < n; i++ {
		// TODO
		fmt.Println()
	}
}

// Erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	// Hinweis: Verallgemeinern Sie die Funktion PrintEmptySquare.
	// Das Prinzip ist exakt das selbe, nur dass Sie dieses Mal kein Leerzeichen
	// hartcodieren, sondern fill benutzen müssen.

	// TODO
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	// Hinweis: Diese Funktion arbeitet genau wie PrintSquare, nur dass die innere
	// Schleife nicht bis zur selben Zahl läuft wie die äußere.
}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	// Hinweis: Diese Funktion arbeitet ebenfalls ähnlich wie PrintSquare() und
	// PrintRectangle(). Hier muss die innere Schleife allerdings eine Variable
	// Anzahl an Wiederholungen haben.
}
