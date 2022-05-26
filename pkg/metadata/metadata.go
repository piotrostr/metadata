package metadata

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/piotrostr/metadata/pkg/db"
)

const DESCRIPTION = "SMPLverse is a collection of synthetic face data from the computational infrastructure of the metaverse, assigned to minters using facial recognition."

const PLACEHOLDER_IMAGE = "ipfs://QmYypT49WH7rYTL2jXpfoNH2DAMHe9VM7pwwEjUVr45XK1"

var ctx = context.Background()

var BlankEntry = Entry{
	TokenId:     "#",
	Name:        "UNCLAIMED SMPL",
	Description: DESCRIPTION,
	ExternalUrl: "",
	Image:       PLACEHOLDER_IMAGE,
	Attributes:  []Attribute(nil),
}

type Metadata struct {
	entries map[string]Entry
	rdb     *redis.Client
}

type Entry struct {
	TokenId     string      `json:"token_id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	ExternalUrl string      `json:"external_url,omitempty"`
	Image       string      `json:"image,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	TraitType string `json:"trait_type,omitempty"`
	Value     string `json:"value,omitempty"`
}

var _ = []string{
	"037544",
	"069701",
	"099370",
	"093321",
	"051039",
	"046594",
	"059759",
	"074727",
	"083824",
	"037661",
	"059324",
}

func New() *Metadata {
	return &Metadata{
		entries: make(map[string]Entry),
		rdb:     db.Client(),
	}
}

func (m *Metadata) Get(tokenId string) *Entry {
	// TODO add cap of tokenId <= totalSupply
	if entry, ok := m.entries[tokenId]; ok {
		return &entry
	}
	return &BlankEntry
}

func (m *Metadata) Add(tokenId string, entry Entry) {
	m.entries[tokenId] = entry
}
