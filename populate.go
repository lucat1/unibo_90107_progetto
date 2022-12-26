package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/interfaces"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/lib/pq"
)

var db *sql.DB

const (
	HOST   = "db"
	DB     = "tommaso"
	USER   = "postgres"
	PASSWD = "postgres"

	ARTISTS_MIN            = 500
	ARTISTS_MAX            = 1000
	SHOWS_MIN              = 5000
	SHOWS_MAX              = 10000
	VENUE_TYPE_MAX         = 5
	EVENTS_MIN             = 100
	EVENTS_MAX             = 500
	VENUES_MIN             = 500
	VENUES_MAX             = 1000
	VENUE_SECTORS_MIN      = VENUES_MIN * 4
	VENUE_SECTORS_MAX      = VENUES_MAX * 6
	VENUE_SECTOR_SEATS_MIN = VENUE_SECTORS_MIN * 30
	VENUE_SECTOR_SEATS_MAX = VENUE_SECTORS_MAX * 50
)

type Populate interface {
	Populate(stmt *sql.Stmt) error
}

type WithID interface {
	GetID() int
}

func bulkInsert[T Populate](table string, fields []string, vec []T) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Could not begin bulk insert for %s: %v", table, err)
	}

	stmt, err := tx.Prepare(pq.CopyIn(table, fields...))
	if err != nil {
		log.Fatalf("Could not prepare bulk insert for %s: %v", table, err)
	}

	for _, t := range vec {
		t.Populate(stmt)
		if err != nil {
			log.Fatalf("Could not populate statement for %s: %v", table, err)
		}
	}

	if _, err = stmt.Exec(); err != nil {
		log.Fatalf("Could not execute bulk insert for %s: %v", table, err)
	}
	if tx.Commit() != nil {
		log.Fatalf("Could not commit bulk insert for %s: %v", table, err)
	}
}

type Artist struct {
	ID   int
	Name string `faker:"first_name"`
}

func ArtistFields() []string {
	return []string{"id", "name"}
}

func (a Artist) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(a.ID, a.Name)
	return
}

func (a Artist) GetID() int {
	return a.ID
}

type Show struct {
	ID        int
	Title     string
	Artist    int
	SIAEPrice float32
	Cachet    float32
}

func ShowFields() []string {
	return []string{"id", "title", "artist", "siae_price", "cachet"}
}

func (s Show) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Title, s.Artist, s.SIAEPrice, s.Cachet)
	return
}

func (s Show) GetID() int {
	return s.ID
}

type VenueType struct {
	ID   int
	Name string
}

func VenueTypeFields() []string {
	return []string{"id", "name"}
}

func (s VenueType) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Name)
	return
}

func (s VenueType) GetID() int {
	return s.ID
}

type Venue struct {
	ID     int
	Type   int    `faker:"oneof: 1,2,3,4,5"`
	Name   string `faker:"first_name"`
	Coords string
	Price  float32
}

func VenueFields() []string {
	return []string{"id", "type", "name", "coords", "price"}
}

func (v Venue) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(v.ID, v.Type, v.Name, v.Coords, v.Price)
	return
}

func (v Venue) GetID() int {
	return v.ID
}

type Event struct {
	ID       int
	Show     int
	Venue    int
	Title    string `faker:"first_name"`
	StartsAt string `faker:"timestamp"`
	EndAt    string `faker:"timestamp"`
}

func EventFields() []string {
	return []string{"id", "show", "venue", "title", "starts_at", "ends_at"}
}

func (e Event) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(e.ID, e.Show, e.Venue, e.Title, e.StartsAt, e.EndAt)
	return
}

func (e Event) GetID() int {
	return e.ID
}

type VenueSector struct {
	ID       int
	Venue    int
	Name     string `faker:"century"`
	Capacity string `faker:"year"`
}

func VenueSectorFields() []string {
	return []string{"id", "venue", "name", "capacity"}
}

func (s VenueSector) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Venue, s.Name, s.Capacity)
	return
}

func (s VenueSector) GetID() int {
	return s.ID
}

type VenueSectorSeat struct {
	ID     int
	Sector int
	Row    string
	Col    int `faker:"oneof: 1,2,3,4,5,6,7,8,9,10"`
}

func VenueSectorSeatFields() []string {
	return []string{"id", "sector", "row", "col"}
}

func (s VenueSectorSeat) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Sector, s.Row, s.Col)
	return
}

func (s VenueSectorSeat) GetID() int {
	return s.ID
}

type VenueSectorEventsPrice struct {
	Sector int
	Event  int
	Price  float64 `faker:"amount"`
}

func VenueSectorEventsPriceFields() []string {
	return []string{"sector", "event", "price"}
}

func (s VenueSectorEventsPrice) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.Sector, s.Event, s.Price)
	return
}

type Ticket struct {
	ID    int
	Seat  int
	Event int
}

func TicketFields() []string {
	return []string{"id", "seat", "event"}
}

func (s Ticket) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Seat, s.Event)
	return
}

func (s Ticket) GetID() int {
	return s.ID
}

func onlyIDs[T WithID](vec []T) interfaces.CustomProviderFunction {
	l := len(vec)
	return func() (interface{}, error) {
		return vec[rand.Intn(l)].GetID(), nil
	}
}

func price(min, max float32) interfaces.CustomProviderFunction {
	return func() (interface{}, error) {
		return min + rand.Float32()*(max-min), nil
	}
}

func title() (interface{}, error) {
	return faker.Name() + " " + faker.Name(), nil
}

func point() (interface{}, error) {
	return "(" + fmt.Sprintf("%f", faker.Latitude()) + "," + fmt.Sprintf("%f", faker.Longitude()) + ")", nil
}

func row() (interface{}, error) {
	chars := []string{"A", "B", "C", "D", "E", "F"}
	return chars[rand.Intn(len(chars))], nil
}

func main() {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWD, DB))
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	schema, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatalf("Could not read schema.sql: %v", err)
	}

	if _, err = db.Exec("DROP SCHEMA public CASCADE;" + "CREATE SCHEMA public;"); err != nil {
		log.Fatalf("Could not purge all data: %v", err)
	}
	if _, err = db.Exec(string(schema)); err != nil {
		log.Fatalf("Could not apply schema.sql: %v", err)
	}
	log.Println("Created schema")

	artists := []Artist{}
	if err = faker.FakeData(
		&artists,
		options.WithRandomMapAndSliceMinSize(ARTISTS_MIN),
		options.WithRandomMapAndSliceMaxSize(ARTISTS_MAX),
	); err != nil {
		log.Fatalf("Could not fill artists data: %v", err)
	}
	for i := range artists {
		artists[i].ID = i + 1
	}
	bulkInsert("artists", ArtistFields(), artists)
	log.Printf("Populated %d artists", len(artists))

	shows := []Show{}
	if err = faker.FakeData(
		&shows,
		options.WithCustomFieldProvider("Title", title),
		options.WithCustomFieldProvider("Cachet", price(100, 100000)),
		options.WithCustomFieldProvider("SIAEPrice", price(500, 100000)),
		options.WithCustomFieldProvider("Artist", onlyIDs(artists)),
		options.WithRandomMapAndSliceMinSize(SHOWS_MIN),
		options.WithRandomMapAndSliceMaxSize(SHOWS_MAX),
	); err != nil {
		log.Fatalf("Could not fill shows data: %v", err)
	}
	for i := range shows {
		shows[i].ID = i + 1
	}
	bulkInsert("shows", ShowFields(), shows)
	log.Printf("Populated %d shows", len(shows))

	venueTypes := []VenueType{
		{1, "stadium"},
		{2, "opera hall"},
		{3, "fair"},
		{4, "pub"},
		{5, "auditorium"},
	}
	bulkInsert("venue_types", VenueTypeFields(), venueTypes)
	log.Printf("Populated %d venue types", len(venueTypes))

	venues := []Venue{}
	if err = faker.FakeData(
		&venues,
		options.WithCustomFieldProvider("Type", onlyIDs(venueTypes)),
		options.WithCustomFieldProvider("Coords", point),
		options.WithCustomFieldProvider("Price", price(1000, 5000)),
		options.WithRandomMapAndSliceMinSize(VENUES_MIN),
		options.WithRandomMapAndSliceMaxSize(VENUES_MAX),
	); err != nil {
		log.Fatalf("Could not fill venues data: %v", err)
	}
	for i := range venues {
		venues[i].ID = i + 1
	}
	bulkInsert("venues", VenueFields(), venues)
	log.Printf("Populated %d venues", len(venues))

	events := []Event{}
	if err = faker.FakeData(
		&events,
		options.WithCustomFieldProvider("Show", onlyIDs(shows)),
		options.WithCustomFieldProvider("Venue", onlyIDs(venues)),
		options.WithRandomMapAndSliceMinSize(EVENTS_MIN),
		options.WithRandomMapAndSliceMaxSize(EVENTS_MAX),
	); err != nil {
		log.Fatalf("Could not fill events data: %v", err)
	}
	for i := range events {
		events[i].ID = i + 1
	}
	bulkInsert("events", EventFields(), events)
	log.Printf("Populated %d events", len(events))

	venueSectors := []VenueSector{}
	if err = faker.FakeData(
		&venueSectors,
		options.WithCustomFieldProvider("Venue", onlyIDs(venues)),
		options.WithRandomMapAndSliceMinSize(VENUE_SECTORS_MIN),
		options.WithRandomMapAndSliceMaxSize(VENUE_SECTORS_MAX),
	); err != nil {
		log.Fatalf("Could not fill venue sectors data: %v", err)
	}
	for i := range venueSectors {
		venueSectors[i].ID = i + 1
	}
	bulkInsert("sectors", VenueSectorFields(), venueSectors)
	log.Printf("Populated %d venue sectors", len(venueSectors))

	venueSectorsSeats := []VenueSectorSeat{}
	if err = faker.FakeData(
		&venueSectorsSeats,
		options.WithCustomFieldProvider("Sector", onlyIDs(venueSectors)),
		options.WithCustomFieldProvider("Row", row),
		options.WithRandomMapAndSliceMinSize(VENUE_SECTOR_SEATS_MIN),
		options.WithRandomMapAndSliceMaxSize(VENUE_SECTOR_SEATS_MAX),
	); err != nil {
		log.Fatalf("Could not fill venue sectors data: %v", err)
	}
	for i := range venueSectorsSeats {
		venueSectorsSeats[i].ID = i + 1
	}
	bulkInsert("seats", VenueSectorSeatFields(), venueSectorsSeats)
	log.Printf("Populated %d venue sectors seats", len(venueSectorsSeats))

	venueSectorsEventsPrices := []VenueSectorEventsPrice{}
	for _, event := range events {
		for _, sector := range venueSectors {
			if sector.Venue == event.Venue {
				var ele VenueSectorEventsPrice
				if err = faker.FakeData(&ele); err != nil {
					log.Fatalf("Could not generate one venue sector events price: %v", err)
				}
				ele.Event = event.GetID()
				ele.Sector = sector.GetID()
				ele.Price += 1
				venueSectorsEventsPrices = append(venueSectorsEventsPrices, ele)
			}
		}
	}
	bulkInsert("sectors_events_prices", VenueSectorEventsPriceFields(), venueSectorsEventsPrices)
	log.Printf("Populated %d venue sectors events prices", len(venueSectorsEventsPrices))

	tickets := []Ticket{}
	i := 0
	for _, event := range events {
		for _, seat := range venueSectors {
			if rand.Intn(10) > 7 {
				var ticket Ticket
				ticket.ID = i
				ticket.Event = event.GetID()
				ticket.Seat = seat.GetID()
				i++
				tickets = append(tickets, ticket)
			}
		}
	}
	bulkInsert("tickets", TicketFields(), tickets)
	log.Printf("Populated %d tickets", len(tickets))
}
