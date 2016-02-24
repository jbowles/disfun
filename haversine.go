package disfun

import (
	"math"
)

var piRadian = func() float64 { return math.Pi / 180.0 }
var sinPointSquare = func(val float64) float64 { return math.Pow(math.Sin(val), 2) }

// Haversine finds the shortest "as-the-crow-flies" distance between 2 points (φ,λ) on a sphere R... accounting for curvuture. In Navigation terms, for the example below, φ (phi) is latitude, λ (lambda) is longitude, R(adius) is earth’s radius (mean radius = 6,371km).
//
//	Example:
//	  a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
//	  c = 2 ⋅ atan2( √a, √(1−a) )
//	  d = R ⋅ c
//
//	References:
//  https://github.com/kellydunn/golang-geo.
//  http://www.codecodex.com/wiki/Calculate_Distance_Between_Two_Points_on_a_Globe
//  http://www.movable-type.co.uk/scripts/latlong.html
//  https://github.com/reddavis/Distance-Measures/blob/master/lib/distance_measures/haversine.rb
func Haversine(onePhi, oneLambda, twoPhi, twoLambda float64) (c float64) {
	dPhi := (twoPhi - onePhi) * piRadian()
	dLambda := (twoLambda - oneLambda) * piRadian()

	phi := onePhi * piRadian()
	lambda := twoPhi * piRadian()

	a1 := sinPointSquare(dPhi / 2)
	a2 := sinPointSquare(dLambda/2) * (math.Cos(phi) * math.Cos(lambda))

	a := a1 + a2
	c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return
}

// HaversinePoint accepts 2 sets of points with a radius, it multiplies radius by the result of the haversine function.
func HaversinePoint(pointA1, pointA2, pointB1, pointB2, radius float64) float64 {
	return (radius * Haversine(pointA1, pointA2, pointB1, pointB2))
}

// HaversineLatLon accepts 2 sets of lat/lon, it multiplies the EARTH_RADIUS by the result of the haversine function.
func HaversineLatLon(lat1, lon1, lat2, lon2 float64) float64 {
	return (EarthRadius * Haversine(lat1, lon1, lat2, lon2))
}
