package drivers

import (
	"database/sql"

	"github.com/Unfield/zugang/config"
	"github.com/Unfield/zugang/db"
	"github.com/Unfield/zugang/models"
)

type DefaultStorageProvider struct {
	cfg *config.ZugangConfig
	db.UserStore
	db.ClientStore
	db.TenantStore
	db.TokenStore
	db.SessionStore
	db.IdentifierStore
	db.CredentialsStore
}

func NewDefaultStorageProvider(cfg *config.ZugangConfig) (*DefaultStorageProvider, error) {
	var kv any
	var db *sql.DB

	//:TODO initialice the db and kv

	return &DefaultStorageProvider{
		cfg:              cfg,
		UserStore:        NewUserStore(db, kv),
		ClientStore:      NewClientStore(db, kv),
		TenantStore:      NewTenantStore(db, kv),
		TokenStore:       NewTokenStore(db, kv),
		SessionStore:     NewSessionStore(db, kv),
		IdentifierStore:  NewIdentifierStore(db, kv),
		CredentialsStore: NewCredentialsStore(db, kv),
	}, nil
}

type DefaultUserStore struct {
	db *sql.DB
	kv any
}

func NewUserStore(db *sql.DB, kv any) *DefaultUserStore {
	return &DefaultUserStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultUserStore) CreateUser(user *models.UserModel) error {
	return nil
}

func (s *DefaultUserStore) GetUserByID(tenantID, userID string) (models.UserModel, error) {
	return models.UserModel{}, nil
}

func (s *DefaultUserStore) GetUserByEmail(tenantID, email string) (models.UserModel, error) {
	return models.UserModel{}, nil
}

func (s *DefaultUserStore) UpdateUserProfile(user *models.UserModel) error {
	return nil
}

func (s *DefaultUserStore) ListUsers(tenantID string) ([]models.UserModel, error) {
	return []models.UserModel{}, nil
}

func (s *DefaultUserStore) DeleteUser(tenantID, userID string) error {
	return nil
}

type DefaultClientStore struct {
	db *sql.DB
	kv any
}

func NewClientStore(db *sql.DB, kv any) *DefaultClientStore {
	return &DefaultClientStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultClientStore) RegisterClient(client *models.ClientModel) error {
	return nil
}

func (s *DefaultClientStore) GetClientByID(tenantID, clientID string) (models.ClientModel, error) {
	return models.ClientModel{}, nil
}

func (s *DefaultClientStore) ValidateSecret(tenantID, clientID, secret string) (bool, error) {
	return false, nil
}

func (s *DefaultClientStore) ListClientsByTenant(tenantID string) ([]models.ClientModel, error) {
	return []models.ClientModel{}, nil
}

func (s *DefaultClientStore) RevokeClient(tenantID, clientID string) error {
	return nil
}

type DefaultTenantStore struct {
	db *sql.DB
	kv any
}

func NewTenantStore(db *sql.DB, kv any) *DefaultTenantStore {
	return &DefaultTenantStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultTenantStore) CreateTenant(tenant *models.TenantModel) error {
	return nil
}

func (s *DefaultTenantStore) GetTenantByID(tenantID string) (models.TenantModel, error) {
	return models.TenantModel{}, nil
}

func (s *DefaultTenantStore) GetTenantByIssuer(issuer string) (models.TenantModel, error) {
	return models.TenantModel{}, nil
}

func (s *DefaultTenantStore) UpdateTenantSettings(tenant *models.TenantModel) error {
	return nil
}

func (s *DefaultTenantStore) ListTenants() ([]models.TenantModel, error) {
	return []models.TenantModel{}, nil
}

func (s *DefaultTenantStore) DeleteTenant(tenantID string) error {
	return nil
}

type DefaultTokenStore struct {
	db *sql.DB
	kv any
}

func NewTokenStore(db *sql.DB, kv any) *DefaultTokenStore {
	return &DefaultTokenStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultTokenStore) SaveAccessToken(token *models.TokenModel) error {
	return nil
}

func (s *DefaultTokenStore) SaveRefreshToken(token *models.TokenModel) error {
	return nil
}

func (s *DefaultTokenStore) GetAccessToken(tenantID, tokenValue string) (models.TokenModel, error) {
	return models.TokenModel{}, nil
}

func (s *DefaultTokenStore) GetRefreshToken(tenantID, tokenValue string) (models.TokenModel, error) {
	return models.TokenModel{}, nil
}

func (s *DefaultTokenStore) ValidateAccessToken(tokenID, tokenValue string) (bool, error) {
	return false, nil
}

func (s *DefaultTokenStore) ValidateRefreshToken(tokenID, tokenValue string) (bool, error) {
	return false, nil
}

func (s *DefaultTokenStore) IsTokenRevoked(tokenID, tokenValue string) (bool, error) {
	return false, nil
}

func (s *DefaultTokenStore) RevokeToken(tenantID, tokenValue string) error {
	return nil
}

func (s *DefaultTokenStore) CleanupExpiredTokens() error {
	return nil
}

type DefaultSessionStore struct {
	db *sql.DB
	kv any
}

func NewSessionStore(db *sql.DB, kv any) *DefaultSessionStore {
	return &DefaultSessionStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultSessionStore) SaveSession(session *models.SessionModel) error {
	return nil
}

func (s *DefaultSessionStore) GetSessionByID(tenantID, sessionID string) (models.SessionModel, error) {
	return models.SessionModel{}, nil
}

func (s *DefaultSessionStore) ListSessionsByUser(tenantID, userID string) ([]models.SessionModel, error) {
	return []models.SessionModel{}, nil
}

func (s *DefaultSessionStore) RevokeSession(tenantID, sessionID string) error {
	return nil
}

func (s *DefaultSessionStore) RevokeAllSessionsForUser(tenantID, userID string) error {
	return nil
}

func (s *DefaultSessionStore) ExpireOldSessions() error {
	return nil
}

type DefaultIdentifierStore struct {
	db *sql.DB
	kv any
}

func NewIdentifierStore(db *sql.DB, kv any) *DefaultIdentifierStore {
	return &DefaultIdentifierStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultIdentifierStore) SaveIdentifier(identifier *models.IdentifierModel) error {
	return nil
}

func (s *DefaultIdentifierStore) GetIdentifierByValue(tenantID, value string) (models.IdentifierModel, error) {
	return models.IdentifierModel{}, nil
}

func (s *DefaultIdentifierStore) ListIdentifiersByUser(tenantID, userID string) ([]models.IdentifierModel, error) {
	return []models.IdentifierModel{}, nil
}

func (s *DefaultIdentifierStore) RemoveIdentifier(tenantID, identifierID string) error {
	return nil
}

type DefaultCredentialsStore struct {
	db *sql.DB
	kv any
}

func NewCredentialsStore(db *sql.DB, kv any) *DefaultCredentialsStore {
	return &DefaultCredentialsStore{
		db: db,
		kv: kv,
	}
}

func (s *DefaultCredentialsStore) SaveCredentials(credentials *models.CredentialsModel) error {
	return nil
}

func (s *DefaultCredentialsStore) GetCredentialsByIdentifierID(tenantID, identifierID string) (models.CredentialsModel, error) {
	return models.CredentialsModel{}, nil
}

func (s *DefaultCredentialsStore) GetAllCredentialsForUser(tenantID, userID string) ([]models.CredentialsModel, error) {
	return []models.CredentialsModel{}, nil
}

func (s *DefaultCredentialsStore) UpdateCredentials(credentials *models.CredentialsModel) error {
	return nil
}

func (s *DefaultCredentialsStore) DeleteCredentials(tenantID, credentialsID string) error {
	return nil
}
