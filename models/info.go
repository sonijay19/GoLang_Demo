package models

type DoctorInfo struct {
	DrName     string
	DrmobNo    string
	DrEmail    string `bson:"_id"`
	DrPassword string
	DrLicense  string
	DrCity     string
	DrPincode  string
	Domain     string
}
