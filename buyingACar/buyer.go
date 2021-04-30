package buyingACar

import (
	"math"
)

func NbMonths(startPriceOld, startPriceNew, savingPerMonth int, percentLossByMonth float64) [2]int {
	var months int
	var priceOld = float64(startPriceOld)
	var priceNew = float64(startPriceNew)
	savings := priceOld - priceNew

	for savings < 0 {
		months++

		if months % 2 == 0 { percentLossByMonth += 0.5 }
		priceOld -= priceOld * (percentLossByMonth / 100)
		priceNew -= priceNew * (percentLossByMonth / 100)

		savings = priceOld - priceNew
		savings += float64(savingPerMonth * months)
	}

	return [2]int{months, int(math.Round(savings))}
}
