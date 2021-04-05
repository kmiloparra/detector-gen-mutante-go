package mainMidle

import (
	"fmt"
	"github.com/kmiloparra/resource-mutant/utilidades"
	"github.com/kmiloparra/resource-mutant/validaciones"
)

func mainMidle() {
	dna := []string{"GTCpAGTA", "TCGAGTAG", "CGAGTAGT", "CGAGTAGT", "pppppppp", "GAGTCGAT", "AGAGTCGT", "CGAGTAGT"}

	fmt.Println(validaciones.ValidacionFilaVacia(dna))
	fmt.Println(validaciones.ValidacionNxN(dna))
	fmt.Println(validaciones.ValidacionDominio(dna))
	//fmt.Println(utilidades.PivotearMatrix(dna))
	//fmt.Println(utilidades.ObtenerDiagonales(dna))
	//fmt.Println(utilidades.EncontrarIncidenciasHash(utilidades.StringToMap(constantes.SECUENCIAS_GEN_MUTANTE_HASHMAP),
	//	utilidades.StringToMap(constantes.SECUENCIAS_INVALIDAS),
	//	"TTTTeAAAA"))
	fmt.Println(utilidades.Hola())
	fmt.Println(utilidades.Hola())

}
