package providers

import (
	"bytes"
	"strings"

	"github.com/igorscandido/ssromanos/graph/model"
)

func SearchHighestRoman(entrada string) model.Response {

	var romanos string = ""

	var acceptedCharacters []byte = []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}

	var numerosRomanos map[byte]int = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var arrNumerosRomanos []string
	var arrValoresDecimais []int

	var maiorValor int = 0
	var maiorValorRomano string = ""

	for x := 0; x < len(entrada); x++ {

		if bytes.Contains(acceptedCharacters, []byte{entrada[x]}) {

			romanos += string(entrada[x])

		} else {

			if len(romanos) != 0 {
				if romanos[len(romanos)-1] != '0' {
					romanos += "0"
				}
			}

		}

	}

	arrNumerosRomanos = strings.Split(romanos, "0")

	for x := 0; x < len(arrNumerosRomanos); x++ {

		valor := 0
		aux := arrNumerosRomanos[x]

		for x := 0; x < len(aux); x++ {
			c := aux[x]
			vc := numerosRomanos[c]
			if x < len(aux)-1 {
				cnext := aux[x+1]
				vcnext := numerosRomanos[cnext]
				if vc < vcnext {
					valor += vcnext - vc
					x++
				} else {
					valor += vc
				}
			} else {
				valor += vc
			}
		}

		arrValoresDecimais = append(arrValoresDecimais, valor)
		valor = 0
		aux = ""
	}

	for x := 0; x < len(arrValoresDecimais); x++ {

		if arrValoresDecimais[x] > maiorValor {
			maiorValor = arrValoresDecimais[x]
			maiorValorRomano = arrNumerosRomanos[x]
		}

	}

	return model.Response{
		Number: maiorValor,
		Value:  maiorValorRomano,
	}

}
