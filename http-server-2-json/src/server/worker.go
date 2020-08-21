package main

import "math/rand"

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

// returns random string of length
func randomStringGenerator(l int) string {
	// letters and number to random pick
	const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, l)
	for i := range result {
		result[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	return string(result)
}
