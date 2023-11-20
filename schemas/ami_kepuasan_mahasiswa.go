package model

import "github.com/golang-module/carbon/v2"

type AmiKepuasanMahasiswa struct {
	IDKepuasanMahasiswa int           `gorm:"primaryKey;column:id_kepuasan_mahasiswa" json:"-"`
	Judul               string        `gorm:"column:judul" json:"judul"`
	File                string        `gorm:"column:file" json:"file"`
	Tanggal             carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin             int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode           int           `gorm:"column:id_periode" json:"id_periode"`
}

type AmiKepuasanMahasiswaJoin struct {
	IDKepuasanMahasiswa int           `gorm:"primaryKey;column:id_kepuasan_mahasiswa" json:"-"`
	Judul               string        `gorm:"column:judul" json:"judul"`
	File                string        `gorm:"column:file" json:"file"`
	Tanggal             carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin             int           `gorm:"column:id_admin" json:"id_admin"`
	NamaAdmin           string        `gorm:"column:nm_admin" json:"nm_admin"`
	IDPeriode           int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun               string        `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiKepuasanMahasiswa) TableName() string {
	return "ami_kepuasan_mahasiswa"
}
