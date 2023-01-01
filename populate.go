package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

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

	PERSONE_MIN    = 500
	PERSONE_MAX    = 1000
	GRUPPI_MIN     = 500
	GRUPPI_MAX     = 1000
	ARTISTI_MIN    = 500
	ARTISTI_MAX    = 1000
	SPETTACOLI_MIN = 5000
	SPETTACOLI_MAX = 10000
	EVENTI_MIN     = 100
	EVENTI_MAX     = 500
	LUOGHI_MIN     = 500
	LUOGHI_MAX     = 1000
	FORNITORI_MIN  = 100
	FORNITORI_MAX  = 100
	SETTORI_MIN    = LUOGHI_MIN * 4
	SETTORI_MAX    = LUOGHI_MAX * 6
	POSTI_MIN      = SETTORI_MIN * 30
	POSTI_MAX      = SETTORI_MAX * 50
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

type Persona struct {
	ID          int
	Nome        string `faker:"first_name"`
	Cognome     string `faker:"last_name"`
	DataNascita string `faker:"date"`
}

func PersonaFields() []string {
	return []string{"id", "nome", "cognome", "data_nascita"}
}

func (a Persona) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(a.ID, a.Nome, a.Cognome, a.DataNascita)
	return
}

func (a Persona) GetID() int {
	return a.ID
}

type Gruppo struct {
	ID             int
	DataFormazione string `faker:"date"`
}

func GruppoFields() []string {
	return []string{"id", "data_formazione"}
}

func (a Gruppo) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(a.ID, a.DataFormazione)
	return
}

func (a Gruppo) GetID() int {
	return a.ID
}

type PersonaGruppoAppartenenza struct {
	Persona int
	Gruppo  int
}

type Artista struct {
	ID       int
	NomeArte string `faker:"username"`
	Persona  *int
	Gruppo   *int
}

func ArtistaFields() []string {
	return []string{"id", "nome_arte", "persona", "gruppo"}
}

func (a Artista) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(a.ID, a.NomeArte, a.Persona, a.Gruppo)
	return
}

func (a Artista) GetID() int {
	return a.ID
}

type Spettacolo struct {
	ID         int
	Titolo     string `faker:"username"`
	Artista    int
	PrezzoSIAE float32 `faker:"amount"`
	Cachet     float32 `faker:"amount"`
}

func SpettacoloFields() []string {
	return []string{"id", "titolo", "artista", "prezzo_siae", "cachet"}
}

func (s Spettacolo) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Titolo, s.Artista, s.PrezzoSIAE, s.Cachet)
	return
}

func (s Spettacolo) GetID() int {
	return s.ID
}

type Luogo struct {
	ID        int
	Tipo      string  `faker:"oneof: arena, palazzetto, parco, piazza, stadio, teatro"`
	Nome      string  `faker:"username,unique"`
	Indirizzo string  `faker:"oneof: roma 23,napoli 11"`
	Citta     string  `faker:"oneof: Bologna,Milano,Roma"`
	Prezzo    float32 `faker:"amount"`
}

func LuogoFields() []string {
	return []string{"id", "tipo", "nome", "indirizzo", "citta", "prezzo"}
}

func (v Luogo) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(v.ID, v.Tipo, v.Nome, v.Indirizzo, v.Citta, v.Prezzo)
	return
}

func (v Luogo) GetID() int {
	return v.ID
}

type Evento struct {
	ID         int
	Spettacolo int
	Luogo      int
	Titolo     string `faker:"first_name"`
	Inizio     string
	Fine       string
}

func EventoFields() []string {
	return []string{"id", "spettacolo", "luogo", "titolo", "inizio", "fine"}
}

func (e Evento) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(e.ID, e.Spettacolo, e.Luogo, e.Titolo, e.Inizio, e.Fine)
	return
}

func (e Evento) GetID() int {
	return e.ID
}

type Settore struct {
	ID       int
	Luogo    int
	Nome     string `faker:"username,unique"`
	Capienza int    `faker:"oneof: 30,40,60"`
}

func SettoreFields() []string {
	return []string{"id", "luogo", "nome", "capienza"}
}

func (s Settore) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Luogo, s.Nome, s.Capienza)
	return
}

func (s Settore) GetID() int {
	return s.ID
}

type Fornitore struct {
	ID          int
	Nome        string `faker:"name,unique"`
	Descrizione string `faker:"sentence"`
}

func FornitoreFields() []string {
	return []string{"id", "nome", "descrizione"}
}

func (s Fornitore) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Nome, s.Descrizione)
	return
}

func (s Fornitore) GetID() int {
	return s.ID
}

type EventoFornitoreServizio struct {
	Fornitore int
	Tipo      string `faker:"oneof: audio,biglietteria,luci,maschere,sicurezza,regia"`
	Evento    int
	Prezzo    float64 `faker:"amount"`
}

func EventoFornitoreServizioFields() []string {
	return []string{"fornitore", "tipo", "evento", "prezzo"}
}

func (s EventoFornitoreServizio) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.Fornitore, s.Tipo, s.Evento, s.Prezzo)
	return
}

type Posto struct {
	ID      int
	Settore int
	Fila    string
	Numero  int
}

func PostoFields() []string {
	return []string{"id", "settore", "fila", "numero"}
}

func (s Posto) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.ID, s.Settore, s.Fila, s.Numero)
	return
}

func (s Posto) GetID() int {
	return s.ID
}

type SettoreEventoCosto struct {
	Settore int
	Evento  int
	Prezzo  float64 `faker:"amount"`
}

func SettoreEventoCostoFields() []string {
	return []string{"settore", "evento", "prezzo"}
}

func (s SettoreEventoCosto) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.Settore, s.Evento, s.Prezzo)
	return
}

type Biglietto struct {
	Codice     int
	Nominativo string `faker:"name"`
	Posto      int
	Evento     int
}

func BigliettoFields() []string {
	return []string{"codice", "nominativo", "posto", "evento"}
}

func (s Biglietto) Populate(stmt *sql.Stmt) (err error) {
	_, err = stmt.Exec(s.Codice, s.Nominativo, s.Posto, s.Evento)
	return
}

func (s Biglietto) GetID() int {
	return s.Codice
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
	chars := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	return chars[rand.Intn(len(chars))], nil
}

func randomTimestamp() (time.Time, time.Time) {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	randomAfterNow := randomNow.Add(time.Hour * time.Duration(rand.Intn(25)))
	return randomNow, randomAfterNow
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
	views, err := ioutil.ReadFile("views.sql")
	if err != nil {
		log.Fatalf("Could not read views.sql: %v", err)
	}

	if _, err = db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;"); err != nil {
		log.Fatalf("Could not purge all data: %v", err)
	}
	if _, err = db.Exec(string(schema)); err != nil {
		log.Fatalf("Could not apply schema.sql: %v", err)
	}
	log.Println("Created schema")
	if _, err = db.Exec(string(views)); err != nil {
		log.Fatalf("Could not apply views.sql: %v", err)
	}
	log.Println("Applied views")

	persone := []Persona{}
	if err = faker.FakeData(
		&persone,
		options.WithRandomMapAndSliceMinSize(PERSONE_MIN),
		options.WithRandomMapAndSliceMaxSize(PERSONE_MAX),
	); err != nil {
		log.Fatalf("Could not fill artists data: %v", err)
	}
	for i := range persone {
		persone[i].ID = i + 1
	}
	bulkInsert("persona", PersonaFields(), persone)
	log.Printf("Popolato %d persone", len(persone))

	gruppi := []Gruppo{}
	if err = faker.FakeData(
		&gruppi,
		options.WithRandomMapAndSliceMinSize(GRUPPI_MIN),
		options.WithRandomMapAndSliceMaxSize(GRUPPI_MAX),
	); err != nil {
		log.Fatalf("Could not fill artists data: %v", err)
	}
	for i := range gruppi {
		gruppi[i].ID = i + 1
	}
	bulkInsert("gruppo", GruppoFields(), gruppi)
	log.Printf("Popolato %d gruppi", len(gruppi))

	artisti := []Artista{}
	if err = faker.FakeData(
		&artisti,
		options.WithRandomMapAndSliceMinSize(ARTISTI_MIN),
		options.WithRandomMapAndSliceMaxSize(ARTISTI_MAX),
	); err != nil {
		log.Fatalf("Could not fill artists data: %v", err)
	}
	for i := range artisti {
		artisti[i].ID = i + 1
		if i%2 == 0 {
			id := persone[rand.Intn(len(persone))].GetID()
			artisti[i].Persona = &id
			artisti[i].Gruppo = nil
		} else {
			id := gruppi[rand.Intn(len(gruppi))].GetID()
			artisti[i].Gruppo = &id
			artisti[i].Persona = nil
		}
	}
	bulkInsert("artista", ArtistaFields(), artisti)
	log.Printf("Popolato %d artisti", len(artisti))

	spettacoli := []Spettacolo{}
	if err = faker.FakeData(
		&spettacoli,
		options.WithCustomFieldProvider("Titolo", title),
		options.WithCustomFieldProvider("Artista", onlyIDs(artisti)),
		options.WithRandomMapAndSliceMinSize(SPETTACOLI_MIN),
		options.WithRandomMapAndSliceMaxSize(SPETTACOLI_MAX),
	); err != nil {
		log.Fatalf("Could not fill shows data: %v", err)
	}
	for i := range spettacoli {
		spettacoli[i].ID = i + 1
	}
	bulkInsert("spettacolo", SpettacoloFields(), spettacoli)
	log.Printf("Popolato %d spettacoli", len(spettacoli))

	fornitori := []Fornitore{}
	if err = faker.FakeData(
		&fornitori,
		options.WithRandomMapAndSliceMinSize(FORNITORI_MIN),
		options.WithRandomMapAndSliceMaxSize(FORNITORI_MAX),
	); err != nil {
		log.Fatalf("Could not fill shows data: %v", err)
	}
	for i := range fornitori {
		fornitori[i].ID = i + 1
	}
	bulkInsert("fornitore", FornitoreFields(), fornitori)
	log.Printf("Popolato %d spettacoli", len(fornitori))

	luoghi := []Luogo{}
	if err = faker.FakeData(
		&luoghi,
		options.WithRandomMapAndSliceMinSize(LUOGHI_MIN),
		options.WithRandomMapAndSliceMaxSize(LUOGHI_MAX),
	); err != nil {
		log.Fatalf("Could not fill venues data: %v", err)
	}
	for i := range luoghi {
		luoghi[i].ID = i + 1
	}
	bulkInsert("luogo", LuogoFields(), luoghi)
	log.Printf("Popolato %d luoghi", len(luoghi))

	eventi := []Evento{}
	if err = faker.FakeData(
		&eventi,
		options.WithCustomFieldProvider("Spettacolo", onlyIDs(spettacoli)),
		options.WithCustomFieldProvider("Luogo", onlyIDs(luoghi)),
		options.WithRandomMapAndSliceMinSize(EVENTI_MIN),
		options.WithRandomMapAndSliceMaxSize(EVENTI_MAX),
	); err != nil {
		log.Fatalf("Could not fill events data: %v", err)
	}
	for i := range eventi {
		eventi[i].ID = i + 1
		n, a := randomTimestamp()
		eventi[i].Inizio = n.String()[:19]
		eventi[i].Fine = a.String()[:19]
	}
	bulkInsert("evento", EventoFields(), eventi)
	log.Printf("Popolato %d eventi", len(eventi))

	settori := []Settore{}
	if err = faker.FakeData(
		&settori,
		options.WithCustomFieldProvider("Luogo", onlyIDs(luoghi)),
		options.WithRandomMapAndSliceMinSize(SETTORI_MIN),
		options.WithRandomMapAndSliceMaxSize(SETTORI_MAX),
	); err != nil {
		log.Fatalf("Could not fill venue sectors data: %v", err)
	}
	for i := range settori {
		settori[i].ID = i + 1
	}
	bulkInsert("settore", SettoreFields(), settori)
	log.Printf("Popolato %d settori", len(settori))

	posti := []Posto{}
	i := 0
	chars := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for _, s := range settori {
		nCols := s.Capienza / len(chars)
		for _, r := range chars {
			for c := 0; c < nCols; c++ {
				posti = append(posti, Posto{
					ID:      i,
					Settore: s.ID, Fila: r, Numero: c})
				i++
			}
		}
	}
	bulkInsert("posto", PostoFields(), posti)
	log.Printf("Popolato %d posti", len(posti))

	settoreEventoCosti := []SettoreEventoCosto{}
	for _, event := range eventi {
		for _, sector := range settori {
			if sector.Luogo == event.Luogo {
				var ele SettoreEventoCosto
				if err = faker.FakeData(&ele); err != nil {
					log.Fatalf("Could not generate one venue sector events price: %v", err)
				}
				ele.Evento = event.GetID()
				ele.Settore = sector.GetID()
				ele.Prezzo += 1
				settoreEventoCosti = append(settoreEventoCosti, ele)
			}
		}
	}
	bulkInsert("settore_evento_costo", SettoreEventoCostoFields(), settoreEventoCosti)
	log.Printf("Popolato %d costi settore evento", len(settoreEventoCosti))

	eventoFornitoreServizi := []EventoFornitoreServizio{}
	if err = faker.FakeData(
		&eventoFornitoreServizi,
		options.WithCustomFieldProvider("Fornitore", onlyIDs(fornitori)),
		options.WithCustomFieldProvider("Evento", onlyIDs(eventi)),
	); err != nil {
		log.Fatalf("Could not fill shows data: %v", err)
	}
	bulkInsert("evento_fornitore_servizio", EventoFornitoreServizioFields(), eventoFornitoreServizi)
	log.Printf("Popolato %d evento fornitore servizio", len(eventoFornitoreServizi))

	biglietti := []Biglietto{}
	i = 0
	for _, settoreEventoCosto := range settoreEventoCosti {
		for _, posto := range posti {
			if posto.Settore == settoreEventoCosto.Settore {
				if rand.Intn(10) > 7 {
					var ticket Biglietto
					ticket.Codice = i
					ticket.Evento = settoreEventoCosto.Evento
					ticket.Posto = posto.GetID()
					ticket.Nominativo = faker.Name()
					i++
					biglietti = append(biglietti, ticket)
				}
			}
		}
	}
	bulkInsert("biglietto", BigliettoFields(), biglietti)
	log.Printf("Popolato %d biglietti", len(biglietti))
}
