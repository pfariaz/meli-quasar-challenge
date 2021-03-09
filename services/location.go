package services

import (
	"math"

	"github.com/pfariaz/meli-quasar-challenge/lib"
)

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float64) (x, y float64) {
	// https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method

	if len(distances) < 3 {
		return 0, 0
	}

	//Kenobi info
	kenobiDistance := distances[0]
	var positionXKenobi float64 = SatellitesLocation[lib.KenobiSatelliteName].X
	var positionYKenobi float64 = SatellitesLocation[lib.KenobiSatelliteName].Y

	//Skywalker info
	skywalkerDistance := distances[1]
	var positionXSkywalker float64 = SatellitesLocation[lib.SkywalkerSatelliteName].X
	var positionYSkywalker float64 = SatellitesLocation[lib.SkywalkerSatelliteName].Y

	//Sato info
	satoDistance := distances[2]
	var positionXSato float64 = SatellitesLocation[lib.SatoSatelliteName].X
	var positionYSato float64 = SatellitesLocation[lib.SatoSatelliteName].Y

	x = ((((math.Pow(kenobiDistance, 2)-math.Pow(skywalkerDistance, 2))+(math.Pow(positionXSkywalker, 2)-math.Pow(positionXKenobi, 2))+(math.Pow(positionYSkywalker, 2)-math.Pow(positionYKenobi, 2)))*(2*positionYSato-2*positionYSkywalker) - ((math.Pow(skywalkerDistance, 2)-math.Pow(satoDistance, 2))+(math.Pow(positionXSato, 2)-math.Pow(positionXSkywalker, 2))+(math.Pow(positionYSato, 2)-math.Pow(positionYSkywalker, 2)))*(2*positionYSkywalker-2*positionYKenobi)) / ((2*positionXSkywalker-2*positionXSato)*(2*positionYSkywalker-2*positionYKenobi) - (2*positionXKenobi-2*positionXSkywalker)*(2*positionYSato-2*positionYSkywalker)))

	y = ((math.Pow(kenobiDistance, 2) - math.Pow(skywalkerDistance, 2)) + (math.Pow(positionXSkywalker, 2) - math.Pow(positionXKenobi, 2)) + (math.Pow(positionYSkywalker, 2) - math.Pow(positionYKenobi, 2)) + x*(2*positionXKenobi-2*positionXSkywalker)) / (2*positionYSkywalker - 2*positionYKenobi)

	return
}
