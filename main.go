package detector_gen_mutante_go

import (
	"fmt"
	"github.com/kmiloparra/resource-mutant/constantes"
	"github.com/kmiloparra/resource-mutant/utilidades"
	"github.com/kmiloparra/resource-mutant/validaciones"
)

func main() {
	dna := []string{"GTCpAGTA", "TCGAGTAG", "CGAGTAGT", "CGAGTAGT", "pppppppp", "GAGTCGAT", "AGAGTCGT", "CGAGTAGT"}

	fmt.Println(validaciones.ValidacionFilaVacia(dna))
	fmt.Println(validaciones.ValidacionNxN(dna))
	fmt.Println(validaciones.ValidacionDominio(dna))
	fmt.Println(utilidades.PivotearMatrix(dna))
	fmt.Println(utilidades.ObtenerDiagonales(dna))
	fmt.Println(utilidades.EncontrarIncidenciasHash(utilidades.StringToMap(constantes.SECUENCIAS_GEN_MUTANTE_HASHMAP),
		utilidades.StringToMap(constantes.SECUENCIAS_INVALIDAS),
		"TTTTeAAAA"))

}
