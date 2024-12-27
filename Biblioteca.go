package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func LetturaFile() {
	var studio string

	fmt.Println("Inserire il nome del File con l'estensione:")
	fmt.Scan(&studio)

	file, err := os.Open(studio)

	defer file.Close()
	if err != nil {
		fmt.Println("Errore, file non trovato, reinserire")
		return
	} else {
		libri := make([]byte, 1024)

		for {
			n, err2 := file.Read(libri)
			if err2 != nil {
				if err2 == io.EOF {
					break
				}
				fmt.Println("Errore lettura file: ", err2)
				return
			}
			fmt.Println(string(libri[:n]))
		}
	}
}

func RitiroLibro() {
	var libreria []string
	var studio, nome string

	fmt.Println("Inserire il nome del libro e il file con l'estensione: ")
	fmt.Scan(&nome, &studio)

	file, err := os.Open(studio)
	defer file.Close()

	if err != nil {
		fmt.Println("Errore nell'apertura del file: ", err)
	}

	filescanner := bufio.NewScanner(file)

	for filescanner.Scan() {
		libreria = append(libreria, filescanner.Text())
	}

	for i, libro := range libreria {
		if strings.Contains(libro, nome) {
			parts := strings.Fields(libro)
			numero, err := strconv.Atoi(string(libro[len(libro)-1]))

			if numero == 0 {
				fmt.Println("Non è presente il libro cercato")
			}
			if err != nil {
				fmt.Println("Errore nella richiesta del libro", err)
			}

			numero--

			parts[len(parts)-1] = strconv.Itoa(numero)
			libreria[i] = strings.Join(parts, " ")

			var err3 error

			_, err3 = file.WriteAt([]byte(libreria[i]), int64(i))

			if err3 != nil {
				fmt.Println("Errore nel prendere in prestito il libro")
			} else {
				fmt.Println("Libro preso in prestito con successo")
			}
		}
	}
}

func RestituzioneLibro() {
	var libreria []string
	var studio, nome string

	fmt.Println("Inserire il nome del libro e il file con l'estensione: ")
	fmt.Scan(&nome, &studio)

	file, err := os.OpenFile(studio, os.O_RDWR, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println("Errore nell'apertura del file: ", err)
	}

	filescanner := bufio.NewScanner(file)

	for filescanner.Scan() {
		libreria = append(libreria, filescanner.Text())
	}

	for i, libro := range libreria {
		if strings.Contains(libro, nome) {
			parts := strings.Fields(libro)
			numero, err := strconv.Atoi(string(libro[len(libro)-1]))

			if numero == 0 {
				fmt.Println("Non è presente il libro cercato")
			}
			if err != nil {
				fmt.Println("Errore nella richiesta del libro", err)
			}

			numero++

			parts[len(parts)-1] = strconv.Itoa(numero)
			libreria[i] = strings.Join(parts, " ")

			var err3 error

			for j, _ := range libreria {
				_, err3 = file.WriteAt([]byte(libreria[j]), int64(j))
			}

			if err3 != nil {
				fmt.Println("Errore nella restituzione")
			} else {
				fmt.Println("Libro restituito con successo")
			}
		}
	}
}

func main() {
	var y int

	fmt.Println("Inserire 1 se si vuole cercare un libro, 2 se si vuole ritirare una copia un libro, 3 per restituire: ")
	fmt.Scan(&y)

	for y < 1 || y > 3 {
		fmt.Println("Reinserire: ")
		fmt.Scan(&y)
	}

	for y != 0 {
		switch y {
		case 1:
			LetturaFile()
		case 2:
			RitiroLibro()
		case 3:
			RestituzioneLibro()
		}
		fmt.Println("Inserire 0 per uscrire:")
		fmt.Scan(&y)

		for y < 0 || y > 3 {
			fmt.Println("Reinserire: ")
			fmt.Scan(&y)
		}
	}
}
