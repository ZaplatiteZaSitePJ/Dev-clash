package jwt

import jwt_config "dev-clash/pkg/jwt"

type TokenManager interface {
	GenerateTokens(userID int64) (access, reresh string, err error)
	ValidateAccess(token string) (int64, error)
	ValidateRefresh(token string) (int64, error)
}

type JWTTokenManager struct {
	manager *jwt_config.JWTConfig
}

func NewJWTTokenManager (manager *jwt_config.JWTConfig) (*JWTTokenManager){
	return &JWTTokenManager{
		manager: manager,
	}
}

func (m *JWTTokenManager) GenerateTokens(userID int64) (access, reresh string, err error) {
	return m.manager.GenerateTokens(userID)
}

func (m *JWTTokenManager) ValidateAccess(token string) (int64, error) {
	claims, err := m.manager.ValidateAccess(token)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}

func (m *JWTTokenManager) ValidateRefresh(token string) (int64, error) {
	claims, err := m.manager.ValidateRefresh(token)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}