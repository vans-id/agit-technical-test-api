package dto

type EmployeeResponse struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Nip          string `json:"nip"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfBirth  string `json:"date_of_birth"`
	Age          uint   `json:"age"`
	Address      string `json:"address"`
	Religion     uint   `json:"religion"`
	Gender       uint   `json:"gender"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
}
