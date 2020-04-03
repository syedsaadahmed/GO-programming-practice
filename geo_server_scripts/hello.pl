#!/usr/bin/perl

use Math::Trig;

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

# convert from permalink OSM format like:
# https://www.openstreetmap.org/?lat=43.731049999999996&lon=15.79375&zoom=13&layers=M
# to OSM "Export" iframe embedded bbox format like:
# https://www.openstreetmap.org/export/embed.html?bbox=15.7444,43.708,15.8431,43.7541&layer=mapnik

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



my $answer = LonLat_to_bbox(60.231050,25.332302,8);
print $answer


