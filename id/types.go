package id

type Type string

const (
	TypeTenant     Type = "tnt"
	TypeUser       Type = "usr"
	TypeIdentifier Type = "idn"
	TypeCredential Type = "crd"
	TypeClient     Type = "clt"
	TypeToken      Type = "tok"
	TypeSession    Type = "ses"
	TypeAudit      Type = "evt"
)

const ID_LENGTH = 24
