package tests

import (
	ssov1 "go-auth-service/protos/gen/sso"
	"go-auth-service/tests/suite"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "test-secret"

	passwordDefaultLen = 10
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	password := gofakeitPassword()

	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    appID,
	})
	require.NoError(t, err)

	loginTime := time.Now()

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, respReg.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, email, claims["email"].(string))
	assert.Equal(t, appID, int(claims["app_id"].(float64)))

	const deltaSeconds = 2.5
	assert.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claims["exp"].(float64), deltaSeconds)
}

func TestRegisterLogin_DuplicatedRegistration(t *testing.T) {
	ctx, st := suite.New(t)

	payload := ssov1.RegisterRequest{
		Email:    gofakeit.Email(),
		Password: gofakeitPassword(),
	}

	respReg, err := st.AuthClient.Register(ctx, &payload)
	require.NoError(t, err)
	require.NotEmpty(t, respReg.GetUserId())

	respReg, err = st.AuthClient.Register(ctx, &payload)
	require.Error(t, err)
	assert.Empty(t, respReg.GetUserId())
	assert.ErrorContains(t, err, "user already exists")
}

func TestRegister_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		email       string
		password    string
		expectedErr string
	}{
		{
			name:     "Register with Empty Password",
			email:    gofakeit.Email(),
			password: "",
		},
		{
			name:     "Register with Empty Email",
			email:    "",
			password: gofakeitPassword(),
		},
		{
			name:     "Register with Empty Email and Password",
			email:    "",
			password: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    test.email,
				Password: test.password,
			})
			require.Error(t, err)
		})
	}
}

func TestLogin_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		email       string
		password    string
		appID       int32
		expectedErr string
	}{
		{
			name:     "Login with Empty Password",
			email:    gofakeit.Email(),
			password: "",
			appID:    appID,
		},
		{
			name:     "Login with Empty Email",
			email:    "",
			password: gofakeitPassword(),
			appID:    appID,
		},
		{
			name:     "Login with Empty Email and Password",
			email:    "",
			password: "",
			appID:    appID,
		},
		{
			name:     "Login with Non-Matching Password",
			email:    gofakeit.Email(),
			password: gofakeitPassword(),
			appID:    appID,
		},
		{
			name:     "Login without AppID",
			email:    gofakeit.Email(),
			password: gofakeitPassword(),
			appID:    emptyAppID,
		},
	}

	for _, test := range tests {
		go t.Run(test.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    gofakeit.Email(),
				Password: gofakeitPassword(),
			})
			require.NoError(t, err)

			_, err = st.AuthClient.Login(ctx, &ssov1.LoginRequest{
				Email:    test.email,
				Password: test.password,
				AppId:    test.appID,
			})
			require.Error(t, err)
		})
	}
}

func gofakeitPassword() string {
	return gofakeit.Password(true, true, true, true, false, passwordDefaultLen)
}
