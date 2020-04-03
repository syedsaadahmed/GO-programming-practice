package main
import (
	"math"
	"fmt"
)

func value2long(x float64,z int) (lon float64) {
	lon = float64(x)/math.Exp2(float64(z))*360.0 - 180.0
	return lon
}

func value2lat(y float64,z int) (lat float64) {
	n := math.Pi - 2.0*math.Pi*float64(y)/math.Exp2(float64(z))
	lat = 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	return lat
}


func value2Tile(lat float64,lon float64,z int) (x float64, y float64) {
	x = float64(math.Floor((lon + 180.0) / 360.0 * (math.Exp2(float64(z)))))
	y = float64(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(z)))))
	return
}

func GetBoundingBox(x float64, y float64, z int) (north float64, south float64,east float64,west float64) {
	var xval,yval float64
	xval,yval = value2Tile(x,y,z)
	north = value2lat(yval,z)
	south = value2lat(yval+1,z)
	west = value2long(xval,z)
	east = value2long(xval+1,z)
	return north,south,west,east
}


func main() {
	var n,s,w,e float64
    n,s,w,e = GetBoundingBox(19.293750,50.681771,5)
    fmt.Printf("%f\n %f\n %f\n %f\n",s,w,n,e)
}
