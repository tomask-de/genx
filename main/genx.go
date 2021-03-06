package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/tomask-de/genx"
)

func main() {
	var source = flag.String("s", "in.csv", "source file name")
	var dest = flag.String("d", "com_objects.csv", "destination file name")
	flag.Parse()

	f, err := os.Open(*source)
	if err != nil {
		fmt.Printf("could not open source file: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()
	d, err := os.OpenFile(*dest, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("could not create destination file: %s", err.Error())
		os.Exit(2)
	}
	defer d.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.FieldsPerRecord = 8

	w := csv.NewWriter(d)
	w.Comma = ','
	defer w.Flush()

	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("could not read sample file: %s", err.Error())
		return
	}
	var floors map[string]*genx.Floor = make(map[string]*genx.Floor, 0)
	var currentFloor string
	currentFloorNumber := 0
	for i, r := range records {
		if i == 0 {
			// skip header
			continue
		}
		if currentFloor == "" || currentFloor != r[0] {
			currentFloor = r[0]
			currentFloorNumber++
		}
		lights, err := strconv.Atoi(r[3])
		shutters, err := strconv.Atoi(r[4])
		heating, err := strconv.Atoi(r[5])
		sockets, err := strconv.Atoi(r[6])
		reedContacts, err := strconv.Atoi(r[7])

		if err != nil {
			fmt.Printf("could not read input: %s", err.Error())
			os.Exit(3)
		}

		room := genx.NewRoom(r[1], r[2], lights, shutters, heating, sockets, reedContacts)
		f, ok := floors[currentFloor]
		if ok {
			f.Rooms = append(f.Rooms, room)
		} else {
			newFloor := genx.NewFloor(currentFloor, currentFloorNumber, room)
			floors[currentFloor] = newFloor
		}
	}
	w.Write([]string{"Main", "Middle", "Sub", "Main", "Middle", "Sub", "Central", "Unfiltered", "Description", "DatapointType", "Security"})
	w.Write([]string{"", "", "", "0", "", "", "", "", "", "", ""})
	w.Write([]string{"", "", "", "0", "0", "", "", "", "", "", ""})
	w.Write([]string{"", "", "", "0", "0", "", "", "", "", "", ""})
	w.Write([]string{"", "", "", "0", "0", "1", "", "", "", "", ""})
	for _, floor := range floors {
		floor.Generate(w, genx.Actuators[genx.MDT])
	}

}
