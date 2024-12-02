package main

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	//Date  string `json:"date"`
}

type ContactsStore interface {
	Add(data Contact) error
	GetAll() ([]Contact, error)
}

type store struct {
	conn *pgx.Conn
}

var connected = false

func NewStore() *store {
	//tunnel := NewSSHTunnel(
	//	"sviatkyo@sviatkyou.cv.ua",
	//	ssh.Password("Thegovernmentsucks1488!"), // 2. password
	//	// The destination host and port of the actual server.
	//	"localhost:5433",
	//)
	tunel()

	conn, err := pgx.Connect(pgx.ConnConfig{Host: "localhost:5433", User: "sviatkyo", Password: "Thegovernmentsucks1488!", Database: "sviatkyo_contacts"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer conn.Close()
	//var connect Contact
	//err = conn.QueryRow("select * from contacts where id=1").Scan(&connect.Name)
	//fmt.Println(connect, err)

	return &store{conn: conn}
}

func (s store) Add(data Contact) error {
	var err error

	contacts, err := s.GetAll()
	if err != nil {
		return err
	}

	if len(contacts) > 0 {
		data.Id = contacts[len(contacts)-1].Id + 1
	} else {
		data.Id = 0
	}

	//data.Date = time.Now().Format("2006-01-02 15:04:05")

	//jsonData.Contacts = append(jsonData.Contacts, data)
	//bytes, err := json.Marshal(jsonData)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(jsonData)

	return s.conn.QueryRow("insert into contacts values ($1, $2, $3, $4)", data.Id, data.Name, data.Email, data.Phone).Scan(&data.Id)
}

func (s store) GetAll() ([]Contact, error) {
	var contacts []Contact
	r, err := s.conn.Query("select * from contacts")
	if err != nil {
		fmt.Println("Error getting contacts", err)
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		var contact Contact
		err := r.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
		}
		contacts = append(contacts, contact)
	}
	fmt.Println(contacts)
	return contacts, nil
}

//func (s store) GetAll() ([]Contact, error) {
//	var jsonData struct{ Contacts []Contact }
//
//	readData, err := os.ReadFile("s.filename")
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(readData, &jsonData)
//	if err != nil {
//		return nil, err
//	}
//
//	if jsonData.Contacts == nil {
//		jsonData.Contacts = []Contact{}
//	}
//
//	return jsonData.Contacts, nil
//}
