package main

import (
	"fmt"
	"strconv"

	// "json-splitter/utils"
	"bufio"
	// "strconv"
	"os"
)

func main() {

	//demande du fichier à l'utilisateur
	var userInput string
	fmt.Println("Quel est le fichier que vous souhaitez traiter : ")
	fmt.Scan(&userInput)

	// Lire le fichier json à découper
	sourceFile, err := os.Open(userInput)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer sourceFile.Close()

	scanner := bufio.NewScanner(sourceFile)

	const linePerFile = 500 // Nb max de lignes par fichier
	nbLines := 0            // Compteur de ligne
	nbFiles := 1            // numéro du fichier en cours de traitement

	// Créer le premier fichier de destination
	destFile, err := os.Create("fichier_" + strconv.Itoa(nbFiles) + ".json")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)

	// Parcourir fichier source

	for scanner.Scan() {
		line := scanner.Text()
		nbLines++

		// écriture de la ligne

		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Erreur lors de l'écriture :", err)
			return
		}

		if nbLines >= linePerFile {
			err = writer.Flush()
			if err != nil {
				fmt.Println("Erreur lors de la finalisation de l'écriture :", err)
				return
			}
			destFile.Close()

			nbLines = 0
			nbFiles++

			// Créer un nouveau fichier
			destFile, err = os.Create("fichier_" + strconv.Itoa(nbFiles) + ".json")
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier :", err)
				return
			}
			writer = bufio.NewWriter(destFile)
		}
	}
	// Finaliser le dernier fichier
	err = writer.Flush()
	if err != nil {
		fmt.Println("Erreur lors de la finalisation de l'écriture :", err)
	}
	destFile.Close()

	// Vérifier les erreurs éventuelles de Scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture :", err)
	}

	fmt.Println("Traitement terminé, les fichiers ont été créés.")

}
