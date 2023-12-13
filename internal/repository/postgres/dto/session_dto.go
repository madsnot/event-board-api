package dto

type SessionDatabaseDTO struct {
	ID           int    `json:"id"`
	UserID       string `json:"userId"`
	TimeZone     string `json:"timeZone"`
	RefreshToken string `json:"refreshToken"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	DeletedAt    string `json:"deletedAt"`
}
