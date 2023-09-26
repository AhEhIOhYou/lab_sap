package main

import (
	"encoding/json"
	"log"
	"os"
)

/*
3. Задать структуру данных для хранения электронных адресов включая ФИО, email, url страницы в одной из социальных сетей, возраст владельца. Реализовать методы добавления/удаления записей об электронных адресах, поиска по социальным сетям, по фамилии, загрузки и выгрузки данных в файл.
*/

const JSONFileName = "email_data.json"

type EmailKeeper struct {
	EmailData []EmailData `json:"email_data"`
}

type EmailData struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Surname   string      `json:"surname"`
	Email     string      `json:"email"`
	SocialUrl SocialLinks `json:"socialLinks"`
	Age       uint8       `json:"age"`
}

type SocialLinks struct {
	VK string `json:"vk"`
	Tg string `json:"tg"`
}

func NewEmailKeeper(data []EmailData) *EmailKeeper {
	return &EmailKeeper{
		data,
	}
}

func NewEmailData(id int, name, surname, email string, socialLinks SocialLinks, age uint8) *EmailData {
	return &EmailData{
		id,
		name,
		surname,
		email,
		socialLinks,
		age,
	}
}

func (ek *EmailKeeper) AddNewEmailData(emailData *EmailData) {
	ek.EmailData = append(ek.EmailData, *emailData)
}

func (ek *EmailKeeper) GetIndexById(id int) int {
	for i, element := range ek.EmailData {
		if element.Id == id {
			return i
		}
	}
	return -1
}

func (ek *EmailKeeper) DeleteEmailDataById(id int) {
	i := ek.GetIndexById(id)
	if i == -1 {
		log.Println("no such id")
		return
	}
	ek.EmailData[i] = ek.EmailData[len(ek.EmailData)-1]
	ek.EmailData[len(ek.EmailData)-1] = EmailData{}
	ek.EmailData = ek.EmailData[:len(ek.EmailData)-1]
}

func (ek *EmailKeeper) SaveToFile() {
	data, _ := json.MarshalIndent(ek.EmailData, "", " ")
	f, err := os.OpenFile(JSONFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	if _, err := f.Write(data); err != nil {
		log.Println(err)
	}
}

func (ek *EmailKeeper) LoadFromFile() {
	data, _ := os.ReadFile(JSONFileName)

	err := json.Unmarshal(data, &ek.EmailData)

	if err != nil {
		log.Println(err)
	}
}

func (ek *EmailKeeper) SearchBySurname(surname string) EmailData {
	for _, element := range ek.EmailData {
		if element.Surname == surname {
			return element
		}
	}
	return EmailData{}
}

func (ek *EmailKeeper) SearchBySocialLink(sl string) EmailData {
	for _, element := range ek.EmailData {
		if element.SocialUrl.Tg == sl || element.SocialUrl.VK == sl {
			return element
		}
	}
	return EmailData{}
}

func main() {
	ek := NewEmailKeeper([]EmailData{})

	ek.AddNewEmailData(NewEmailData(1, "NT1", "ST1", "mail@mail.ru", SocialLinks{Tg: "ahhaha"}, 21))

	ek.AddNewEmailData(NewEmailData(2, "NT2", "ST2", "m1l@mail.ru", SocialLinks{VK: "kek"}, 13))
	ek.AddNewEmailData(NewEmailData(3, "NT3", "ST3", "mal3@mail.ru", SocialLinks{"err", "op"}, 90))

	ek.SaveToFile()
}
