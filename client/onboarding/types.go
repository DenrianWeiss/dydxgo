package onboarding

import (
	"github.com/denrianweiss/dydxgo/client/private"
	"github.com/denrianweiss/dydxgo/types"
)

const SigningMethodTypedData = "TypedData"

type OnboardingResponse struct {
	APIKey  types.ApiKeyCredentials `json:"apiKey"`
	User    private.User            `json:"user"`
	Account private.Account         `json:"account"`
}
