package models

import (
	"time"
)

type User struct {
	Base
	Name      string `gorm:"size:100;not null" json:"name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null" json:"password"`
	Telephone string `gorm:"size:20;not null" json:"telephone"`

	// Relacionamentos acadêmicos
	UniversityID *int        `json:"university_id,omitempty"`
	University   *University `gorm:"foreignKey:UniversityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"university,omitempty"`
	CourseID     *int        `json:"course_id,omitempty"`
	Course       *Course     `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"course,omitempty"`

	// Informações acadêmicas
	StudentStatus   string `gorm:"size:20;default:'viewer'" json:"student_status"` // 'active', 'graduated', 'viewer', 'prospective'
	CurrentSemester *int   `json:"current_semester,omitempty"`                     // Semestre atual (1-10)
	GraduationYear  *int   `json:"graduation_year,omitempty"`                      // Ano de formatura
	EntryYear       *int   `json:"entry_year,omitempty"`                           // Ano de ingresso

	// Campos pessoais
	Bio         *string    `gorm:"size:500" json:"bio,omitempty"` // Descrição pessoal
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`

	// Campos para atlética
	AthleticRole *string    `gorm:"size:50" json:"athletic_role,omitempty"` // 'member', 'director', 'president'
	JoinDate     *time.Time `json:"join_date,omitempty"`                    // Data que entrou na atlética
}
