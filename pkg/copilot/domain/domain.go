package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"createdAt,omitempty"`
}
type Attendants struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	REF        string             `bson:"ref" json:"ref"`
	CPF        string             `bson:"cpf" json:"cpf"`
	Customer   primitive.ObjectID `bson:"customer" json:"customer"`
	Gasstation primitive.ObjectID `bson:"gasstation" json:"gasstation"`
	Active     bool               `bson:"active" json:"active"`
	CreatedAt  time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

type ConcentradorBombas struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	IDConcentradorBomba string             `bson:"idConcentradorBomba" json:"idConcentradorBomba"`
	NomeConcentrador    string             `bson:"nomeConcentrador" json:"nomeConcentrador"`
	Active              bool               `bson:"active" json:"active"`
	Gasstation          primitive.ObjectID `bson:"gasstation" json:"gasstation"`
	CreatedAt           time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

type Customers struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
	Phone string             `bson:"phone" json:"phone"`
}

type Fills struct {
	MongoID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                  string             `bson:"ID" json:"ID"`
	DataHora            time.Time          `bson:"dataHora" json:"dataHora"`
	DataCaptura         time.Time          `bson:"dataCaptura" json:"dataCaptura"`
	GUIDAutomacaoBomba  string             `bson:"GUIDAutomacaoBomba" json:"GUIDAutomacaoBomba"`
	Volume              float64            `bson:"volume" json:"volume"`
	PrecoLitro          float64            `bson:"precoLitro" json:"precoLitro"`
	EncerranteFinal     float64            `bson:"encerranteFinal" json:"encerranteFinal"`
	EncerranteInicial   float64            `bson:"encerranteInicial" json:"encerranteInicial"`
	CodigoBico          string             `bson:"codigoBico" json:"codigoBico"`
	RetornoAutomacao    string             `bson:"retorno_automacao" json:"retorno_automacao"`
	Tempo               int                `bson:"tempo" json:"tempo"`
	IdentificadorI      string             `bson:"identficadorI" json:"identficadorI"`
	IdentificadorII     string             `bson:"identficadorII" json:"identficadorII"`
	Situacao            int                `bson:"situacao" json:"situacao"`
	NumeroAbastecimento int                `bson:"numeroAbastecimento" json:"numeroAbastecimento"`
	NumeroBicoPista     int                `bson:"numeroBicoPista" json:"numeroBicoPista"`
	CustomerID          primitive.ObjectID `bson:"customer" json:"customer"`
	GasStationID        primitive.ObjectID `bson:"gasstation" json:"gasstation"`
	CombustivelID       primitive.ObjectID `bson:"combustivel" json:"combustivel"`
	Checksum            string             `bson:"checksum" json:"checksum"`
	EmUso               string             `bson:"em_uso" json:"em_uso"`
}

type Fuelnozzles struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NumeroBico          string             `bson:"numeroBico" json:"numeroBico"`
	CodigoHexadecimal   string             `bson:"codigoHexadecimal" json:"codigoHexadecimal"`
	Combustivel         primitive.ObjectID `bson:"combustivel" json:"combustivel"`
	IDConcentradorBomba string             `bson:"idConcentradorBomba" json:"idConcentradorBomba"`
	Active              bool               `bson:"active" json:"active"`
	PreviousState       primitive.ObjectID `bson:"previousState,omitempty" json:"previousState,omitempty"`
}

type GasStation struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	ProductGas    []primitive.ObjectID `bson:"productGas" json:"productGas"`
	Name          string               `bson:"name" json:"name"`
	CNPJ          string               `bson:"cnpj" json:"cnpj"`
	CompanyName   string               `bson:"companyName" json:"companyName"`
	Flag          string               `bson:"flag" json:"flag"`
	City          string               `bson:"city" json:"city"`
	CEP           string               `bson:"cep" json:"cep"`
	UF            string               `bson:"uf" json:"uf"`
	Address       string               `bson:"address" json:"address"`
	Phone         string               `bson:"phone" json:"phone"`
	N             string               `bson:"N" json:"N"`
	Latitude      string               `bson:"lat" json:"lat"`
	Longitude     string               `bson:"lng" json:"lng"`
	Image         string               `bson:"image" json:"image"`
	CustomerID    primitive.ObjectID   `bson:"customer" json:"customer"`
	Active        bool                 `bson:"active" json:"active"`
	Neighborhood  string               `bson:"neighborhood" json:"neighborhood"`
	WhatsApp      string               `bson:"whatsapp" json:"whatsapp"`
	Facebook      string               `bson:"facebook" json:"facebook"`
	Instagram     string               `bson:"instagram" json:"instagram"`
	Site          string               `bson:"site" json:"site"`
	CreatedAt     time.Time            `bson:"createdAd" json:"createdAd"`
	Version       int                  `bson:"__v" json:"__v"`
	ClientRated   int                  `bson:"clientrated" json:"clientrated"`
	Evaluation    string               `bson:"evaluation" json:"evaluation"`
	UpdatedAt     time.Time            `bson:"updatedAt" json:"updatedAt"`
	Timezone      int64                `bson:"timezone" json:"timezone"`
	UserUpdateID  primitive.ObjectID   `bson:"userUpdate" json:"userUpdate"`
	FlagImage     string               `bson:"flagImage" json:"flagImage"`
	DenyPayments  []string             `bson:"denyPayments" json:"denyPayments"`
	HiddenInApp   bool                 `bson:"hiddeninApp" json:"hiddeninApp"`
	AtendenteCode []string             `bson:"atendenteCode" json:"atendenteCode"`
}

type ProductGas struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name            string             `bson:"name" json:"name"`
	Icon            string             `bson:"icon" json:"icon"`
	CustomerID      primitive.ObjectID `bson:"customer" json:"customer"`
	Ref             string             `bson:"ref" json:"ref"`
	Active          bool               `bson:"active" json:"active"`
	CreatedAt       time.Time          `bson:"createdAd" json:"createdAd"`
	Version         int                `bson:"__v" json:"__v"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
	UnitID          primitive.ObjectID `bson:"unit" json:"unit"`
	ProductGasGroup *string            `bson:"productgasGroup" json:"productgasGroup"`
}
