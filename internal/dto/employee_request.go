package dto

type EmployeeRequest struct {
	Name         string `json:"name" binding:"required"`
	Nip          string `json:"nip" binding:"required"`
	PlaceOfBirth string `json:"place_of_birth" binding:"required"`
	DateOfBirth  string `json:"date_of_birth" binding:"required"`
	Age          uint   `json:"age"`
	Address      string `json:"address" binding:"required"`
	Religion     uint   `json:"religion" binding:"required"`
	Gender       uint   `json:"gender" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Email        string `json:"email" binding:"required"`
}
