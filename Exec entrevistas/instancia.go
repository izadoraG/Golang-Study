package main

import "fmt"

func CriaInstancias(Cpus []int) int {

	instancia := 4
	result := 0

	for _, cpu := range Cpus {
		if cpu > 80 {
			result = instancia * 2
		} else if cpu < 40 {
			result = instancia / 2
		}
	}
	return result

}

func main() {
	Cpus := []int{70, 30, 20, 45, 67, 81}
	instanciaAtualizadas := CriaInstancias(Cpus)
	fmt.Println(instanciaAtualizadas)
}
