package types

type Estabelecimento struct {
	ID                      uint   `gorm:"primaryKey;autoIncrement"`
	CNPJBase                string `gorm:"size:8"`
	CNPJOrdem               string `gorm:"size:4"`
	CNPJDV                  string `gorm:"size:2"`
	IdentificadorMatriz     string `gorm:"size:1"`
	NomeFantasia            string `gorm:"size:255"`
	SituacaoCadastral       string `gorm:"size:2"`
	DataSituacaoCadastral   string `gorm:"size:8"`
	MotivoSituacaoCadastral string `gorm:"size:2"`
	NomeCidadeExterior      string `gorm:"size:255"`
	Pais                    string `gorm:"size:3"`
	DataInicioAtividade     string `gorm:"size:8"`
	CNAEPrincipal           string `gorm:"size:7"`
	CNAESecundario          string `gorm:"size:255"`
	TipoLogradouro          string `gorm:"size:50"`
	Logradouro              string `gorm:"size:255"`
	Numero                  string `gorm:"size:20"`
	Complemento             string `gorm:"size:255"`
	Bairro                  string `gorm:"size:100"`
	CEP                     string `gorm:"size:8"`
	UF                      string `gorm:"size:2"`
	CodigoMunicipio         string `gorm:"size:10"`
	DDD1                    string `gorm:"size:2"`
	Telefone1               string `gorm:"size:15"`
	DDD2                    string `gorm:"size:2"`
	Telefone2               string `gorm:"size:15"`
	DDDFax                  string `gorm:"size:2"`
	Fax                     string `gorm:"size:15"`
	Email                   string `gorm:"size:255"`
	SituacaoEspecial        string `gorm:"size:255"`
	DataSituacaoEspecial    string `gorm:"size:8"`
}