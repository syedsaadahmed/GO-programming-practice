package main
import (
	"math"
	"fmt"
)

type Tile struct {
	Z    float64
	X    float64
	Y    float64
	//Lat  float64
	//Long float64
}

// type Conversion interface {
// 	deg2num(t *Tile) (x float64, y float64)
// 	num2deg(t *Tile) (lat float64, long float64)
// }

// func (*Tile) Deg2num(t *Tile) (x float64, y float64) {
// 	x = float64(math.Floor((t.Long + 180.0) / 360.0 * (math.Exp2(float64(t.Z)))))
// 	y = float64(math.Floor((1.0 - math.Log(math.Tan(t.Lat*math.Pi/180.0)+1.0/math.Cos(t.Lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(t.Z)))))
// 	return
// }

func Num2deg(t *Tile) (lat float64, long float64) {
	n := math.Pi - 2.0*math.Pi*float64(t.Y)/math.Exp2(float64(t.Z))
	lat = 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	long = float64(t.X)/math.Exp2(float64(t.Z))*360.0 - 180.0
	fmt.Printf("%f\n",lat)
	fmt.Printf("%f\n",long)
	return lat, long
}

func GetBoundingBox(x float64, y float64, z float64) (north double, south double,east double,west double) {
	lat,long = Num2deg(x,y,z)
	lat_,long_ = Num2deg(x,y+1,z)
	_lat,l_ong = Num2deg(x,y,z)
	__lat,__long_ = Num2deg(x+1,y,z)
	// north = lat
	// south = lat_
	// east = _long
	// west = __long
	//fmt.Printf("%f\n%f\n%f\n%f\n",north,south,east,west)
	return
}


func main() {
    cls := Tile {X:52.517057 ,Y:13.377228 , Z:18,}
    GetBoundingBox(&cls)
}

