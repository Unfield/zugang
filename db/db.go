package db

import "github.com/Unfield/zugang/models"

type StorageProvider interface {
	UserStore
	ClientStore
	TenantStore
	TokenStore
	SessionStore
}

type UserStore interface {
	CreateUser(user *models.UserModel) error

	GetUserByID(id string) (models.UserModel, error)
	GetUserByEmail(email string) (models.UserModel, error)

	UpdateUserProfile(user *models.UserModel) error

	ListUsers() ([]models.UserModel, error)

	DeleteUser(id string) error
}

type ClientStore interface {
	RegisterClient(client *models.ClientModel) error

	GetClientByID(id string) (models.ClientModel, error)

	ValidateSecret(id, secret string) (bool, error)

	ListClientsByTenant(tenantID string) ([]models.ClientModel, error)

	RevokeClient(id string) error
}

type TenantStore interface {
	CreateTenant(tenant *models.TenantModel) error

	GetTenantByID(id string) (models.TenantModel, error)
	GetTenantByIssuer(issuer string) (models.TenantModel, error)

	UpdateTenantSettings(tenant *models.TenantModel) error

	ListTenants() ([]models.TenantModel, error)

	DeleteTenant(id string) error
}

type TokenStore interface {
	SaveAccessToken(token *models.TokenModel) error
	SaveRefreshToken(token *models.TokenModel) error

	GetAccessToken(tokenValue string) (models.TokenModel, error)
	GetRefreshToken(tokenValue string) (models.TokenModel, error)

	ValidateAccessToken(tokenValue string) (bool, error)
	ValidateRefreshToken(tokenValue string) (bool, error)

	IsTokenRevoked(tokenValue string) (bool, error)

	RevokeToken(tokenValue string) error

	CleanupExpiredTokens()
}

type SessionStore interface {
	SaveSession(session *models.SessionModel) error

	GetSessionByID(id string) (models.SessionModel, error)

	ListSessionsByUser(userID string) ([]models.SessionModel, error)

	RevokeSession(id string) error
	RevokeAllSessionsForUser(userID string) error

	ExpireOldSessions()
}

type IdentifierStore interface {
	SaveIdentifier(identifier *models.IdentifierModel) error

	GetIdentifierByValue(value string) (models.IdentifierModel, error)

	ListIdentifiersByUser(userID string) ([]models.IdentifierModel, error)

	RemoveIdentifier(id string) error
}

type CredentialsStore interface {
}
