package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)
	for i := 0; i < length/2; i++ {
		podium[i], podium[length-1-1] = podium[length-1-1], podium[i]
	}
	return podium
}
