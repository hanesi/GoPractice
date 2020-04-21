package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-geom"
)

// Rat's nest of a KML generated from this website:
// https://www.onlinetool.io/xmltogo/
// Then deconstructed into non-inline types
type Kml struct {
	XMLName  xml.Name `xml:"kml"`
	Document Document `xml:"Document"`
}

type Document struct {
	Placemark Placemark `xml:"Placemark"`
}

type Placemark struct {
	Polygon Polygon `xml:"Polygon"`
}

type Polygon struct {
	Text            string          `xml:",chardata"`
	Extrude         string          `xml:"extrude"`
	Tessellate      string          `xml:"tessellate"`
	AltitudeMode    string          `xml:"altitudeMode"`
	OuterBoundaryIs OuterBoundaryIs `xml:"outerBoundaryIs"`
}

type OuterBoundaryIs struct {
	Text       string     `xml:",chardata"`
	LinearRing LinearRing `xml:"LinearRing"`
}

type LinearRing struct {
	Text        string `xml:",chardata"`
	Coordinates string `xml:"coordinates"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("/Users/ianhanes/Desktop/General/KMLs/all-zips/zip77345.kml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var data Kml
	xml.Unmarshal(byteValue, &data)
	polyLine := data.Document.Placemark.Polygon.OuterBoundaryIs.LinearRing.Coordinates
	polyList := strings.Split(polyLine, ",0.0 -")
	coordList := [][]geom.Coord{}
	for _, v := range polyList[:len(polyList)-2] {
		v = strings.Replace(v, "-", "", -1)
		//v = strings.Replace(v, " 0.0", "", -1)
		coordList = append(coordList, coordStringToFloat(v))
	}

	var q geom.Layout = 1
	shape := geom.NewPolygon(q).MustSetCoords(coordList)
	f := *shape
	fmt.Println(f.Area())
}

func coordStringToFloat(coord string) []geom.Coord {
	retArray := []float64{}
	whatever := []geom.Coord{}
	coords := strings.Split(coord, ",")
	for _, v := range coords {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			break
		}
		retArray = append(retArray, num)
	}
	x := geom.Coord(retArray)
	whatever = append(whatever, x)
	return whatever
}
