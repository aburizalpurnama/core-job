package model

import (
	database "core-users-service/databases/rdbms"
	"time"
)

const TableNameCif = "users.cifs"

// Cif mapped from table <cifs>
type Cif struct {
	database.BaseModel
	Cif                          int64       `gorm:"column:cif" json:"cif"`
	BrinetsCif                   *string     `gorm:"column:brinets_cif" json:"brinets_cif"`
	IDNumber                     string      `gorm:"column:id_number;not null" json:"id_number"`
	Gender                       string      `gorm:"column:gender;not null;default:male" json:"gender"`
	FullName                     string      `gorm:"column:full_name;not null" json:"full_name"`
	ShortName                    string      `gorm:"column:short_name" json:"short_name"`
	BirthDate                    time.Time   `gorm:"column:birth_date;not null" json:"birth_date"`
	Occupation                   string      `gorm:"column:occupation;not null" json:"occupation"`
	MotherName                   string      `gorm:"column:mother_name;not null" json:"mother_name"`
	PhoneNumber                  string      `gorm:"column:phone_number;not null" json:"phone_number"`
	Email                        string      `gorm:"column:email;not null" json:"email"`
	Address                      string      `gorm:"column:address;not null" json:"address"`
	Rt                           string      `gorm:"column:rt" json:"rt"`
	Rw                           string      `gorm:"column:rw" json:"rw"`
	StateCode                    string      `gorm:"column:state_code;not null" json:"state_code"`
	CityCode                     string      `gorm:"column:city_code;not null" json:"city_code"`
	District                     string      `gorm:"column:district;not null;comment:kecamatan" json:"district"`       // kecamatan
	Subdistrict                  string      `gorm:"column:subdistrict;not null;comment:kelurahan" json:"subdistrict"` // kelurahan
	ZipCode                      string      `gorm:"column:zip_code;not null" json:"zip_code"`
	MonthlyExpense               float64     `gorm:"column:monthly_expense" json:"monthly_expense"`
	MonthlyNettIncome            float64     `gorm:"column:monthly_nett_income" json:"monthly_nett_income"`
	MaritalStatus                string      `gorm:"column:marital_status;not null;default:single" json:"marital_status"`
	EmploymentType               string      `gorm:"column:employment_type;not null" json:"employment_type"`
	Education                    string      `gorm:"column:education;not null" json:"education"`
	EmergencyContactName         string      `gorm:"column:emergency_contact_name" json:"emergency_contact_name"`
	EmergencyContactRelationship string      `gorm:"column:emergency_contact_relationship" json:"emergency_contact_relationship"`
	ReferenceID                  string      `gorm:"column:reference_id;not null" json:"reference_id"`
	Collectibility               int32       `gorm:"column:collectibility;not null;default:1" json:"collectibility"`
	BlockCode                    *string     `gorm:"column:block_code" json:"block_code"`
	CreatedBy                    interface{} `gorm:"column:created_by;type:json;not null" json:"created_by"`
}

type CifReport struct {
	Cif         string    `json:"cif"`
	IDNumber    string    `json:"id_number"`
	Gender      string    `json:"gender"`
	FullName    string    `json:"full_name"`
	ShortName   string    `json:"short_name"`
	BirthDate   time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	OpenCifDate string    `json:"open_cif_date"`
}

// TableName Cif's table name
func (*Cif) TableName() string {
	return TableNameCif
}
