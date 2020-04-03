
package main
import (
	"math"
	"fmt"
)

func lon2xtile(lon float64,z int) (x int) {
	x = int(math.Floor((lon + 180.0) / 360.0 * (math.Exp2(float64(z)))))
	return x
}

func lat2ytile(lat float64,z int) (y int) {
	y = int(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(z)))))
	return y
}

func xval2long(x int,z int) (lon float64) {
	lon = float64(x)/math.Exp2(float64(z))*360.0 - 180.0
	return lon
}


func yval2lat(y int,z int) (lat float64) {
	n := math.Pi - 2.0*math.Pi*float64(y)/math.Exp2(float64(z))
	lat = 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	return lat
}


func GetBoundingBox(x int, y int, z int) (north float64, south float64,east float64,west float64) {
	// var xval,yval int
	// xval,yval = val2tile(x,y,z)
	north = yval2lat(y,z)
	south = yval2lat(y+1,z)
	east = xval2long(x+1,z)
	west = xval2long(x,z)
	return north,south,west,east
}


func main() {
	// var llong, llat float64
	// llong = val2long(70406,17)
	// llat = val2lat(42987,17)
	// fmt.Printf("%f\n %f\n",llong,llat)

	// var x,y int
	// x = lon2xtile(13.375854,17)
	// y = lat2ytile(52.517892,17)
	// fmt.Printf("%d\n %d\n",x,y)

	var n,s,e,w float64
    n,s,e,w = GetBoundingBox(70405,42987,17)
    fmt.Printf("%f\n %f\n %f\n %f\n",n,s,e,w)
}
