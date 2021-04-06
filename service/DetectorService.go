package service


type DetectorService interface {
	IsMutant( dna []string) bool
}
