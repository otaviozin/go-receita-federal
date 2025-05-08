package types

import (
	"time"

	"github.com/shopspring/decimal"
)

type Empresa struct {
	ID                      uint            `gorm:"primaryKey"`
	CNPJRaiz                string          `gorm:"size:8;not null"`
	NomeEmpresarial         string          `gorm:"size:255;not null"`
	CNAEPrincipal           string          `gorm:"size:10;not null"`
	NaturezaJuridica        string          `gorm:"size:10;not null"`
	CapitalSocial           decimal.Decimal `gorm:"type:decimal(20,2);not null"`
	QualificacaoResponsavel string          `gorm:"size:5"`
	CampoReservado          string          `gorm:"size:255"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}