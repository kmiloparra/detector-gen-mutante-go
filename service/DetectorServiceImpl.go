package service

import (
	"fmt"
	"github.com/kmiloparra/resource-mutant/constantes"
	"github.com/kmiloparra/resource-mutant/utilidades"
)

type DetectorGenMutante struct {
	secuenciasValidas   map[string]string
	secuenciasInvalidas map[string]string
}

func (dect DetectorGenMutante) IsMutant(dna []string) bool {
	dect.secuenciasInvalidas = utilidades.StringToMap(constantes.SECUENCIAS_INVALIDAS)
	dect.secuenciasValidas = utilidades.StringToMap(constantes.SECUENCIAS_GEN_MUTANTE_HASHMAP)
	contadorSecuenciasMutantes := contarSecuenciasGenomicasHorizontales(dna,dect.secuenciasValidas, dect.secuenciasInvalidas)
	fmt.Println("contadorSecuenciasMutantes Hori",contadorSecuenciasMutantes)
	if contadorSecuenciasMutantes < constantes.CANTIDAD_SECUENCIA_MUTANTE {
		contadorSecuenciasMutantes = contarSecuenciasGenomicasVerticales(dna,dect.secuenciasValidas, dect.secuenciasInvalidas)
		fmt.Println("contadorSecuenciasMutantes Vert",contadorSecuenciasMutantes)
		if contadorSecuenciasMutantes < constantes.CANTIDAD_SECUENCIA_MUTANTE {
			contadorSecuenciasMutantes = contarSecuenciasGenomicasDiagonales(dna,dect.secuenciasValidas, dect.secuenciasInvalidas)
			fmt.Println("contadorSecuenciasMutantes Diag",contadorSecuenciasMutantes)

		}
	}
	return contadorSecuenciasMutantes >= constantes.CANTIDAD_SECUENCIA_MUTANTE
}

func contarSecuenciasGenomicasHorizontales(dna []string,
	secuenciasValidas map[string]string,
	secuenciasInvalidas map[string]string) (contadorGenMutante uint8) {

	ejecutarBusquedaGenomicaCadenas(dna, secuenciasValidas, secuenciasInvalidas, &contadorGenMutante)
	return
}

func contarSecuenciasGenomicasVerticales(dna []string,
	secuenciasValidas map[string]string,
	secuenciasInvalidas map[string]string) (contadorGenMutante uint8) {
	dna = utilidades.PivotearMatrix(dna)
	ejecutarBusquedaGenomicaCadenas(dna, secuenciasValidas,
		secuenciasInvalidas, &contadorGenMutante)
	return
}

func contarSecuenciasGenomicasDiagonales(dna []string,
	secuenciasValidas map[string]string,
	secuenciasInvalidas map[string]string) (contadorGenMutante uint8) {
	dna = utilidades.ObtenerDiagonales(dna)
	ejecutarBusquedaGenomicaCadenas(dna, secuenciasValidas,
		secuenciasInvalidas, &contadorGenMutante)
	return
}

func ejecutarBusquedaGenomicaCadenas(dna []string,
	secuenciasValidas map[string]string,
	secuenciasInvalidas map[string]string,
	contadorGenMutante *uint8) {
	for _, cadena := range dna {
		*contadorGenMutante += utilidades.EncontrarIncidenciasHash(secuenciasValidas, secuenciasInvalidas, cadena)
		if *contadorGenMutante >=constantes.CANTIDAD_SECUENCIA_MUTANTE {
			return
		}
	}
}
