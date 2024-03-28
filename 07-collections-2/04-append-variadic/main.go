package main



func Merge(a []string, b []string) []string {
	putTogether := append(a,b...)
	return putTogether
}

func Remove(a []string, index int) []string {
    // Stelle sicher, dass index gültig ist
    if index < 0 || index >= len(a) {
        return a // Gebe den ursprünglichen Slice zurück, wenn der Index ungültig ist
    }
    return append(a[:index], a[index+1:]...) // Entferne das Element und gebe den neuen Slice zurück
}


func RemoveLast(a []string) []string {
    if len(a) == 0 {
        return a // Wenn der Slice leer ist, gibt einfach den ursprünglichen Slice zurück
    }
    lastElementIndex := len(a) - 1 // Korrektur hier, um den Index des letzten Elements zu bekommen
    return Remove(a, lastElementIndex) // Wiederverwende Remove, um das letzte Element zu entfernen
}

