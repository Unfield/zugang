package db

import "github.com/Unfield/zugang/models"

type StorageProvider interface {
	UserStore
	ClientStore
	TenantStore
	TokenStore
	SessionStore
	IdentifierStore
	CredentialsStore
}

type UserStore interface {
	CreateUser(user *models.UserModel) error

	GetUserByID(tenantID, userID string) (models.UserModel, error)
	GetUserByEmail(tenantID, email string) (models.UserModel, error)

	UpdateUserProfile(user *models.UserModel) error

	ListUsers(tenantID string) ([]models.UserModel, error)

	DeleteUser(tenantID, userID string) error
}

type ClientStore interface {
	RegisterClient(client *models.ClientModel) error

	GetClientByID(tenantID, clientID string) (models.ClientModel, error)

	ValidateSecret(tenantID, clientID, secret string) (bool, error)

	ListClientsByTenant(tenantID string) ([]models.ClientModel, error)

	RevokeClient(tenantID, clientID string) error
}

type TenantStore interface {
	CreateTenant(tenant *models.TenantModel) error

	GetTenantByID(tenantID string) (models.TenantModel, error)
	GetTenantByIssuer(issuer string) (models.TenantModel, error)

	UpdateTenantSettings(tenant *models.TenantModel) error

	ListTenants() ([]models.TenantModel, error)

	DeleteTenant(tenantID string) error
}

type TokenStore interface {
	SaveAccessToken(token *models.TokenModel) error
	SaveRefreshToken(token *models.TokenModel) error

	GetAccessToken(tenantID, tokenValue string) (models.TokenModel, error)
	GetRefreshToken(tenantID, tokenValue string) (models.TokenModel, error)

	ValidateAccessToken(tenantID, tokenValue string) (bool, error)
	ValidateRefreshToken(tenantID, tokenValue string) (bool, error)

	IsTokenRevoked(tenantID, tokenValue string) (bool, error)

	RevokeToken(tenantID, tokenValue string) error

	CleanupExpiredTokens() error
}

type SessionStore interface {
	SaveSession(session *models.SessionModel) error

	GetSessionByID(tenantID, sessionID string) (models.SessionModel, error)

	ListSessionsByUser(tenantID, userID string) ([]models.SessionModel, error)

	RevokeSession(tenantID, sessionID string) error
	RevokeAllSessionsForUser(tenantID, userID string) error

	ExpireOldSessions() error
}

type IdentifierStore interface {
	SaveIdentifier(identifier *models.IdentifierModel) error

	GetIdentifierByValue(tenantID, value string) (models.IdentifierModel, error)

	ListIdentifiersByUser(tenantID, userID string) ([]models.IdentifierModel, error)

	RemoveIdentifier(tenantID, identifierID string) error
}

type CredentialsStore interface {
	SaveCredentials(credentials *models.CredentialsModel) error

	GetCredentialsByIdentifierID(tenantID, identifierID string) (models.CredentialsModel, error)
	GetAllCredentialsForUser(tenantID, userID string) ([]models.CredentialsModel, error)

	UpdateCredentials(credentials *models.CredentialsModel) error

	DeleteCredentials(tenantID, credentialsID string) error
}
