package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Affiche le texte de résultat en console
// Affiche les statistiques
// - Nombre d'occurences du mot
// - Numéro de lignes où le mot a été trouvé

// FindReplaceInFile remplace l'ancien mot par le nouveau
func FindReplaceInFile(src, dst, old, new string) (occ int, lines []int, err error) {
	// gérér l'erreur d'ouverture de fichier
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := ProcessLigne(scanner.Text(), old, new) // t contient une ligne
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Println(res)
		lineIdx++
	}
	// writer := bufio.NewWriter(dstFile)
	// defer writer.Flush()
	// fmt.Fprintln(writer, "Texte d'une ligne")
	return occ, lines, err
}

// ProcessLigne remplace l'ancien mot par le nouveau
func ProcessLigne(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func main() {

	fichier := "fichier.txt"
	destFile := "newFile.txt"
	oldMot := "ancien"
	newMot := "nouveau"

	occurence, numLignes, err := FindReplaceInFile(fichier, destFile, oldMot, newMot)
	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
	}

	fmt.Println("__ Summary __")
	fmt.Printf("Number of occurences of Go: %d\n", occurence)
	fmt.Printf("Numner of lines: %d\n", len(numLignes))
	fmt.Printf("Lines: %v", numLignes)
	fmt.Println("__ End of summary __")

}
