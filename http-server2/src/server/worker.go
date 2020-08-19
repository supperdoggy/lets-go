package main

type worker struct {
	name        string `json:"name"`
	age         string `json:"age"`
	position    string `json:"position"`
	job         string `json:"job"`
	phoneNumber string `json:"phone_number"`
	email       string `json:"email"`
	id          string `json:"id"`
}

func (w *worker) getInfo() (result string) {
	result = "Name: " + w.name + "\nAge: " + w.age +
		"\nPhone: " + w.phoneNumber + "\nEmail: " + w.email +
		"\nJob: " + w.job + "\nPosition " + w.position + "\n"
	return
}
