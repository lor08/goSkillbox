package providers

import "github.com/lor08/goSkillbox/final/internal/models"

type Provider interface {
	GetStatus(map[string]*models.Country) string
}
