// package main
// import (
//     "math"
//     "fmt"
// )
// // Semi-axes of WGS-84 geoidal reference
// const (
//     WGS84A = 6378137.0
//     WGS84B = 6356752.3
// )
// // BoundingBox represents the geo-polygon that encompasses the given point and radius
// type BoundingBox struct {
//     LatMin float64
//     LatMax float64
//     LonMin float64
//     LonMax float64
// }
// // Convert a degree value to radians
// func deg2Rad(deg float64) float64 {
//     return math.Pi * deg / 180.0
// }
// // Convert a radian value to degrees
// func rad2Deg(rad float64) float64 {
//     return 180.0 * rad / math.Pi
// }
// // Get the Earth's radius in meters at a given latitude based on the WGS84 ellipsoid
// func getWgs84EarthRadius(lat float64) float64 {
//     an := WGS84A * WGS84A * math.Cos(lat)
//     bn := WGS84B * WGS84B * math.Sin(lat)
//     ad := WGS84A * math.Cos(lat)
//     bd := WGS84B * math.Sin(lat)
//     return math.Sqrt((an*an + bn*bn) / (ad*ad + bd*bd))
// }

// // GetBoundingBox returns a BoundingBox encompassing the given lat/long point and radius
// func GetBoundingBox(latDeg float64, longDeg float64, radiusKm float64) BoundingBox {
//     lat := deg2Rad(latDeg)
//     lon := deg2Rad(longDeg)
//     halfSide := 1000 * radiusKm

//     // Radius of Earth at given latitude
//     radius := getWgs84EarthRadius(lat)

//     pradius := radius * math.Cos(lat)

//     latMin := lat - halfSide/radius
//     latMax := lat + halfSide/radius
//     lonMin := lon - halfSide/pradius
//     lonMax := lon + halfSide/pradius

//     return BoundingBox{
//         LatMin: rad2Deg(latMin),
//         LatMax: rad2Deg(latMax),
//         LonMin: rad2Deg(lonMin),
//         LonMax: rad2Deg(lonMax),
//     }
// }


// // func (e BoundingBox) printBB(){
// // 	fmt.Printf("%f\n %f\n %f\n %f\n ", e.LatMin,e.LatMax,e.LonMin,e.LonMax)
// // }

// func main(){
// 	// e := BoundingBox.New()
// 	// e.GetBoundingBox(41.500000,12.067091,6)
// 	// e.printBB()
// 	fmt.Printf("%f\n %f\n %f\n %f\n ", GetBoundingBox(12.067091,41.500000,6))
// 	// xs := make([]float64, 4)
// 	// xs = GetBoundingBox(41.500000,12.067091,6)
// 	// fmt.Printf("%f",xs)
// }




package main
import (
	"math"
	"fmt"
)

type Tile struct {
	Z    float64
	X    float64
	Y    float64
	Lat  float64
	Long float64
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


func main() {
    cls := Tile {X:"52.517057" ,Y:"13.377228" , Z:"18",}

    Num2deg(&cls)
}








// func GetBoundingBox(x float64, y float64, z float64) (north double, south double,east double,west double) {
// 	lat,long = Num2deg(x,y,z)
// 	lat_,long_ = Num2deg(x,y+1,z)
// 	_lat,l_ong = Num2deg(x,y,z)
// 	__lat,__long_ = Num2deg(x+1,y,z)
// 	// north = lat
// 	// south = lat_
// 	// east = _long
// 	// west = __long
// 	//fmt.Printf("%f\n%f\n%f\n%f\n",north,south,east,west)
// 	return
// }

// func main(){
// 	GetBoundingBox(52.517057,13.377228,18)
// }


  class BoundingBox {
    double north;
    double south;
    double east;
    double west;   
  }
  BoundingBox tile2boundingBox(final int x, final int y, final int zoom) {
    BoundingBox bb = new BoundingBox();
    bb.north = tile2lat(y, zoom);
    bb.south = tile2lat(y + 1, zoom);
    bb.west = tile2lon(x, zoom);
    bb.east = tile2lon(x + 1, zoom);
    return bb;
  }

  static double tile2lon(int x, int z) {
     return x / Math.pow(2.0, z) * 360.0 - 180;
  }

  static double tile2lat(int y, int z) {
    double n = Math.PI - (2.0 * Math.PI * y) / Math.pow(2.0, z);
    return Math.toDegrees(Math.atan(Math.sinh(n)));
  }












sub getTileNumber {
    my ($lat,$lon,$zoom) = @_;
    my $xtile = int( ($lon+180)/360 * 2**$zoom ) ;
    my $ytile = int( (1 - log(tan(deg2rad($lat)) + sec(deg2rad($lat)))/pi)/2 * 2**$zoom ) ;
    return ($xtile, $ytile);
}

sub getLonLat {
	my ($xtile, $ytile, $zoom) = @_;
	my $n = 2 ** $zoom;
	my $lon_deg = $xtile / $n * 360.0 - 180.0;
	my $lat_deg = rad2deg(atan(sinh(pi * (1 - 2 * $ytile / $n))));
	return ($lon_deg, $lat_deg);
}

sub LonLat_to_bbox {
	my ($lat, $lon, $zoom) = @_;

	my $width = 425; my $height = 350;	# note: must modify this to match your embed map width/height in pixels
	my $tile_size = 256;

	my ($xtile, $ytile) = getTileNumber ($lat, $lon, $zoom);

	my $xtile_s = ($xtile * $tile_size - $width/2) / $tile_size;
	my $ytile_s = ($ytile * $tile_size - $height/2) / $tile_size;
	my $xtile_e = ($xtile * $tile_size + $width/2) / $tile_size;
	my $ytile_e = ($ytile * $tile_size + $height/2) / $tile_size;

	my ($lon_s, $lat_s) = getLonLat($xtile_s, $ytile_s, $zoom);
	my ($lon_e, $lat_e) = getLonLat($xtile_e, $ytile_e, $zoom);

	my $bbox = "$lon_s,$lat_s,$lon_e,$lat_e";
	return $bbox;
}