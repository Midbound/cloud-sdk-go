// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package midboundcloud

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"

	"github.com/Midbound/cloud-sdk-go/internal/apijson"
	"github.com/Midbound/cloud-sdk-go/internal/requestconfig"
	"github.com/Midbound/cloud-sdk-go/option"
	"github.com/Midbound/cloud-sdk-go/packages/respjson"
	"github.com/Midbound/cloud-sdk-go/shared/constant"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookService contains methods and other services that help with interacting
// with the midbound-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookSecret
	if key == "" {
		return nil, errors.New("The WebhookSecret option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *WebhookService) UnwrapUnsafe(payload []byte, opts ...option.RequestOption) (*UnwrapUnsafeWebhookEventUnion, error) {
	res := &UnwrapUnsafeWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type IdentityResolvedWebhookEvent struct {
	// Unique event identifier
	ID string `json:"id" api:"required"`
	// Unix timestamp (milliseconds) when the event was created
	Created int64                            `json:"created" api:"required"`
	Data    IdentityResolvedWebhookEventData `json:"data" api:"required"`
	// Event type identifier
	Type constant.IdentityResolved `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventData struct {
	Attribution IdentityResolvedWebhookEventDataAttribution `json:"attribution" api:"required"`
	Identity    IdentityResolvedWebhookEventDataIdentity    `json:"identity" api:"required"`
	Session     IdentityResolvedWebhookEventDataSession     `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribution respjson.Field
		Identity    respjson.Field
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataAttribution struct {
	PixelID string `json:"pixelId" api:"required"`
	// Any of "jNNdd", "OMzN4".
	Prid      string `json:"prid" api:"required"`
	SessionID string `json:"sessionId" api:"required"`
	// Any of "website".
	Type string `json:"type" api:"required"`
	Vid  string `json:"vid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PixelID     respjson.Field
		Prid        respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		Vid         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataAttribution) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataAttribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataIdentity struct {
	Demographics IdentityResolvedWebhookEventDataIdentityDemographics `json:"demographics" api:"required"`
	Emails       []IdentityResolvedWebhookEventDataIdentityEmail      `json:"emails" api:"required"`
	LinkedinURL  string                                               `json:"linkedinUrl" api:"required"`
	Location     IdentityResolvedWebhookEventDataIdentityLocation     `json:"location" api:"required"`
	Phones       []string                                             `json:"phones" api:"required"`
	Professional IdentityResolvedWebhookEventDataIdentityProfessional `json:"professional" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Demographics respjson.Field
		Emails       respjson.Field
		LinkedinURL  respjson.Field
		Location     respjson.Field
		Phones       respjson.Field
		Professional respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataIdentity) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataIdentity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataIdentityDemographics struct {
	FirstName   string `json:"firstName" api:"required"`
	HasChildren bool   `json:"hasChildren" api:"required"`
	IsHomeowner bool   `json:"isHomeowner" api:"required"`
	IsMarried   bool   `json:"isMarried" api:"required"`
	LastName    string `json:"lastName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		HasChildren respjson.Field
		IsHomeowner respjson.Field
		IsMarried   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataIdentityDemographics) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataIdentityDemographics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataIdentityEmail struct {
	// Any of "personal", "professional".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataIdentityEmail) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataIdentityEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataIdentityLocation struct {
	City    string `json:"city" api:"required"`
	Country string `json:"country" api:"required"`
	State   string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataIdentityLocation) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataIdentityLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataIdentityProfessional struct {
	CompanyName string `json:"companyName" api:"required"`
	JobTitle    string `json:"jobTitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyName respjson.Field
		JobTitle    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataIdentityProfessional) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataIdentityProfessional) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataSession struct {
	CreatedAt    float64                                        `json:"createdAt" api:"required"`
	EndedAt      float64                                        `json:"endedAt" api:"required"`
	Fbclid       string                                         `json:"fbclid" api:"required"`
	Gclid        string                                         `json:"gclid" api:"required"`
	LandingPage  string                                         `json:"landingPage" api:"required"`
	LandingTitle string                                         `json:"landingTitle" api:"required"`
	Network      IdentityResolvedWebhookEventDataSessionNetwork `json:"network" api:"required"`
	Pid          string                                         `json:"pid" api:"required"`
	Referrer     string                                         `json:"referrer" api:"required"`
	Screen       IdentityResolvedWebhookEventDataSessionScreen  `json:"screen" api:"required"`
	Sid          string                                         `json:"sid" api:"required"`
	Tenant       string                                         `json:"tenant" api:"required"`
	Timezone     string                                         `json:"timezone" api:"required"`
	UserAgent    string                                         `json:"userAgent" api:"required"`
	Utm          IdentityResolvedWebhookEventDataSessionUtm     `json:"utm" api:"required"`
	Vid          string                                         `json:"vid" api:"required"`
	Options      map[string]any                                 `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataSession) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataSessionNetwork struct {
	Asn           IdentityResolvedWebhookEventDataSessionNetworkAsn           `json:"asn" api:"required"`
	BotManagement IdentityResolvedWebhookEventDataSessionNetworkBotManagement `json:"botManagement" api:"required"`
	City          string                                                      `json:"city" api:"required"`
	Colo          string                                                      `json:"colo" api:"required"`
	Continent     string                                                      `json:"continent" api:"required"`
	Country       string                                                      `json:"country" api:"required"`
	IP            string                                                      `json:"ip" api:"required"`
	IsEu          bool                                                        `json:"isEU" api:"required"`
	Latitude      string                                                      `json:"latitude" api:"required"`
	Longitude     string                                                      `json:"longitude" api:"required"`
	MetroCode     string                                                      `json:"metroCode" api:"required"`
	PostalCode    string                                                      `json:"postalCode" api:"required"`
	Region        string                                                      `json:"region" api:"required"`
	RegionCode    string                                                      `json:"regionCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataSessionNetwork) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataSessionNetworkAsn struct {
	Code float64 `json:"code" api:"required"`
	Name string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataSessionNetworkAsn) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy" api:"required"`
	Score               float64 `json:"score" api:"required"`
	VerifiedBot         bool    `json:"verifiedBot" api:"required"`
	VerifiedBotCategory string  `json:"verifiedBotCategory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataSessionNetworkBotManagement) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityResolvedWebhookEventDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataSessionScreen struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataSessionScreen) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityResolvedWebhookEventDataSessionUtm struct {
	Campaign string `json:"campaign" api:"required"`
	Content  string `json:"content" api:"required"`
	Medium   string `json:"medium" api:"required"`
	Source   string `json:"source" api:"required"`
	Term     string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityResolvedWebhookEventDataSessionUtm) RawJSON() string { return r.JSON.raw }
func (r *IdentityResolvedWebhookEventDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEvent struct {
	// Unique event identifier
	ID string `json:"id" api:"required"`
	// Unix timestamp (milliseconds) when the event was created
	Created int64                             `json:"created" api:"required"`
	Data    IdentityQualifiedWebhookEventData `json:"data" api:"required"`
	// Event type identifier
	Type constant.IdentityQualified `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventData struct {
	Attribution IdentityQualifiedWebhookEventDataAttribution `json:"attribution" api:"required"`
	Identity    IdentityQualifiedWebhookEventDataIdentity    `json:"identity" api:"required"`
	Session     IdentityQualifiedWebhookEventDataSession     `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribution respjson.Field
		Identity    respjson.Field
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataAttribution struct {
	PixelID string `json:"pixelId" api:"required"`
	// Any of "jNNdd", "OMzN4".
	Prid      string `json:"prid" api:"required"`
	SessionID string `json:"sessionId" api:"required"`
	// Any of "website".
	Type string `json:"type" api:"required"`
	Vid  string `json:"vid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PixelID     respjson.Field
		Prid        respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		Vid         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataAttribution) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataAttribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataIdentity struct {
	Demographics IdentityQualifiedWebhookEventDataIdentityDemographics `json:"demographics" api:"required"`
	Emails       []IdentityQualifiedWebhookEventDataIdentityEmail      `json:"emails" api:"required"`
	LinkedinURL  string                                                `json:"linkedinUrl" api:"required"`
	Location     IdentityQualifiedWebhookEventDataIdentityLocation     `json:"location" api:"required"`
	Phones       []string                                              `json:"phones" api:"required"`
	Professional IdentityQualifiedWebhookEventDataIdentityProfessional `json:"professional" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Demographics respjson.Field
		Emails       respjson.Field
		LinkedinURL  respjson.Field
		Location     respjson.Field
		Phones       respjson.Field
		Professional respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataIdentity) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataIdentity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataIdentityDemographics struct {
	FirstName   string `json:"firstName" api:"required"`
	HasChildren bool   `json:"hasChildren" api:"required"`
	IsHomeowner bool   `json:"isHomeowner" api:"required"`
	IsMarried   bool   `json:"isMarried" api:"required"`
	LastName    string `json:"lastName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstName   respjson.Field
		HasChildren respjson.Field
		IsHomeowner respjson.Field
		IsMarried   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataIdentityDemographics) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataIdentityDemographics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataIdentityEmail struct {
	// Any of "personal", "professional".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataIdentityEmail) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataIdentityEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataIdentityLocation struct {
	City    string `json:"city" api:"required"`
	Country string `json:"country" api:"required"`
	State   string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataIdentityLocation) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataIdentityLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataIdentityProfessional struct {
	CompanyName string `json:"companyName" api:"required"`
	JobTitle    string `json:"jobTitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyName respjson.Field
		JobTitle    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataIdentityProfessional) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataIdentityProfessional) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataSession struct {
	CreatedAt    float64                                         `json:"createdAt" api:"required"`
	EndedAt      float64                                         `json:"endedAt" api:"required"`
	Fbclid       string                                          `json:"fbclid" api:"required"`
	Gclid        string                                          `json:"gclid" api:"required"`
	LandingPage  string                                          `json:"landingPage" api:"required"`
	LandingTitle string                                          `json:"landingTitle" api:"required"`
	Network      IdentityQualifiedWebhookEventDataSessionNetwork `json:"network" api:"required"`
	Pid          string                                          `json:"pid" api:"required"`
	Referrer     string                                          `json:"referrer" api:"required"`
	Screen       IdentityQualifiedWebhookEventDataSessionScreen  `json:"screen" api:"required"`
	Sid          string                                          `json:"sid" api:"required"`
	Tenant       string                                          `json:"tenant" api:"required"`
	Timezone     string                                          `json:"timezone" api:"required"`
	UserAgent    string                                          `json:"userAgent" api:"required"`
	Utm          IdentityQualifiedWebhookEventDataSessionUtm     `json:"utm" api:"required"`
	Vid          string                                          `json:"vid" api:"required"`
	Options      map[string]any                                  `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataSession) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataSessionNetwork struct {
	Asn           IdentityQualifiedWebhookEventDataSessionNetworkAsn           `json:"asn" api:"required"`
	BotManagement IdentityQualifiedWebhookEventDataSessionNetworkBotManagement `json:"botManagement" api:"required"`
	City          string                                                       `json:"city" api:"required"`
	Colo          string                                                       `json:"colo" api:"required"`
	Continent     string                                                       `json:"continent" api:"required"`
	Country       string                                                       `json:"country" api:"required"`
	IP            string                                                       `json:"ip" api:"required"`
	IsEu          bool                                                         `json:"isEU" api:"required"`
	Latitude      string                                                       `json:"latitude" api:"required"`
	Longitude     string                                                       `json:"longitude" api:"required"`
	MetroCode     string                                                       `json:"metroCode" api:"required"`
	PostalCode    string                                                       `json:"postalCode" api:"required"`
	Region        string                                                       `json:"region" api:"required"`
	RegionCode    string                                                       `json:"regionCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataSessionNetwork) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataSessionNetworkAsn struct {
	Code float64 `json:"code" api:"required"`
	Name string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataSessionNetworkAsn) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy" api:"required"`
	Score               float64 `json:"score" api:"required"`
	VerifiedBot         bool    `json:"verifiedBot" api:"required"`
	VerifiedBotCategory string  `json:"verifiedBotCategory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataSessionNetworkBotManagement) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityQualifiedWebhookEventDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataSessionScreen struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataSessionScreen) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityQualifiedWebhookEventDataSessionUtm struct {
	Campaign string `json:"campaign" api:"required"`
	Content  string `json:"content" api:"required"`
	Medium   string `json:"medium" api:"required"`
	Source   string `json:"source" api:"required"`
	Term     string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityQualifiedWebhookEventDataSessionUtm) RawJSON() string { return r.JSON.raw }
func (r *IdentityQualifiedWebhookEventDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEvent struct {
	// Unique event identifier
	ID string `json:"id" api:"required"`
	// Unix timestamp (milliseconds) when the event was created
	Created int64                            `json:"created" api:"required"`
	Data    IdentityEnrichedWebhookEventData `json:"data" api:"required"`
	// Event type identifier
	Type constant.IdentityEnriched `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventData struct {
	Attribution       IdentityEnrichedWebhookEventDataAttribution `json:"attribution" api:"required"`
	CompaniesEnriched float64                                     `json:"companiesEnriched" api:"required"`
	Enrichment        IdentityEnrichedWebhookEventDataEnrichment  `json:"enrichment" api:"required"`
	Query             IdentityEnrichedWebhookEventDataQuery       `json:"query" api:"required"`
	Session           IdentityEnrichedWebhookEventDataSession     `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribution       respjson.Field
		CompaniesEnriched respjson.Field
		Enrichment        respjson.Field
		Query             respjson.Field
		Session           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataAttribution struct {
	PixelID string `json:"pixelId" api:"required"`
	// Any of "jNNdd", "OMzN4".
	Prid      string `json:"prid" api:"required"`
	SessionID string `json:"sessionId" api:"required"`
	// Any of "website".
	Type string `json:"type" api:"required"`
	Vid  string `json:"vid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PixelID     respjson.Field
		Prid        respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		Vid         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataAttribution) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataAttribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichment struct {
	Companies        []IdentityEnrichedWebhookEventDataEnrichmentCompany         `json:"companies" api:"required"`
	Person           IdentityEnrichedWebhookEventDataEnrichmentPerson            `json:"person" api:"required"`
	EmailValidations []IdentityEnrichedWebhookEventDataEnrichmentEmailValidation `json:"emailValidations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Companies        respjson.Field
		Person           respjson.Field
		EmailValidations respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichment) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompany struct {
	// Any of "full".
	Type       string `json:"_type" api:"required"`
	EnrichedAt string `json:"enrichedAt" api:"required"`
	Name       string `json:"name" api:"required"`
	// Any of "xK9mP", "qR7nL".
	Provider          string                                                             `json:"provider" api:"required"`
	Address           IdentityEnrichedWebhookEventDataEnrichmentCompanyAddress           `json:"address" api:"nullable"`
	Description       string                                                             `json:"description" api:"nullable"`
	Domain            string                                                             `json:"domain" api:"nullable"`
	EmployeeCount     float64                                                            `json:"employeeCount" api:"nullable"`
	EstimatedRevenue  IdentityEnrichedWebhookEventDataEnrichmentCompanyEstimatedRevenue  `json:"estimatedRevenue" api:"nullable"`
	FoundedYear       float64                                                            `json:"foundedYear" api:"nullable"`
	Funding           IdentityEnrichedWebhookEventDataEnrichmentCompanyFunding           `json:"funding" api:"nullable"`
	Headcount         IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcount         `json:"headcount" api:"nullable"`
	Industry          string                                                             `json:"industry" api:"nullable"`
	LinkedinFollowers IdentityEnrichedWebhookEventDataEnrichmentCompanyLinkedinFollowers `json:"linkedinFollowers" api:"nullable"`
	LinkedinID        string                                                             `json:"linkedinId" api:"nullable"`
	LinkedinURL       string                                                             `json:"linkedinUrl" api:"nullable" format:"uri"`
	LogoURL           string                                                             `json:"logoUrl" api:"nullable" format:"uri"`
	Seo               IdentityEnrichedWebhookEventDataEnrichmentCompanySeo               `json:"seo" api:"nullable"`
	Specialties       []string                                                           `json:"specialties"`
	Taxonomy          IdentityEnrichedWebhookEventDataEnrichmentCompanyTaxonomy          `json:"taxonomy" api:"nullable"`
	Website           string                                                             `json:"website" api:"nullable" format:"uri"`
	WebTraffic        IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTraffic        `json:"webTraffic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		EnrichedAt        respjson.Field
		Name              respjson.Field
		Provider          respjson.Field
		Address           respjson.Field
		Description       respjson.Field
		Domain            respjson.Field
		EmployeeCount     respjson.Field
		EstimatedRevenue  respjson.Field
		FoundedYear       respjson.Field
		Funding           respjson.Field
		Headcount         respjson.Field
		Industry          respjson.Field
		LinkedinFollowers respjson.Field
		LinkedinID        respjson.Field
		LinkedinURL       respjson.Field
		LogoURL           respjson.Field
		Seo               respjson.Field
		Specialties       respjson.Field
		Taxonomy          respjson.Field
		Website           respjson.Field
		WebTraffic        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompany) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyAddress struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyAddress) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyEstimatedRevenue struct {
	Currency string  `json:"currency"`
	Max      float64 `json:"max" api:"nullable"`
	Min      float64 `json:"min" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyEstimatedRevenue) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyEstimatedRevenue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyFunding struct {
	DaysSinceLastRound float64  `json:"daysSinceLastRound" api:"nullable"`
	Investors          []string `json:"investors"`
	LastRoundAmountUsd float64  `json:"lastRoundAmountUsd" api:"nullable"`
	LastRoundType      string   `json:"lastRoundType" api:"nullable"`
	TotalRaisedUsd     float64  `json:"totalRaisedUsd" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysSinceLastRound respjson.Field
		Investors          respjson.Field
		LastRoundAmountUsd respjson.Field
		LastRoundType      respjson.Field
		TotalRaisedUsd     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyFunding) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyFunding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcount struct {
	Growth IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcountGrowth `json:"growth" api:"nullable"`
	Total  float64                                                          `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Growth      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcount) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcountGrowth struct {
	Mom float64 `json:"mom" api:"nullable"`
	Qoq float64 `json:"qoq" api:"nullable"`
	Yoy float64 `json:"yoy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mom         respjson.Field
		Qoq         respjson.Field
		Yoy         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcountGrowth) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyHeadcountGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyLinkedinFollowers struct {
	Count            float64 `json:"count" api:"nullable"`
	MomGrowthPercent float64 `json:"momGrowthPercent" api:"nullable"`
	QoqGrowthPercent float64 `json:"qoqGrowthPercent" api:"nullable"`
	YoyGrowthPercent float64 `json:"yoyGrowthPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count            respjson.Field
		MomGrowthPercent respjson.Field
		QoqGrowthPercent respjson.Field
		YoyGrowthPercent respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyLinkedinFollowers) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyLinkedinFollowers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanySeo struct {
	AverageAdRank        float64 `json:"averageAdRank" api:"nullable"`
	AverageOrganicRank   float64 `json:"averageOrganicRank" api:"nullable"`
	MonthlyAdsSpend      float64 `json:"monthlyAdsSpend" api:"nullable"`
	MonthlyOrganicClicks float64 `json:"monthlyOrganicClicks" api:"nullable"`
	MonthlyOrganicValue  float64 `json:"monthlyOrganicValue" api:"nullable"`
	MonthlyPaidClicks    float64 `json:"monthlyPaidClicks" api:"nullable"`
	TotalOrganicKeywords float64 `json:"totalOrganicKeywords" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageAdRank        respjson.Field
		AverageOrganicRank   respjson.Field
		MonthlyAdsSpend      respjson.Field
		MonthlyOrganicClicks respjson.Field
		MonthlyOrganicValue  respjson.Field
		MonthlyPaidClicks    respjson.Field
		TotalOrganicKeywords respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanySeo) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanySeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyTaxonomy struct {
	LinkedinIndustry    string   `json:"linkedinIndustry" api:"nullable"`
	LinkedinSpecialties []string `json:"linkedinSpecialties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedinIndustry    respjson.Field
		LinkedinSpecialties respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyTaxonomy) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTraffic struct {
	MomGrowthPercent float64                                                                   `json:"momGrowthPercent" api:"nullable"`
	MonthlyVisitors  float64                                                                   `json:"monthlyVisitors" api:"nullable"`
	QoqGrowthPercent float64                                                                   `json:"qoqGrowthPercent" api:"nullable"`
	TrafficSources   IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources `json:"trafficSources" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MomGrowthPercent respjson.Field
		MonthlyVisitors  respjson.Field
		QoqGrowthPercent respjson.Field
		TrafficSources   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTraffic) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources struct {
	DirectPercent       float64 `json:"directPercent" api:"nullable"`
	PaidReferralPercent float64 `json:"paidReferralPercent" api:"nullable"`
	ReferralPercent     float64 `json:"referralPercent" api:"nullable"`
	SearchPercent       float64 `json:"searchPercent" api:"nullable"`
	SocialPercent       float64 `json:"socialPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirectPercent       respjson.Field
		PaidReferralPercent respjson.Field
		ReferralPercent     respjson.Field
		SearchPercent       respjson.Field
		SocialPercent       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPerson struct {
	EnrichedAt  string `json:"enrichedAt" api:"required"`
	LinkedinURL string `json:"linkedinUrl" api:"required" format:"uri"`
	// Any of "xK9mP", "qR7nL".
	Provider          string                                                       `json:"provider" api:"required"`
	Connections       float64                                                      `json:"connections" api:"nullable"`
	Education         []IdentityEnrichedWebhookEventDataEnrichmentPersonEducation  `json:"education"`
	Email             string                                                       `json:"email" api:"nullable" format:"email"`
	Employments       []IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment `json:"employments"`
	Experience        []IdentityEnrichedWebhookEventDataEnrichmentPersonExperience `json:"experience"`
	FirstName         string                                                       `json:"firstName" api:"nullable"`
	FullName          string                                                       `json:"fullName" api:"nullable"`
	Headline          string                                                       `json:"headline" api:"nullable"`
	Languages         []string                                                     `json:"languages"`
	LastName          string                                                       `json:"lastName" api:"nullable"`
	LinkedinID        string                                                       `json:"linkedinId" api:"nullable"`
	Location          IdentityEnrichedWebhookEventDataEnrichmentPersonLocation     `json:"location" api:"nullable"`
	ProfilePictureURL string                                                       `json:"profilePictureUrl" api:"nullable" format:"uri"`
	Skills            []string                                                     `json:"skills"`
	Summary           string                                                       `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnrichedAt        respjson.Field
		LinkedinURL       respjson.Field
		Provider          respjson.Field
		Connections       respjson.Field
		Education         respjson.Field
		Email             respjson.Field
		Employments       respjson.Field
		Experience        respjson.Field
		FirstName         respjson.Field
		FullName          respjson.Field
		Headline          respjson.Field
		Languages         respjson.Field
		LastName          respjson.Field
		LinkedinID        respjson.Field
		Location          respjson.Field
		ProfilePictureURL respjson.Field
		Skills            respjson.Field
		Summary           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPerson) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichmentPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEducation struct {
	InstituteName       string `json:"instituteName" api:"required"`
	Degree              string `json:"degree" api:"nullable"`
	EndDate             string `json:"endDate" api:"nullable"`
	FieldOfStudy        string `json:"fieldOfStudy" api:"nullable"`
	InstituteLinkedinID string `json:"instituteLinkedinId" api:"nullable"`
	InstituteLogoURL    string `json:"instituteLogoUrl" api:"nullable" format:"uri"`
	StartDate           string `json:"startDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstituteName       respjson.Field
		Degree              respjson.Field
		EndDate             respjson.Field
		FieldOfStudy        respjson.Field
		InstituteLinkedinID respjson.Field
		InstituteLogoURL    respjson.Field
		StartDate           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEducation) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment struct {
	Company     IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompany `json:"company" api:"required"`
	Title       string                                                            `json:"title" api:"required"`
	Description string                                                            `json:"description" api:"nullable"`
	EndDate     string                                                            `json:"endDate" api:"nullable"`
	Location    string                                                            `json:"location" api:"nullable"`
	StartDate   string                                                            `json:"startDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Title       respjson.Field
		Description respjson.Field
		EndDate     respjson.Field
		Location    respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompany struct {
	// Any of "full".
	Type       string `json:"_type" api:"required"`
	EnrichedAt string `json:"enrichedAt" api:"required"`
	Name       string `json:"name" api:"required"`
	// Any of "xK9mP", "qR7nL".
	Provider          string                                                                             `json:"provider" api:"required"`
	Address           IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress           `json:"address" api:"nullable"`
	Description       string                                                                             `json:"description" api:"nullable"`
	Domain            string                                                                             `json:"domain" api:"nullable"`
	EmployeeCount     float64                                                                            `json:"employeeCount" api:"nullable"`
	EstimatedRevenue  IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue  `json:"estimatedRevenue" api:"nullable"`
	FoundedYear       float64                                                                            `json:"foundedYear" api:"nullable"`
	Funding           IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding           `json:"funding" api:"nullable"`
	Headcount         IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount         `json:"headcount" api:"nullable"`
	Industry          string                                                                             `json:"industry" api:"nullable"`
	LinkedinFollowers IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers `json:"linkedinFollowers" api:"nullable"`
	LinkedinID        string                                                                             `json:"linkedinId" api:"nullable"`
	LinkedinURL       string                                                                             `json:"linkedinUrl" api:"nullable" format:"uri"`
	LogoURL           string                                                                             `json:"logoUrl" api:"nullable" format:"uri"`
	Seo               IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanySeo               `json:"seo" api:"nullable"`
	Specialties       []string                                                                           `json:"specialties"`
	Taxonomy          IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy          `json:"taxonomy" api:"nullable"`
	Website           string                                                                             `json:"website" api:"nullable" format:"uri"`
	WebTraffic        IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic        `json:"webTraffic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		EnrichedAt        respjson.Field
		Name              respjson.Field
		Provider          respjson.Field
		Address           respjson.Field
		Description       respjson.Field
		Domain            respjson.Field
		EmployeeCount     respjson.Field
		EstimatedRevenue  respjson.Field
		FoundedYear       respjson.Field
		Funding           respjson.Field
		Headcount         respjson.Field
		Industry          respjson.Field
		LinkedinFollowers respjson.Field
		LinkedinID        respjson.Field
		LinkedinURL       respjson.Field
		LogoURL           respjson.Field
		Seo               respjson.Field
		Specialties       respjson.Field
		Taxonomy          respjson.Field
		Website           respjson.Field
		WebTraffic        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompany) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue struct {
	Currency string  `json:"currency"`
	Max      float64 `json:"max" api:"nullable"`
	Min      float64 `json:"min" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding struct {
	DaysSinceLastRound float64  `json:"daysSinceLastRound" api:"nullable"`
	Investors          []string `json:"investors"`
	LastRoundAmountUsd float64  `json:"lastRoundAmountUsd" api:"nullable"`
	LastRoundType      string   `json:"lastRoundType" api:"nullable"`
	TotalRaisedUsd     float64  `json:"totalRaisedUsd" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysSinceLastRound respjson.Field
		Investors          respjson.Field
		LastRoundAmountUsd respjson.Field
		LastRoundType      respjson.Field
		TotalRaisedUsd     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount struct {
	Growth IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth `json:"growth" api:"nullable"`
	Total  float64                                                                          `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Growth      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth struct {
	Mom float64 `json:"mom" api:"nullable"`
	Qoq float64 `json:"qoq" api:"nullable"`
	Yoy float64 `json:"yoy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mom         respjson.Field
		Qoq         respjson.Field
		Yoy         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers struct {
	Count            float64 `json:"count" api:"nullable"`
	MomGrowthPercent float64 `json:"momGrowthPercent" api:"nullable"`
	QoqGrowthPercent float64 `json:"qoqGrowthPercent" api:"nullable"`
	YoyGrowthPercent float64 `json:"yoyGrowthPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count            respjson.Field
		MomGrowthPercent respjson.Field
		QoqGrowthPercent respjson.Field
		YoyGrowthPercent respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanySeo struct {
	AverageAdRank        float64 `json:"averageAdRank" api:"nullable"`
	AverageOrganicRank   float64 `json:"averageOrganicRank" api:"nullable"`
	MonthlyAdsSpend      float64 `json:"monthlyAdsSpend" api:"nullable"`
	MonthlyOrganicClicks float64 `json:"monthlyOrganicClicks" api:"nullable"`
	MonthlyOrganicValue  float64 `json:"monthlyOrganicValue" api:"nullable"`
	MonthlyPaidClicks    float64 `json:"monthlyPaidClicks" api:"nullable"`
	TotalOrganicKeywords float64 `json:"totalOrganicKeywords" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageAdRank        respjson.Field
		AverageOrganicRank   respjson.Field
		MonthlyAdsSpend      respjson.Field
		MonthlyOrganicClicks respjson.Field
		MonthlyOrganicValue  respjson.Field
		MonthlyPaidClicks    respjson.Field
		TotalOrganicKeywords respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanySeo) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanySeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy struct {
	LinkedinIndustry    string   `json:"linkedinIndustry" api:"nullable"`
	LinkedinSpecialties []string `json:"linkedinSpecialties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedinIndustry    respjson.Field
		LinkedinSpecialties respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic struct {
	MomGrowthPercent float64                                                                                   `json:"momGrowthPercent" api:"nullable"`
	MonthlyVisitors  float64                                                                                   `json:"monthlyVisitors" api:"nullable"`
	QoqGrowthPercent float64                                                                                   `json:"qoqGrowthPercent" api:"nullable"`
	TrafficSources   IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources `json:"trafficSources" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MomGrowthPercent respjson.Field
		MonthlyVisitors  respjson.Field
		QoqGrowthPercent respjson.Field
		TrafficSources   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources struct {
	DirectPercent       float64 `json:"directPercent" api:"nullable"`
	PaidReferralPercent float64 `json:"paidReferralPercent" api:"nullable"`
	ReferralPercent     float64 `json:"referralPercent" api:"nullable"`
	SearchPercent       float64 `json:"searchPercent" api:"nullable"`
	SocialPercent       float64 `json:"socialPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirectPercent       respjson.Field
		PaidReferralPercent respjson.Field
		ReferralPercent     respjson.Field
		SearchPercent       respjson.Field
		SocialPercent       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonExperience struct {
	Company     IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompany `json:"company" api:"required"`
	Title       string                                                            `json:"title" api:"required"`
	Description string                                                            `json:"description" api:"nullable"`
	EndDate     string                                                            `json:"endDate" api:"nullable"`
	Location    string                                                            `json:"location" api:"nullable"`
	StartDate   string                                                            `json:"startDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Title       respjson.Field
		Description respjson.Field
		EndDate     respjson.Field
		Location    respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonExperience) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompany struct {
	// Any of "basic".
	Type          string                                                                   `json:"_type" api:"required"`
	Name          string                                                                   `json:"name" api:"required"`
	Address       IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompanyAddress `json:"address" api:"nullable"`
	Description   string                                                                   `json:"description" api:"nullable"`
	Domain        string                                                                   `json:"domain" api:"nullable"`
	EmployeeCount float64                                                                  `json:"employeeCount" api:"nullable"`
	FoundedYear   float64                                                                  `json:"foundedYear" api:"nullable"`
	Industry      string                                                                   `json:"industry" api:"nullable"`
	LinkedinID    string                                                                   `json:"linkedinId" api:"nullable"`
	LinkedinURL   string                                                                   `json:"linkedinUrl" api:"nullable" format:"uri"`
	LogoURL       string                                                                   `json:"logoUrl" api:"nullable" format:"uri"`
	Website       string                                                                   `json:"website" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		Name          respjson.Field
		Address       respjson.Field
		Description   respjson.Field
		Domain        respjson.Field
		EmployeeCount respjson.Field
		FoundedYear   respjson.Field
		Industry      respjson.Field
		LinkedinID    respjson.Field
		LinkedinURL   respjson.Field
		LogoURL       respjson.Field
		Website       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompany) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompanyAddress struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompanyAddress) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonExperienceCompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentPersonLocation struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Raw         string `json:"raw" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Raw         respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentPersonLocation) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataEnrichmentPersonLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentEmailValidation struct {
	Email    string                                                         `json:"email" api:"required"`
	Error    IdentityEnrichedWebhookEventDataEnrichmentEmailValidationError `json:"error" api:"required"`
	Provider string                                                         `json:"provider" api:"required"`
	// Any of "valid", "invalid", "catch_all", "valid_catch_all".
	Validity string `json:"validity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Error       respjson.Field
		Provider    respjson.Field
		Validity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentEmailValidation) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentEmailValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataEnrichmentEmailValidationError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataEnrichmentEmailValidationError) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataEnrichmentEmailValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataQuery struct {
	// Any of "linkedinUrl".
	Match string `json:"match" api:"required"`
	// Any of "person".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Match       respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataQuery) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataQuery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataSession struct {
	CreatedAt    float64                                        `json:"createdAt" api:"required"`
	EndedAt      float64                                        `json:"endedAt" api:"required"`
	Fbclid       string                                         `json:"fbclid" api:"required"`
	Gclid        string                                         `json:"gclid" api:"required"`
	LandingPage  string                                         `json:"landingPage" api:"required"`
	LandingTitle string                                         `json:"landingTitle" api:"required"`
	Network      IdentityEnrichedWebhookEventDataSessionNetwork `json:"network" api:"required"`
	Pid          string                                         `json:"pid" api:"required"`
	Referrer     string                                         `json:"referrer" api:"required"`
	Screen       IdentityEnrichedWebhookEventDataSessionScreen  `json:"screen" api:"required"`
	Sid          string                                         `json:"sid" api:"required"`
	Tenant       string                                         `json:"tenant" api:"required"`
	Timezone     string                                         `json:"timezone" api:"required"`
	UserAgent    string                                         `json:"userAgent" api:"required"`
	Utm          IdentityEnrichedWebhookEventDataSessionUtm     `json:"utm" api:"required"`
	Vid          string                                         `json:"vid" api:"required"`
	Options      map[string]any                                 `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataSession) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataSessionNetwork struct {
	Asn           IdentityEnrichedWebhookEventDataSessionNetworkAsn           `json:"asn" api:"required"`
	BotManagement IdentityEnrichedWebhookEventDataSessionNetworkBotManagement `json:"botManagement" api:"required"`
	City          string                                                      `json:"city" api:"required"`
	Colo          string                                                      `json:"colo" api:"required"`
	Continent     string                                                      `json:"continent" api:"required"`
	Country       string                                                      `json:"country" api:"required"`
	IP            string                                                      `json:"ip" api:"required"`
	IsEu          bool                                                        `json:"isEU" api:"required"`
	Latitude      string                                                      `json:"latitude" api:"required"`
	Longitude     string                                                      `json:"longitude" api:"required"`
	MetroCode     string                                                      `json:"metroCode" api:"required"`
	PostalCode    string                                                      `json:"postalCode" api:"required"`
	Region        string                                                      `json:"region" api:"required"`
	RegionCode    string                                                      `json:"regionCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataSessionNetwork) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataSessionNetworkAsn struct {
	Code float64 `json:"code" api:"required"`
	Name string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataSessionNetworkAsn) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy" api:"required"`
	Score               float64 `json:"score" api:"required"`
	VerifiedBot         bool    `json:"verifiedBot" api:"required"`
	VerifiedBotCategory string  `json:"verifiedBotCategory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataSessionNetworkBotManagement) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityEnrichedWebhookEventDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataSessionScreen struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataSessionScreen) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityEnrichedWebhookEventDataSessionUtm struct {
	Campaign string `json:"campaign" api:"required"`
	Content  string `json:"content" api:"required"`
	Medium   string `json:"medium" api:"required"`
	Source   string `json:"source" api:"required"`
	Term     string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityEnrichedWebhookEventDataSessionUtm) RawJSON() string { return r.JSON.raw }
func (r *IdentityEnrichedWebhookEventDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEvent struct {
	// Unique event identifier
	ID string `json:"id" api:"required"`
	// Unix timestamp (milliseconds) when the event was created
	Created int64                             `json:"created" api:"required"`
	Data    IdentityValidatedWebhookEventData `json:"data" api:"required"`
	// Event type identifier
	Type constant.IdentityValidated `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventData struct {
	Attribution IdentityValidatedWebhookEventDataAttribution  `json:"attribution" api:"required"`
	Session     IdentityValidatedWebhookEventDataSession      `json:"session" api:"required"`
	Validations []IdentityValidatedWebhookEventDataValidation `json:"validations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribution respjson.Field
		Session     respjson.Field
		Validations respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataAttribution struct {
	PixelID string `json:"pixelId" api:"required"`
	// Any of "jNNdd", "OMzN4".
	Prid      string `json:"prid" api:"required"`
	SessionID string `json:"sessionId" api:"required"`
	// Any of "website".
	Type string `json:"type" api:"required"`
	Vid  string `json:"vid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PixelID     respjson.Field
		Prid        respjson.Field
		SessionID   respjson.Field
		Type        respjson.Field
		Vid         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataAttribution) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataAttribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataSession struct {
	CreatedAt    float64                                         `json:"createdAt" api:"required"`
	EndedAt      float64                                         `json:"endedAt" api:"required"`
	Fbclid       string                                          `json:"fbclid" api:"required"`
	Gclid        string                                          `json:"gclid" api:"required"`
	LandingPage  string                                          `json:"landingPage" api:"required"`
	LandingTitle string                                          `json:"landingTitle" api:"required"`
	Network      IdentityValidatedWebhookEventDataSessionNetwork `json:"network" api:"required"`
	Pid          string                                          `json:"pid" api:"required"`
	Referrer     string                                          `json:"referrer" api:"required"`
	Screen       IdentityValidatedWebhookEventDataSessionScreen  `json:"screen" api:"required"`
	Sid          string                                          `json:"sid" api:"required"`
	Tenant       string                                          `json:"tenant" api:"required"`
	Timezone     string                                          `json:"timezone" api:"required"`
	UserAgent    string                                          `json:"userAgent" api:"required"`
	Utm          IdentityValidatedWebhookEventDataSessionUtm     `json:"utm" api:"required"`
	Vid          string                                          `json:"vid" api:"required"`
	Options      map[string]any                                  `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataSession) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataSessionNetwork struct {
	Asn           IdentityValidatedWebhookEventDataSessionNetworkAsn           `json:"asn" api:"required"`
	BotManagement IdentityValidatedWebhookEventDataSessionNetworkBotManagement `json:"botManagement" api:"required"`
	City          string                                                       `json:"city" api:"required"`
	Colo          string                                                       `json:"colo" api:"required"`
	Continent     string                                                       `json:"continent" api:"required"`
	Country       string                                                       `json:"country" api:"required"`
	IP            string                                                       `json:"ip" api:"required"`
	IsEu          bool                                                         `json:"isEU" api:"required"`
	Latitude      string                                                       `json:"latitude" api:"required"`
	Longitude     string                                                       `json:"longitude" api:"required"`
	MetroCode     string                                                       `json:"metroCode" api:"required"`
	PostalCode    string                                                       `json:"postalCode" api:"required"`
	Region        string                                                       `json:"region" api:"required"`
	RegionCode    string                                                       `json:"regionCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataSessionNetwork) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataSessionNetworkAsn struct {
	Code float64 `json:"code" api:"required"`
	Name string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataSessionNetworkAsn) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy" api:"required"`
	Score               float64 `json:"score" api:"required"`
	VerifiedBot         bool    `json:"verifiedBot" api:"required"`
	VerifiedBotCategory string  `json:"verifiedBotCategory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataSessionNetworkBotManagement) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentityValidatedWebhookEventDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataSessionScreen struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataSessionScreen) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataSessionUtm struct {
	Campaign string `json:"campaign" api:"required"`
	Content  string `json:"content" api:"required"`
	Medium   string `json:"medium" api:"required"`
	Source   string `json:"source" api:"required"`
	Term     string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataSessionUtm) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataValidation struct {
	Email    string                                           `json:"email" api:"required"`
	Error    IdentityValidatedWebhookEventDataValidationError `json:"error" api:"required"`
	Provider string                                           `json:"provider" api:"required"`
	// Any of "valid", "invalid", "catch_all", "valid_catch_all".
	Validity string `json:"validity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Error       respjson.Field
		Provider    respjson.Field
		Validity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataValidation) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityValidatedWebhookEventDataValidationError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityValidatedWebhookEventDataValidationError) RawJSON() string { return r.JSON.raw }
func (r *IdentityValidatedWebhookEventDataValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEvent struct {
	// Unique event identifier
	ID string `json:"id" api:"required"`
	// Unix timestamp (milliseconds) when the event was created
	Created int64                                    `json:"created" api:"required"`
	Data    IdentitySessionFinalizedWebhookEventData `json:"data" api:"required"`
	// Event type identifier
	Type constant.IdentitySessionFinalized `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventData struct {
	Enrichment     IdentitySessionFinalizedWebhookEventDataEnrichment `json:"enrichment" api:"required"`
	EventCount     float64                                            `json:"eventCount" api:"required"`
	FinalizedAt    float64                                            `json:"finalizedAt" api:"required"`
	Session        IdentitySessionFinalizedWebhookEventDataSession    `json:"session" api:"required"`
	SessionEndedAt float64                                            `json:"sessionEndedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enrichment     respjson.Field
		EventCount     respjson.Field
		FinalizedAt    respjson.Field
		Session        respjson.Field
		SessionEndedAt respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichment struct {
	Companies        []IdentitySessionFinalizedWebhookEventDataEnrichmentCompany         `json:"companies" api:"required"`
	Person           IdentitySessionFinalizedWebhookEventDataEnrichmentPerson            `json:"person" api:"required"`
	EmailValidations []IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation `json:"emailValidations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Companies        respjson.Field
		Person           respjson.Field
		EmailValidations respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichment) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataEnrichment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompany struct {
	// Any of "full".
	Type       string `json:"_type" api:"required"`
	EnrichedAt string `json:"enrichedAt" api:"required"`
	Name       string `json:"name" api:"required"`
	// Any of "xK9mP", "qR7nL".
	Provider          string                                                                     `json:"provider" api:"required"`
	Address           IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyAddress           `json:"address" api:"nullable"`
	Description       string                                                                     `json:"description" api:"nullable"`
	Domain            string                                                                     `json:"domain" api:"nullable"`
	EmployeeCount     float64                                                                    `json:"employeeCount" api:"nullable"`
	EstimatedRevenue  IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyEstimatedRevenue  `json:"estimatedRevenue" api:"nullable"`
	FoundedYear       float64                                                                    `json:"foundedYear" api:"nullable"`
	Funding           IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyFunding           `json:"funding" api:"nullable"`
	Headcount         IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcount         `json:"headcount" api:"nullable"`
	Industry          string                                                                     `json:"industry" api:"nullable"`
	LinkedinFollowers IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyLinkedinFollowers `json:"linkedinFollowers" api:"nullable"`
	LinkedinID        string                                                                     `json:"linkedinId" api:"nullable"`
	LinkedinURL       string                                                                     `json:"linkedinUrl" api:"nullable" format:"uri"`
	LogoURL           string                                                                     `json:"logoUrl" api:"nullable" format:"uri"`
	Seo               IdentitySessionFinalizedWebhookEventDataEnrichmentCompanySeo               `json:"seo" api:"nullable"`
	Specialties       []string                                                                   `json:"specialties"`
	Taxonomy          IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyTaxonomy          `json:"taxonomy" api:"nullable"`
	Website           string                                                                     `json:"website" api:"nullable" format:"uri"`
	WebTraffic        IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTraffic        `json:"webTraffic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		EnrichedAt        respjson.Field
		Name              respjson.Field
		Provider          respjson.Field
		Address           respjson.Field
		Description       respjson.Field
		Domain            respjson.Field
		EmployeeCount     respjson.Field
		EstimatedRevenue  respjson.Field
		FoundedYear       respjson.Field
		Funding           respjson.Field
		Headcount         respjson.Field
		Industry          respjson.Field
		LinkedinFollowers respjson.Field
		LinkedinID        respjson.Field
		LinkedinURL       respjson.Field
		LogoURL           respjson.Field
		Seo               respjson.Field
		Specialties       respjson.Field
		Taxonomy          respjson.Field
		Website           respjson.Field
		WebTraffic        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompany) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyAddress struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyAddress) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyEstimatedRevenue struct {
	Currency string  `json:"currency"`
	Max      float64 `json:"max" api:"nullable"`
	Min      float64 `json:"min" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyEstimatedRevenue) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyEstimatedRevenue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyFunding struct {
	DaysSinceLastRound float64  `json:"daysSinceLastRound" api:"nullable"`
	Investors          []string `json:"investors"`
	LastRoundAmountUsd float64  `json:"lastRoundAmountUsd" api:"nullable"`
	LastRoundType      string   `json:"lastRoundType" api:"nullable"`
	TotalRaisedUsd     float64  `json:"totalRaisedUsd" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysSinceLastRound respjson.Field
		Investors          respjson.Field
		LastRoundAmountUsd respjson.Field
		LastRoundType      respjson.Field
		TotalRaisedUsd     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyFunding) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyFunding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcount struct {
	Growth IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcountGrowth `json:"growth" api:"nullable"`
	Total  float64                                                                  `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Growth      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcount) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcountGrowth struct {
	Mom float64 `json:"mom" api:"nullable"`
	Qoq float64 `json:"qoq" api:"nullable"`
	Yoy float64 `json:"yoy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mom         respjson.Field
		Qoq         respjson.Field
		Yoy         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcountGrowth) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyHeadcountGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyLinkedinFollowers struct {
	Count            float64 `json:"count" api:"nullable"`
	MomGrowthPercent float64 `json:"momGrowthPercent" api:"nullable"`
	QoqGrowthPercent float64 `json:"qoqGrowthPercent" api:"nullable"`
	YoyGrowthPercent float64 `json:"yoyGrowthPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count            respjson.Field
		MomGrowthPercent respjson.Field
		QoqGrowthPercent respjson.Field
		YoyGrowthPercent respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyLinkedinFollowers) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyLinkedinFollowers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanySeo struct {
	AverageAdRank        float64 `json:"averageAdRank" api:"nullable"`
	AverageOrganicRank   float64 `json:"averageOrganicRank" api:"nullable"`
	MonthlyAdsSpend      float64 `json:"monthlyAdsSpend" api:"nullable"`
	MonthlyOrganicClicks float64 `json:"monthlyOrganicClicks" api:"nullable"`
	MonthlyOrganicValue  float64 `json:"monthlyOrganicValue" api:"nullable"`
	MonthlyPaidClicks    float64 `json:"monthlyPaidClicks" api:"nullable"`
	TotalOrganicKeywords float64 `json:"totalOrganicKeywords" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageAdRank        respjson.Field
		AverageOrganicRank   respjson.Field
		MonthlyAdsSpend      respjson.Field
		MonthlyOrganicClicks respjson.Field
		MonthlyOrganicValue  respjson.Field
		MonthlyPaidClicks    respjson.Field
		TotalOrganicKeywords respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanySeo) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanySeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyTaxonomy struct {
	LinkedinIndustry    string   `json:"linkedinIndustry" api:"nullable"`
	LinkedinSpecialties []string `json:"linkedinSpecialties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedinIndustry    respjson.Field
		LinkedinSpecialties respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyTaxonomy) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTraffic struct {
	MomGrowthPercent float64                                                                           `json:"momGrowthPercent" api:"nullable"`
	MonthlyVisitors  float64                                                                           `json:"monthlyVisitors" api:"nullable"`
	QoqGrowthPercent float64                                                                           `json:"qoqGrowthPercent" api:"nullable"`
	TrafficSources   IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources `json:"trafficSources" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MomGrowthPercent respjson.Field
		MonthlyVisitors  respjson.Field
		QoqGrowthPercent respjson.Field
		TrafficSources   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTraffic) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources struct {
	DirectPercent       float64 `json:"directPercent" api:"nullable"`
	PaidReferralPercent float64 `json:"paidReferralPercent" api:"nullable"`
	ReferralPercent     float64 `json:"referralPercent" api:"nullable"`
	SearchPercent       float64 `json:"searchPercent" api:"nullable"`
	SocialPercent       float64 `json:"socialPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirectPercent       respjson.Field
		PaidReferralPercent respjson.Field
		ReferralPercent     respjson.Field
		SearchPercent       respjson.Field
		SocialPercent       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentCompanyWebTrafficTrafficSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPerson struct {
	EnrichedAt  string `json:"enrichedAt" api:"required"`
	LinkedinURL string `json:"linkedinUrl" api:"required" format:"uri"`
	// Any of "xK9mP", "qR7nL".
	Provider          string                                                               `json:"provider" api:"required"`
	Connections       float64                                                              `json:"connections" api:"nullable"`
	Education         []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation  `json:"education"`
	Email             string                                                               `json:"email" api:"nullable" format:"email"`
	Employments       []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment `json:"employments"`
	Experience        []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience `json:"experience"`
	FirstName         string                                                               `json:"firstName" api:"nullable"`
	FullName          string                                                               `json:"fullName" api:"nullable"`
	Headline          string                                                               `json:"headline" api:"nullable"`
	Languages         []string                                                             `json:"languages"`
	LastName          string                                                               `json:"lastName" api:"nullable"`
	LinkedinID        string                                                               `json:"linkedinId" api:"nullable"`
	Location          IdentitySessionFinalizedWebhookEventDataEnrichmentPersonLocation     `json:"location" api:"nullable"`
	ProfilePictureURL string                                                               `json:"profilePictureUrl" api:"nullable" format:"uri"`
	Skills            []string                                                             `json:"skills"`
	Summary           string                                                               `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnrichedAt        respjson.Field
		LinkedinURL       respjson.Field
		Provider          respjson.Field
		Connections       respjson.Field
		Education         respjson.Field
		Email             respjson.Field
		Employments       respjson.Field
		Experience        respjson.Field
		FirstName         respjson.Field
		FullName          respjson.Field
		Headline          respjson.Field
		Languages         respjson.Field
		LastName          respjson.Field
		LinkedinID        respjson.Field
		Location          respjson.Field
		ProfilePictureURL respjson.Field
		Skills            respjson.Field
		Summary           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPerson) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation struct {
	InstituteName       string `json:"instituteName" api:"required"`
	Degree              string `json:"degree" api:"nullable"`
	EndDate             string `json:"endDate" api:"nullable"`
	FieldOfStudy        string `json:"fieldOfStudy" api:"nullable"`
	InstituteLinkedinID string `json:"instituteLinkedinId" api:"nullable"`
	InstituteLogoURL    string `json:"instituteLogoUrl" api:"nullable" format:"uri"`
	StartDate           string `json:"startDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstituteName       respjson.Field
		Degree              respjson.Field
		EndDate             respjson.Field
		FieldOfStudy        respjson.Field
		InstituteLinkedinID respjson.Field
		InstituteLogoURL    respjson.Field
		StartDate           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment struct {
	Company     IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompany `json:"company" api:"required"`
	Title       string                                                                    `json:"title" api:"required"`
	Description string                                                                    `json:"description" api:"nullable"`
	EndDate     string                                                                    `json:"endDate" api:"nullable"`
	Location    string                                                                    `json:"location" api:"nullable"`
	StartDate   string                                                                    `json:"startDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Title       respjson.Field
		Description respjson.Field
		EndDate     respjson.Field
		Location    respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompany struct {
	// Any of "full".
	Type       string `json:"_type" api:"required"`
	EnrichedAt string `json:"enrichedAt" api:"required"`
	Name       string `json:"name" api:"required"`
	// Any of "xK9mP", "qR7nL".
	Provider          string                                                                                     `json:"provider" api:"required"`
	Address           IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress           `json:"address" api:"nullable"`
	Description       string                                                                                     `json:"description" api:"nullable"`
	Domain            string                                                                                     `json:"domain" api:"nullable"`
	EmployeeCount     float64                                                                                    `json:"employeeCount" api:"nullable"`
	EstimatedRevenue  IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue  `json:"estimatedRevenue" api:"nullable"`
	FoundedYear       float64                                                                                    `json:"foundedYear" api:"nullable"`
	Funding           IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding           `json:"funding" api:"nullable"`
	Headcount         IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount         `json:"headcount" api:"nullable"`
	Industry          string                                                                                     `json:"industry" api:"nullable"`
	LinkedinFollowers IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers `json:"linkedinFollowers" api:"nullable"`
	LinkedinID        string                                                                                     `json:"linkedinId" api:"nullable"`
	LinkedinURL       string                                                                                     `json:"linkedinUrl" api:"nullable" format:"uri"`
	LogoURL           string                                                                                     `json:"logoUrl" api:"nullable" format:"uri"`
	Seo               IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanySeo               `json:"seo" api:"nullable"`
	Specialties       []string                                                                                   `json:"specialties"`
	Taxonomy          IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy          `json:"taxonomy" api:"nullable"`
	Website           string                                                                                     `json:"website" api:"nullable" format:"uri"`
	WebTraffic        IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic        `json:"webTraffic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		EnrichedAt        respjson.Field
		Name              respjson.Field
		Provider          respjson.Field
		Address           respjson.Field
		Description       respjson.Field
		Domain            respjson.Field
		EmployeeCount     respjson.Field
		EstimatedRevenue  respjson.Field
		FoundedYear       respjson.Field
		Funding           respjson.Field
		Headcount         respjson.Field
		Industry          respjson.Field
		LinkedinFollowers respjson.Field
		LinkedinID        respjson.Field
		LinkedinURL       respjson.Field
		LogoURL           respjson.Field
		Seo               respjson.Field
		Specialties       respjson.Field
		Taxonomy          respjson.Field
		Website           respjson.Field
		WebTraffic        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompany) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue struct {
	Currency string  `json:"currency"`
	Max      float64 `json:"max" api:"nullable"`
	Min      float64 `json:"min" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyEstimatedRevenue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding struct {
	DaysSinceLastRound float64  `json:"daysSinceLastRound" api:"nullable"`
	Investors          []string `json:"investors"`
	LastRoundAmountUsd float64  `json:"lastRoundAmountUsd" api:"nullable"`
	LastRoundType      string   `json:"lastRoundType" api:"nullable"`
	TotalRaisedUsd     float64  `json:"totalRaisedUsd" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DaysSinceLastRound respjson.Field
		Investors          respjson.Field
		LastRoundAmountUsd respjson.Field
		LastRoundType      respjson.Field
		TotalRaisedUsd     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyFunding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount struct {
	Growth IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth `json:"growth" api:"nullable"`
	Total  float64                                                                                  `json:"total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Growth      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth struct {
	Mom float64 `json:"mom" api:"nullable"`
	Qoq float64 `json:"qoq" api:"nullable"`
	Yoy float64 `json:"yoy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mom         respjson.Field
		Qoq         respjson.Field
		Yoy         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyHeadcountGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers struct {
	Count            float64 `json:"count" api:"nullable"`
	MomGrowthPercent float64 `json:"momGrowthPercent" api:"nullable"`
	QoqGrowthPercent float64 `json:"qoqGrowthPercent" api:"nullable"`
	YoyGrowthPercent float64 `json:"yoyGrowthPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count            respjson.Field
		MomGrowthPercent respjson.Field
		QoqGrowthPercent respjson.Field
		YoyGrowthPercent respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyLinkedinFollowers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanySeo struct {
	AverageAdRank        float64 `json:"averageAdRank" api:"nullable"`
	AverageOrganicRank   float64 `json:"averageOrganicRank" api:"nullable"`
	MonthlyAdsSpend      float64 `json:"monthlyAdsSpend" api:"nullable"`
	MonthlyOrganicClicks float64 `json:"monthlyOrganicClicks" api:"nullable"`
	MonthlyOrganicValue  float64 `json:"monthlyOrganicValue" api:"nullable"`
	MonthlyPaidClicks    float64 `json:"monthlyPaidClicks" api:"nullable"`
	TotalOrganicKeywords float64 `json:"totalOrganicKeywords" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageAdRank        respjson.Field
		AverageOrganicRank   respjson.Field
		MonthlyAdsSpend      respjson.Field
		MonthlyOrganicClicks respjson.Field
		MonthlyOrganicValue  respjson.Field
		MonthlyPaidClicks    respjson.Field
		TotalOrganicKeywords respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanySeo) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanySeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy struct {
	LinkedinIndustry    string   `json:"linkedinIndustry" api:"nullable"`
	LinkedinSpecialties []string `json:"linkedinSpecialties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedinIndustry    respjson.Field
		LinkedinSpecialties respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic struct {
	MomGrowthPercent float64                                                                                           `json:"momGrowthPercent" api:"nullable"`
	MonthlyVisitors  float64                                                                                           `json:"monthlyVisitors" api:"nullable"`
	QoqGrowthPercent float64                                                                                           `json:"qoqGrowthPercent" api:"nullable"`
	TrafficSources   IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources `json:"trafficSources" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MomGrowthPercent respjson.Field
		MonthlyVisitors  respjson.Field
		QoqGrowthPercent respjson.Field
		TrafficSources   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources struct {
	DirectPercent       float64 `json:"directPercent" api:"nullable"`
	PaidReferralPercent float64 `json:"paidReferralPercent" api:"nullable"`
	ReferralPercent     float64 `json:"referralPercent" api:"nullable"`
	SearchPercent       float64 `json:"searchPercent" api:"nullable"`
	SocialPercent       float64 `json:"socialPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirectPercent       respjson.Field
		PaidReferralPercent respjson.Field
		ReferralPercent     respjson.Field
		SearchPercent       respjson.Field
		SocialPercent       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmploymentCompanyWebTrafficTrafficSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience struct {
	Company     IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompany `json:"company" api:"required"`
	Title       string                                                                    `json:"title" api:"required"`
	Description string                                                                    `json:"description" api:"nullable"`
	EndDate     string                                                                    `json:"endDate" api:"nullable"`
	Location    string                                                                    `json:"location" api:"nullable"`
	StartDate   string                                                                    `json:"startDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Title       respjson.Field
		Description respjson.Field
		EndDate     respjson.Field
		Location    respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompany struct {
	// Any of "basic".
	Type          string                                                                           `json:"_type" api:"required"`
	Name          string                                                                           `json:"name" api:"required"`
	Address       IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompanyAddress `json:"address" api:"nullable"`
	Description   string                                                                           `json:"description" api:"nullable"`
	Domain        string                                                                           `json:"domain" api:"nullable"`
	EmployeeCount float64                                                                          `json:"employeeCount" api:"nullable"`
	FoundedYear   float64                                                                          `json:"foundedYear" api:"nullable"`
	Industry      string                                                                           `json:"industry" api:"nullable"`
	LinkedinID    string                                                                           `json:"linkedinId" api:"nullable"`
	LinkedinURL   string                                                                           `json:"linkedinUrl" api:"nullable" format:"uri"`
	LogoURL       string                                                                           `json:"logoUrl" api:"nullable" format:"uri"`
	Website       string                                                                           `json:"website" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		Name          respjson.Field
		Address       respjson.Field
		Description   respjson.Field
		Domain        respjson.Field
		EmployeeCount respjson.Field
		FoundedYear   respjson.Field
		Industry      respjson.Field
		LinkedinID    respjson.Field
		LinkedinURL   respjson.Field
		LogoURL       respjson.Field
		Website       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompany) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompanyAddress struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompanyAddress) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceCompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentPersonLocation struct {
	Country     string `json:"country" api:"nullable"`
	CountryCode string `json:"countryCode" api:"nullable"`
	Locality    string `json:"locality" api:"nullable"`
	PostalCode  string `json:"postalCode" api:"nullable"`
	Raw         string `json:"raw" api:"nullable"`
	Region      string `json:"region" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Raw         respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentPersonLocation) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentPersonLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation struct {
	Email    string                                                                 `json:"email" api:"required"`
	Error    IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidationError `json:"error" api:"required"`
	Provider string                                                                 `json:"provider" api:"required"`
	// Any of "valid", "invalid", "catch_all", "valid_catch_all".
	Validity string `json:"validity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Error       respjson.Field
		Provider    respjson.Field
		Validity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidationError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidationError) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSession struct {
	CreatedAt    float64                                                     `json:"createdAt" api:"required"`
	EndedAt      float64                                                     `json:"endedAt" api:"required"`
	Events       []IdentitySessionFinalizedWebhookEventDataSessionEvent      `json:"events" api:"required"`
	Fbclid       string                                                      `json:"fbclid" api:"required"`
	Gclid        string                                                      `json:"gclid" api:"required"`
	LandingPage  string                                                      `json:"landingPage" api:"required"`
	LandingTitle string                                                      `json:"landingTitle" api:"required"`
	LastActivity float64                                                     `json:"lastActivity" api:"required"`
	Network      IdentitySessionFinalizedWebhookEventDataSessionNetwork      `json:"network" api:"required"`
	Pid          string                                                      `json:"pid" api:"required"`
	Referrer     string                                                      `json:"referrer" api:"required"`
	Screen       IdentitySessionFinalizedWebhookEventDataSessionScreen       `json:"screen" api:"required"`
	Sid          string                                                      `json:"sid" api:"required"`
	Tenant       string                                                      `json:"tenant" api:"required"`
	Timezone     string                                                      `json:"timezone" api:"required"`
	UserAgent    string                                                      `json:"userAgent" api:"required"`
	Utm          IdentitySessionFinalizedWebhookEventDataSessionUtm          `json:"utm" api:"required"`
	Vid          string                                                      `json:"vid" api:"required"`
	Enrichments  []IdentitySessionFinalizedWebhookEventDataSessionEnrichment `json:"enrichments"`
	Options      map[string]any                                              `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Events       respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		LastActivity respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Enrichments  respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSession) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionEvent struct {
	Outlink string         `json:"outlink" api:"required"`
	Props   map[string]any `json:"props" api:"required"`
	Ref     string         `json:"ref" api:"required"`
	Signal  string         `json:"signal" api:"required"`
	Title   string         `json:"title" api:"required"`
	Ts      float64        `json:"ts" api:"required"`
	URL     string         `json:"url" api:"required"`
	V       string         `json:"v" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outlink     respjson.Field
		Props       respjson.Field
		Ref         respjson.Field
		Signal      respjson.Field
		Title       respjson.Field
		Ts          respjson.Field
		URL         respjson.Field
		V           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionEvent) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataSessionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionNetwork struct {
	Asn           IdentitySessionFinalizedWebhookEventDataSessionNetworkAsn           `json:"asn" api:"required"`
	BotManagement IdentitySessionFinalizedWebhookEventDataSessionNetworkBotManagement `json:"botManagement" api:"required"`
	City          string                                                              `json:"city" api:"required"`
	Colo          string                                                              `json:"colo" api:"required"`
	Continent     string                                                              `json:"continent" api:"required"`
	Country       string                                                              `json:"country" api:"required"`
	IP            string                                                              `json:"ip" api:"required"`
	IsEu          bool                                                                `json:"isEU" api:"required"`
	Latitude      string                                                              `json:"latitude" api:"required"`
	Longitude     string                                                              `json:"longitude" api:"required"`
	MetroCode     string                                                              `json:"metroCode" api:"required"`
	PostalCode    string                                                              `json:"postalCode" api:"required"`
	Region        string                                                              `json:"region" api:"required"`
	RegionCode    string                                                              `json:"regionCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionNetwork) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionNetworkAsn struct {
	Code float64 `json:"code" api:"required"`
	Name string  `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionNetworkAsn) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy" api:"required"`
	Score               float64 `json:"score" api:"required"`
	VerifiedBot         bool    `json:"verifiedBot" api:"required"`
	VerifiedBotCategory string  `json:"verifiedBotCategory" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionNetworkBotManagement) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionScreen struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionScreen) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionUtm struct {
	Campaign string `json:"campaign" api:"required"`
	Content  string `json:"content" api:"required"`
	Medium   string `json:"medium" api:"required"`
	Source   string `json:"source" api:"required"`
	Term     string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionUtm) RawJSON() string { return r.JSON.raw }
func (r *IdentitySessionFinalizedWebhookEventDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentitySessionFinalizedWebhookEventDataSessionEnrichment struct {
	EnrichedAt     float64 `json:"enrichedAt" api:"required"`
	EnrichmentData string  `json:"enrichmentData" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnrichedAt     respjson.Field
		EnrichmentData respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentitySessionFinalizedWebhookEventDataSessionEnrichment) RawJSON() string {
	return r.JSON.raw
}
func (r *IdentitySessionFinalizedWebhookEventDataSessionEnrichment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnion contains all possible properties and values from
// [IdentityResolvedWebhookEvent], [IdentityQualifiedWebhookEvent],
// [IdentityEnrichedWebhookEvent], [IdentityValidatedWebhookEvent],
// [IdentitySessionFinalizedWebhookEvent].
//
// Use the [UnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	// This field is a union of [IdentityResolvedWebhookEventData],
	// [IdentityQualifiedWebhookEventData], [IdentityEnrichedWebhookEventData],
	// [IdentityValidatedWebhookEventData], [IdentitySessionFinalizedWebhookEventData]
	Data UnwrapWebhookEventUnionData `json:"data"`
	// Any of "identity.resolved", "identity.qualified", "identity.enriched",
	// "identity.validated", "identity.session.finalized".
	Type string `json:"type"`
	JSON struct {
		ID      respjson.Field
		Created respjson.Field
		Data    respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

// anyUnwrapWebhookEvent is implemented by each variant of
// [UnwrapWebhookEventUnion] to add type safety for the return type of
// [UnwrapWebhookEventUnion.AsAny]
type anyUnwrapWebhookEvent interface {
	implUnwrapWebhookEventUnion()
}

func (IdentityResolvedWebhookEvent) implUnwrapWebhookEventUnion()         {}
func (IdentityQualifiedWebhookEvent) implUnwrapWebhookEventUnion()        {}
func (IdentityEnrichedWebhookEvent) implUnwrapWebhookEventUnion()         {}
func (IdentityValidatedWebhookEvent) implUnwrapWebhookEventUnion()        {}
func (IdentitySessionFinalizedWebhookEvent) implUnwrapWebhookEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UnwrapWebhookEventUnion.AsAny().(type) {
//	case midboundcloud.IdentityResolvedWebhookEvent:
//	case midboundcloud.IdentityQualifiedWebhookEvent:
//	case midboundcloud.IdentityEnrichedWebhookEvent:
//	case midboundcloud.IdentityValidatedWebhookEvent:
//	case midboundcloud.IdentitySessionFinalizedWebhookEvent:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UnwrapWebhookEventUnion) AsAny() anyUnwrapWebhookEvent {
	switch u.Type {
	case "identity.resolved":
		return u.AsIdentityResolved()
	case "identity.qualified":
		return u.AsIdentityQualified()
	case "identity.enriched":
		return u.AsIdentityEnriched()
	case "identity.validated":
		return u.AsIdentityValidated()
	case "identity.session.finalized":
		return u.AsIdentitySessionFinalized()
	}
	return nil
}

func (u UnwrapWebhookEventUnion) AsIdentityResolved() (v IdentityResolvedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsIdentityQualified() (v IdentityQualifiedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsIdentityEnriched() (v IdentityEnrichedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsIdentityValidated() (v IdentityValidatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsIdentitySessionFinalized() (v IdentitySessionFinalizedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	// This field is a union of [IdentityResolvedWebhookEventDataAttribution],
	// [IdentityQualifiedWebhookEventDataAttribution],
	// [IdentityEnrichedWebhookEventDataAttribution],
	// [IdentityValidatedWebhookEventDataAttribution]
	Attribution UnwrapWebhookEventUnionDataAttribution `json:"attribution"`
	// This field is a union of [IdentityResolvedWebhookEventDataIdentity],
	// [IdentityQualifiedWebhookEventDataIdentity]
	Identity UnwrapWebhookEventUnionDataIdentity `json:"identity"`
	// This field is a union of [IdentityResolvedWebhookEventDataSession],
	// [IdentityQualifiedWebhookEventDataSession],
	// [IdentityEnrichedWebhookEventDataSession],
	// [IdentityValidatedWebhookEventDataSession],
	// [IdentitySessionFinalizedWebhookEventDataSession]
	Session UnwrapWebhookEventUnionDataSession `json:"session"`
	// This field is from variant [IdentityEnrichedWebhookEventData].
	CompaniesEnriched float64 `json:"companiesEnriched"`
	// This field is a union of [IdentityEnrichedWebhookEventDataEnrichment],
	// [IdentitySessionFinalizedWebhookEventDataEnrichment]
	Enrichment UnwrapWebhookEventUnionDataEnrichment `json:"enrichment"`
	// This field is from variant [IdentityEnrichedWebhookEventData].
	Query IdentityEnrichedWebhookEventDataQuery `json:"query"`
	// This field is from variant [IdentityValidatedWebhookEventData].
	Validations []IdentityValidatedWebhookEventDataValidation `json:"validations"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventData].
	EventCount float64 `json:"eventCount"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventData].
	FinalizedAt float64 `json:"finalizedAt"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventData].
	SessionEndedAt float64 `json:"sessionEndedAt"`
	JSON           struct {
		Attribution       respjson.Field
		Identity          respjson.Field
		Session           respjson.Field
		CompaniesEnriched respjson.Field
		Enrichment        respjson.Field
		Query             respjson.Field
		Validations       respjson.Field
		EventCount        respjson.Field
		FinalizedAt       respjson.Field
		SessionEndedAt    respjson.Field
		raw               string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataAttribution is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataAttribution provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataAttribution struct {
	PixelID   string `json:"pixelId"`
	Prid      string `json:"prid"`
	SessionID string `json:"sessionId"`
	Type      string `json:"type"`
	Vid       string `json:"vid"`
	JSON      struct {
		PixelID   respjson.Field
		Prid      respjson.Field
		SessionID respjson.Field
		Type      respjson.Field
		Vid       respjson.Field
		raw       string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataAttribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataIdentity is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataIdentity provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataIdentity struct {
	// This field is a union of [IdentityResolvedWebhookEventDataIdentityDemographics],
	// [IdentityQualifiedWebhookEventDataIdentityDemographics]
	Demographics UnwrapWebhookEventUnionDataIdentityDemographics `json:"demographics"`
	// This field is a union of [[]IdentityResolvedWebhookEventDataIdentityEmail],
	// [[]IdentityQualifiedWebhookEventDataIdentityEmail]
	Emails      UnwrapWebhookEventUnionDataIdentityEmails `json:"emails"`
	LinkedinURL string                                    `json:"linkedinUrl"`
	// This field is a union of [IdentityResolvedWebhookEventDataIdentityLocation],
	// [IdentityQualifiedWebhookEventDataIdentityLocation]
	Location UnwrapWebhookEventUnionDataIdentityLocation `json:"location"`
	Phones   []string                                    `json:"phones"`
	// This field is a union of [IdentityResolvedWebhookEventDataIdentityProfessional],
	// [IdentityQualifiedWebhookEventDataIdentityProfessional]
	Professional UnwrapWebhookEventUnionDataIdentityProfessional `json:"professional"`
	JSON         struct {
		Demographics respjson.Field
		Emails       respjson.Field
		LinkedinURL  respjson.Field
		Location     respjson.Field
		Phones       respjson.Field
		Professional respjson.Field
		raw          string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataIdentity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataIdentityDemographics is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataIdentityDemographics
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataIdentityDemographics struct {
	FirstName   string `json:"firstName"`
	HasChildren bool   `json:"hasChildren"`
	IsHomeowner bool   `json:"isHomeowner"`
	IsMarried   bool   `json:"isMarried"`
	LastName    string `json:"lastName"`
	JSON        struct {
		FirstName   respjson.Field
		HasChildren respjson.Field
		IsHomeowner respjson.Field
		IsMarried   respjson.Field
		LastName    respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataIdentityDemographics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataIdentityEmails is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataIdentityEmails provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityResolvedWebhookEventDataIdentityEmails
// OfIdentityQualifiedWebhookEventDataIdentityEmails]
type UnwrapWebhookEventUnionDataIdentityEmails struct {
	// This field will be present if the value is a
	// [[]IdentityResolvedWebhookEventDataIdentityEmail] instead of an object.
	OfIdentityResolvedWebhookEventDataIdentityEmails []IdentityResolvedWebhookEventDataIdentityEmail `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentityQualifiedWebhookEventDataIdentityEmail] instead of an object.
	OfIdentityQualifiedWebhookEventDataIdentityEmails []IdentityQualifiedWebhookEventDataIdentityEmail `json:",inline"`
	JSON                                              struct {
		OfIdentityResolvedWebhookEventDataIdentityEmails  respjson.Field
		OfIdentityQualifiedWebhookEventDataIdentityEmails respjson.Field
		raw                                               string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataIdentityEmails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataIdentityLocation is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataIdentityLocation provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataIdentityLocation struct {
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	JSON    struct {
		City    respjson.Field
		Country respjson.Field
		State   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataIdentityLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataIdentityProfessional is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataIdentityProfessional
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataIdentityProfessional struct {
	CompanyName string `json:"companyName"`
	JobTitle    string `json:"jobTitle"`
	JSON        struct {
		CompanyName respjson.Field
		JobTitle    respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataIdentityProfessional) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSession is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataSession provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataSession struct {
	CreatedAt    float64 `json:"createdAt"`
	EndedAt      float64 `json:"endedAt"`
	Fbclid       string  `json:"fbclid"`
	Gclid        string  `json:"gclid"`
	LandingPage  string  `json:"landingPage"`
	LandingTitle string  `json:"landingTitle"`
	// This field is a union of [IdentityResolvedWebhookEventDataSessionNetwork],
	// [IdentityQualifiedWebhookEventDataSessionNetwork],
	// [IdentityEnrichedWebhookEventDataSessionNetwork],
	// [IdentityValidatedWebhookEventDataSessionNetwork],
	// [IdentitySessionFinalizedWebhookEventDataSessionNetwork]
	Network  UnwrapWebhookEventUnionDataSessionNetwork `json:"network"`
	Pid      string                                    `json:"pid"`
	Referrer string                                    `json:"referrer"`
	// This field is a union of [IdentityResolvedWebhookEventDataSessionScreen],
	// [IdentityQualifiedWebhookEventDataSessionScreen],
	// [IdentityEnrichedWebhookEventDataSessionScreen],
	// [IdentityValidatedWebhookEventDataSessionScreen],
	// [IdentitySessionFinalizedWebhookEventDataSessionScreen]
	Screen    UnwrapWebhookEventUnionDataSessionScreen `json:"screen"`
	Sid       string                                   `json:"sid"`
	Tenant    string                                   `json:"tenant"`
	Timezone  string                                   `json:"timezone"`
	UserAgent string                                   `json:"userAgent"`
	// This field is a union of [IdentityResolvedWebhookEventDataSessionUtm],
	// [IdentityQualifiedWebhookEventDataSessionUtm],
	// [IdentityEnrichedWebhookEventDataSessionUtm],
	// [IdentityValidatedWebhookEventDataSessionUtm],
	// [IdentitySessionFinalizedWebhookEventDataSessionUtm]
	Utm     UnwrapWebhookEventUnionDataSessionUtm `json:"utm"`
	Vid     string                                `json:"vid"`
	Options any                                   `json:"options"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventDataSession].
	Events []IdentitySessionFinalizedWebhookEventDataSessionEvent `json:"events"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventDataSession].
	LastActivity float64 `json:"lastActivity"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventDataSession].
	Enrichments []IdentitySessionFinalizedWebhookEventDataSessionEnrichment `json:"enrichments"`
	JSON        struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Options      respjson.Field
		Events       respjson.Field
		LastActivity respjson.Field
		Enrichments  respjson.Field
		raw          string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSessionNetwork is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataSessionNetwork provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataSessionNetwork struct {
	// This field is a union of [IdentityResolvedWebhookEventDataSessionNetworkAsn],
	// [IdentityQualifiedWebhookEventDataSessionNetworkAsn],
	// [IdentityEnrichedWebhookEventDataSessionNetworkAsn],
	// [IdentityValidatedWebhookEventDataSessionNetworkAsn],
	// [IdentitySessionFinalizedWebhookEventDataSessionNetworkAsn]
	Asn UnwrapWebhookEventUnionDataSessionNetworkAsn `json:"asn"`
	// This field is a union of
	// [IdentityResolvedWebhookEventDataSessionNetworkBotManagement],
	// [IdentityQualifiedWebhookEventDataSessionNetworkBotManagement],
	// [IdentityEnrichedWebhookEventDataSessionNetworkBotManagement],
	// [IdentityValidatedWebhookEventDataSessionNetworkBotManagement],
	// [IdentitySessionFinalizedWebhookEventDataSessionNetworkBotManagement]
	BotManagement UnwrapWebhookEventUnionDataSessionNetworkBotManagement `json:"botManagement"`
	City          string                                                 `json:"city"`
	Colo          string                                                 `json:"colo"`
	Continent     string                                                 `json:"continent"`
	Country       string                                                 `json:"country"`
	IP            string                                                 `json:"ip"`
	IsEu          bool                                                   `json:"isEU"`
	Latitude      string                                                 `json:"latitude"`
	Longitude     string                                                 `json:"longitude"`
	MetroCode     string                                                 `json:"metroCode"`
	PostalCode    string                                                 `json:"postalCode"`
	Region        string                                                 `json:"region"`
	RegionCode    string                                                 `json:"regionCode"`
	JSON          struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		raw           string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSessionNetworkAsn is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataSessionNetworkAsn provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataSessionNetworkAsn struct {
	Code float64 `json:"code"`
	Name string  `json:"name"`
	JSON struct {
		Code respjson.Field
		Name respjson.Field
		raw  string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSessionNetworkBotManagement is an implicit subunion
// of [UnwrapWebhookEventUnion].
// UnwrapWebhookEventUnionDataSessionNetworkBotManagement provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy"`
	Score               float64 `json:"score"`
	VerifiedBot         bool    `json:"verifiedBot"`
	VerifiedBotCategory string  `json:"verifiedBotCategory"`
	JSON                struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSessionScreen is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataSessionScreen provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataSessionScreen struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	JSON   struct {
		Height respjson.Field
		Width  respjson.Field
		raw    string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataSessionUtm is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataSessionUtm provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataSessionUtm struct {
	Campaign string `json:"campaign"`
	Content  string `json:"content"`
	Medium   string `json:"medium"`
	Source   string `json:"source"`
	Term     string `json:"term"`
	JSON     struct {
		Campaign respjson.Field
		Content  respjson.Field
		Medium   respjson.Field
		Source   respjson.Field
		Term     respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichment is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichment provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataEnrichment struct {
	// This field is a union of [[]IdentityEnrichedWebhookEventDataEnrichmentCompany],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentCompany]
	Companies UnwrapWebhookEventUnionDataEnrichmentCompanies `json:"companies"`
	// This field is a union of [IdentityEnrichedWebhookEventDataEnrichmentPerson],
	// [IdentitySessionFinalizedWebhookEventDataEnrichmentPerson]
	Person UnwrapWebhookEventUnionDataEnrichmentPerson `json:"person"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentEmailValidation],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation]
	EmailValidations UnwrapWebhookEventUnionDataEnrichmentEmailValidations `json:"emailValidations"`
	JSON             struct {
		Companies        respjson.Field
		Person           respjson.Field
		EmailValidations respjson.Field
		raw              string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentCompanies is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichmentCompanies
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentCompanies
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentCompanies]
type UnwrapWebhookEventUnionDataEnrichmentCompanies struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentCompany] instead of an object.
	OfIdentityEnrichedWebhookEventDataEnrichmentCompanies []IdentityEnrichedWebhookEventDataEnrichmentCompany `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentCompany] instead of an
	// object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentCompanies []IdentitySessionFinalizedWebhookEventDataEnrichmentCompany `json:",inline"`
	JSON                                                          struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentCompanies         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentCompanies respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentCompanies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentPerson is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichmentPerson provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataEnrichmentPerson struct {
	EnrichedAt  string  `json:"enrichedAt"`
	LinkedinURL string  `json:"linkedinUrl"`
	Provider    string  `json:"provider"`
	Connections float64 `json:"connections"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEducation],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation]
	Education UnwrapWebhookEventUnionDataEnrichmentPersonEducation `json:"education"`
	Email     string                                               `json:"email"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment]
	Employments UnwrapWebhookEventUnionDataEnrichmentPersonEmployments `json:"employments"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonExperience],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience]
	Experience UnwrapWebhookEventUnionDataEnrichmentPersonExperience `json:"experience"`
	FirstName  string                                                `json:"firstName"`
	FullName   string                                                `json:"fullName"`
	Headline   string                                                `json:"headline"`
	Languages  []string                                              `json:"languages"`
	LastName   string                                                `json:"lastName"`
	LinkedinID string                                                `json:"linkedinId"`
	// This field is a union of
	// [IdentityEnrichedWebhookEventDataEnrichmentPersonLocation],
	// [IdentitySessionFinalizedWebhookEventDataEnrichmentPersonLocation]
	Location          UnwrapWebhookEventUnionDataEnrichmentPersonLocation `json:"location"`
	ProfilePictureURL string                                              `json:"profilePictureUrl"`
	Skills            []string                                            `json:"skills"`
	Summary           string                                              `json:"summary"`
	JSON              struct {
		EnrichedAt        respjson.Field
		LinkedinURL       respjson.Field
		Provider          respjson.Field
		Connections       respjson.Field
		Education         respjson.Field
		Email             respjson.Field
		Employments       respjson.Field
		Experience        respjson.Field
		FirstName         respjson.Field
		FullName          respjson.Field
		Headline          respjson.Field
		Languages         respjson.Field
		LastName          respjson.Field
		LinkedinID        respjson.Field
		Location          respjson.Field
		ProfilePictureURL respjson.Field
		Skills            respjson.Field
		Summary           respjson.Field
		raw               string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentPersonEducation is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichmentPersonEducation
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentPersonEducationArray
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducationArray]
type UnwrapWebhookEventUnionDataEnrichmentPersonEducation struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEducation] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentPersonEducationArray []IdentityEnrichedWebhookEventDataEnrichmentPersonEducation `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation] instead of
	// an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducationArray []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation `json:",inline"`
	JSON                                                                     struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentPersonEducationArray         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducationArray respjson.Field
		raw                                                                      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentPersonEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentPersonEmployments is an implicit subunion
// of [UnwrapWebhookEventUnion].
// UnwrapWebhookEventUnionDataEnrichmentPersonEmployments provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentPersonEmployments
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployments]
type UnwrapWebhookEventUnionDataEnrichmentPersonEmployments struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentPersonEmployments []IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment] instead
	// of an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployments []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment `json:",inline"`
	JSON                                                                  struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentPersonEmployments         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployments respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentPersonEmployments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentPersonExperience is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichmentPersonExperience
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentPersonExperienceArray
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceArray]
type UnwrapWebhookEventUnionDataEnrichmentPersonExperience struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonExperience] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentPersonExperienceArray []IdentityEnrichedWebhookEventDataEnrichmentPersonExperience `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience] instead
	// of an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceArray []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience `json:",inline"`
	JSON                                                                      struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentPersonExperienceArray         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceArray respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentPersonExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentPersonLocation is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichmentPersonLocation
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataEnrichmentPersonLocation struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Locality    string `json:"locality"`
	PostalCode  string `json:"postalCode"`
	Raw         string `json:"raw"`
	Region      string `json:"region"`
	JSON        struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Raw         respjson.Field
		Region      respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentPersonLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataEnrichmentEmailValidations is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataEnrichmentEmailValidations
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentEmailValidations
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidations]
type UnwrapWebhookEventUnionDataEnrichmentEmailValidations struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentEmailValidation] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentEmailValidations []IdentityEnrichedWebhookEventDataEnrichmentEmailValidation `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation] instead of
	// an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidations []IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation `json:",inline"`
	JSON                                                                 struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentEmailValidations         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidations respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataEnrichmentEmailValidations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnion contains all possible properties and values from
// [IdentityResolvedWebhookEvent], [IdentityQualifiedWebhookEvent],
// [IdentityEnrichedWebhookEvent], [IdentityValidatedWebhookEvent],
// [IdentitySessionFinalizedWebhookEvent].
//
// Use the [UnwrapUnsafeWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapUnsafeWebhookEventUnion struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	// This field is a union of [IdentityResolvedWebhookEventData],
	// [IdentityQualifiedWebhookEventData], [IdentityEnrichedWebhookEventData],
	// [IdentityValidatedWebhookEventData], [IdentitySessionFinalizedWebhookEventData]
	Data UnwrapUnsafeWebhookEventUnionData `json:"data"`
	// Any of "identity.resolved", "identity.qualified", "identity.enriched",
	// "identity.validated", "identity.session.finalized".
	Type string `json:"type"`
	JSON struct {
		ID      respjson.Field
		Created respjson.Field
		Data    respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

// anyUnwrapUnsafeWebhookEvent is implemented by each variant of
// [UnwrapUnsafeWebhookEventUnion] to add type safety for the return type of
// [UnwrapUnsafeWebhookEventUnion.AsAny]
type anyUnwrapUnsafeWebhookEvent interface {
	implUnwrapUnsafeWebhookEventUnion()
}

func (IdentityResolvedWebhookEvent) implUnwrapUnsafeWebhookEventUnion()         {}
func (IdentityQualifiedWebhookEvent) implUnwrapUnsafeWebhookEventUnion()        {}
func (IdentityEnrichedWebhookEvent) implUnwrapUnsafeWebhookEventUnion()         {}
func (IdentityValidatedWebhookEvent) implUnwrapUnsafeWebhookEventUnion()        {}
func (IdentitySessionFinalizedWebhookEvent) implUnwrapUnsafeWebhookEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UnwrapUnsafeWebhookEventUnion.AsAny().(type) {
//	case midboundcloud.IdentityResolvedWebhookEvent:
//	case midboundcloud.IdentityQualifiedWebhookEvent:
//	case midboundcloud.IdentityEnrichedWebhookEvent:
//	case midboundcloud.IdentityValidatedWebhookEvent:
//	case midboundcloud.IdentitySessionFinalizedWebhookEvent:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UnwrapUnsafeWebhookEventUnion) AsAny() anyUnwrapUnsafeWebhookEvent {
	switch u.Type {
	case "identity.resolved":
		return u.AsIdentityResolved()
	case "identity.qualified":
		return u.AsIdentityQualified()
	case "identity.enriched":
		return u.AsIdentityEnriched()
	case "identity.validated":
		return u.AsIdentityValidated()
	case "identity.session.finalized":
		return u.AsIdentitySessionFinalized()
	}
	return nil
}

func (u UnwrapUnsafeWebhookEventUnion) AsIdentityResolved() (v IdentityResolvedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapUnsafeWebhookEventUnion) AsIdentityQualified() (v IdentityQualifiedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapUnsafeWebhookEventUnion) AsIdentityEnriched() (v IdentityEnrichedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapUnsafeWebhookEventUnion) AsIdentityValidated() (v IdentityValidatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapUnsafeWebhookEventUnion) AsIdentitySessionFinalized() (v IdentitySessionFinalizedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapUnsafeWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapUnsafeWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionData is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionData struct {
	// This field is a union of [IdentityResolvedWebhookEventDataAttribution],
	// [IdentityQualifiedWebhookEventDataAttribution],
	// [IdentityEnrichedWebhookEventDataAttribution],
	// [IdentityValidatedWebhookEventDataAttribution]
	Attribution UnwrapUnsafeWebhookEventUnionDataAttribution `json:"attribution"`
	// This field is a union of [IdentityResolvedWebhookEventDataIdentity],
	// [IdentityQualifiedWebhookEventDataIdentity]
	Identity UnwrapUnsafeWebhookEventUnionDataIdentity `json:"identity"`
	// This field is a union of [IdentityResolvedWebhookEventDataSession],
	// [IdentityQualifiedWebhookEventDataSession],
	// [IdentityEnrichedWebhookEventDataSession],
	// [IdentityValidatedWebhookEventDataSession],
	// [IdentitySessionFinalizedWebhookEventDataSession]
	Session UnwrapUnsafeWebhookEventUnionDataSession `json:"session"`
	// This field is from variant [IdentityEnrichedWebhookEventData].
	CompaniesEnriched float64 `json:"companiesEnriched"`
	// This field is a union of [IdentityEnrichedWebhookEventDataEnrichment],
	// [IdentitySessionFinalizedWebhookEventDataEnrichment]
	Enrichment UnwrapUnsafeWebhookEventUnionDataEnrichment `json:"enrichment"`
	// This field is from variant [IdentityEnrichedWebhookEventData].
	Query IdentityEnrichedWebhookEventDataQuery `json:"query"`
	// This field is from variant [IdentityValidatedWebhookEventData].
	Validations []IdentityValidatedWebhookEventDataValidation `json:"validations"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventData].
	EventCount float64 `json:"eventCount"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventData].
	FinalizedAt float64 `json:"finalizedAt"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventData].
	SessionEndedAt float64 `json:"sessionEndedAt"`
	JSON           struct {
		Attribution       respjson.Field
		Identity          respjson.Field
		Session           respjson.Field
		CompaniesEnriched respjson.Field
		Enrichment        respjson.Field
		Query             respjson.Field
		Validations       respjson.Field
		EventCount        respjson.Field
		FinalizedAt       respjson.Field
		SessionEndedAt    respjson.Field
		raw               string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataAttribution is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataAttribution
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataAttribution struct {
	PixelID   string `json:"pixelId"`
	Prid      string `json:"prid"`
	SessionID string `json:"sessionId"`
	Type      string `json:"type"`
	Vid       string `json:"vid"`
	JSON      struct {
		PixelID   respjson.Field
		Prid      respjson.Field
		SessionID respjson.Field
		Type      respjson.Field
		Vid       respjson.Field
		raw       string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataAttribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataIdentity is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataIdentity
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataIdentity struct {
	// This field is a union of [IdentityResolvedWebhookEventDataIdentityDemographics],
	// [IdentityQualifiedWebhookEventDataIdentityDemographics]
	Demographics UnwrapUnsafeWebhookEventUnionDataIdentityDemographics `json:"demographics"`
	// This field is a union of [[]IdentityResolvedWebhookEventDataIdentityEmail],
	// [[]IdentityQualifiedWebhookEventDataIdentityEmail]
	Emails      UnwrapUnsafeWebhookEventUnionDataIdentityEmails `json:"emails"`
	LinkedinURL string                                          `json:"linkedinUrl"`
	// This field is a union of [IdentityResolvedWebhookEventDataIdentityLocation],
	// [IdentityQualifiedWebhookEventDataIdentityLocation]
	Location UnwrapUnsafeWebhookEventUnionDataIdentityLocation `json:"location"`
	Phones   []string                                          `json:"phones"`
	// This field is a union of [IdentityResolvedWebhookEventDataIdentityProfessional],
	// [IdentityQualifiedWebhookEventDataIdentityProfessional]
	Professional UnwrapUnsafeWebhookEventUnionDataIdentityProfessional `json:"professional"`
	JSON         struct {
		Demographics respjson.Field
		Emails       respjson.Field
		LinkedinURL  respjson.Field
		Location     respjson.Field
		Phones       respjson.Field
		Professional respjson.Field
		raw          string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataIdentity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataIdentityDemographics is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataIdentityDemographics provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataIdentityDemographics struct {
	FirstName   string `json:"firstName"`
	HasChildren bool   `json:"hasChildren"`
	IsHomeowner bool   `json:"isHomeowner"`
	IsMarried   bool   `json:"isMarried"`
	LastName    string `json:"lastName"`
	JSON        struct {
		FirstName   respjson.Field
		HasChildren respjson.Field
		IsHomeowner respjson.Field
		IsMarried   respjson.Field
		LastName    respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataIdentityDemographics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataIdentityEmails is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataIdentityEmails
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityResolvedWebhookEventDataIdentityEmails
// OfIdentityQualifiedWebhookEventDataIdentityEmails]
type UnwrapUnsafeWebhookEventUnionDataIdentityEmails struct {
	// This field will be present if the value is a
	// [[]IdentityResolvedWebhookEventDataIdentityEmail] instead of an object.
	OfIdentityResolvedWebhookEventDataIdentityEmails []IdentityResolvedWebhookEventDataIdentityEmail `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentityQualifiedWebhookEventDataIdentityEmail] instead of an object.
	OfIdentityQualifiedWebhookEventDataIdentityEmails []IdentityQualifiedWebhookEventDataIdentityEmail `json:",inline"`
	JSON                                              struct {
		OfIdentityResolvedWebhookEventDataIdentityEmails  respjson.Field
		OfIdentityQualifiedWebhookEventDataIdentityEmails respjson.Field
		raw                                               string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataIdentityEmails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataIdentityLocation is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataIdentityLocation provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataIdentityLocation struct {
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	JSON    struct {
		City    respjson.Field
		Country respjson.Field
		State   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataIdentityLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataIdentityProfessional is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataIdentityProfessional provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataIdentityProfessional struct {
	CompanyName string `json:"companyName"`
	JobTitle    string `json:"jobTitle"`
	JSON        struct {
		CompanyName respjson.Field
		JobTitle    respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataIdentityProfessional) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataSession is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataSession
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataSession struct {
	CreatedAt    float64 `json:"createdAt"`
	EndedAt      float64 `json:"endedAt"`
	Fbclid       string  `json:"fbclid"`
	Gclid        string  `json:"gclid"`
	LandingPage  string  `json:"landingPage"`
	LandingTitle string  `json:"landingTitle"`
	// This field is a union of [IdentityResolvedWebhookEventDataSessionNetwork],
	// [IdentityQualifiedWebhookEventDataSessionNetwork],
	// [IdentityEnrichedWebhookEventDataSessionNetwork],
	// [IdentityValidatedWebhookEventDataSessionNetwork],
	// [IdentitySessionFinalizedWebhookEventDataSessionNetwork]
	Network  UnwrapUnsafeWebhookEventUnionDataSessionNetwork `json:"network"`
	Pid      string                                          `json:"pid"`
	Referrer string                                          `json:"referrer"`
	// This field is a union of [IdentityResolvedWebhookEventDataSessionScreen],
	// [IdentityQualifiedWebhookEventDataSessionScreen],
	// [IdentityEnrichedWebhookEventDataSessionScreen],
	// [IdentityValidatedWebhookEventDataSessionScreen],
	// [IdentitySessionFinalizedWebhookEventDataSessionScreen]
	Screen    UnwrapUnsafeWebhookEventUnionDataSessionScreen `json:"screen"`
	Sid       string                                         `json:"sid"`
	Tenant    string                                         `json:"tenant"`
	Timezone  string                                         `json:"timezone"`
	UserAgent string                                         `json:"userAgent"`
	// This field is a union of [IdentityResolvedWebhookEventDataSessionUtm],
	// [IdentityQualifiedWebhookEventDataSessionUtm],
	// [IdentityEnrichedWebhookEventDataSessionUtm],
	// [IdentityValidatedWebhookEventDataSessionUtm],
	// [IdentitySessionFinalizedWebhookEventDataSessionUtm]
	Utm     UnwrapUnsafeWebhookEventUnionDataSessionUtm `json:"utm"`
	Vid     string                                      `json:"vid"`
	Options any                                         `json:"options"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventDataSession].
	Events []IdentitySessionFinalizedWebhookEventDataSessionEvent `json:"events"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventDataSession].
	LastActivity float64 `json:"lastActivity"`
	// This field is from variant [IdentitySessionFinalizedWebhookEventDataSession].
	Enrichments []IdentitySessionFinalizedWebhookEventDataSessionEnrichment `json:"enrichments"`
	JSON        struct {
		CreatedAt    respjson.Field
		EndedAt      respjson.Field
		Fbclid       respjson.Field
		Gclid        respjson.Field
		LandingPage  respjson.Field
		LandingTitle respjson.Field
		Network      respjson.Field
		Pid          respjson.Field
		Referrer     respjson.Field
		Screen       respjson.Field
		Sid          respjson.Field
		Tenant       respjson.Field
		Timezone     respjson.Field
		UserAgent    respjson.Field
		Utm          respjson.Field
		Vid          respjson.Field
		Options      respjson.Field
		Events       respjson.Field
		LastActivity respjson.Field
		Enrichments  respjson.Field
		raw          string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataSessionNetwork is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataSessionNetwork
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataSessionNetwork struct {
	// This field is a union of [IdentityResolvedWebhookEventDataSessionNetworkAsn],
	// [IdentityQualifiedWebhookEventDataSessionNetworkAsn],
	// [IdentityEnrichedWebhookEventDataSessionNetworkAsn],
	// [IdentityValidatedWebhookEventDataSessionNetworkAsn],
	// [IdentitySessionFinalizedWebhookEventDataSessionNetworkAsn]
	Asn UnwrapUnsafeWebhookEventUnionDataSessionNetworkAsn `json:"asn"`
	// This field is a union of
	// [IdentityResolvedWebhookEventDataSessionNetworkBotManagement],
	// [IdentityQualifiedWebhookEventDataSessionNetworkBotManagement],
	// [IdentityEnrichedWebhookEventDataSessionNetworkBotManagement],
	// [IdentityValidatedWebhookEventDataSessionNetworkBotManagement],
	// [IdentitySessionFinalizedWebhookEventDataSessionNetworkBotManagement]
	BotManagement UnwrapUnsafeWebhookEventUnionDataSessionNetworkBotManagement `json:"botManagement"`
	City          string                                                       `json:"city"`
	Colo          string                                                       `json:"colo"`
	Continent     string                                                       `json:"continent"`
	Country       string                                                       `json:"country"`
	IP            string                                                       `json:"ip"`
	IsEu          bool                                                         `json:"isEU"`
	Latitude      string                                                       `json:"latitude"`
	Longitude     string                                                       `json:"longitude"`
	MetroCode     string                                                       `json:"metroCode"`
	PostalCode    string                                                       `json:"postalCode"`
	Region        string                                                       `json:"region"`
	RegionCode    string                                                       `json:"regionCode"`
	JSON          struct {
		Asn           respjson.Field
		BotManagement respjson.Field
		City          respjson.Field
		Colo          respjson.Field
		Continent     respjson.Field
		Country       respjson.Field
		IP            respjson.Field
		IsEu          respjson.Field
		Latitude      respjson.Field
		Longitude     respjson.Field
		MetroCode     respjson.Field
		PostalCode    respjson.Field
		Region        respjson.Field
		RegionCode    respjson.Field
		raw           string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataSessionNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataSessionNetworkAsn is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataSessionNetworkAsn provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataSessionNetworkAsn struct {
	Code float64 `json:"code"`
	Name string  `json:"name"`
	JSON struct {
		Code respjson.Field
		Name respjson.Field
		raw  string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataSessionNetworkAsn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataSessionNetworkBotManagement is an implicit
// subunion of [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataSessionNetworkBotManagement provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataSessionNetworkBotManagement struct {
	CorporateProxy      bool    `json:"corporateProxy"`
	Score               float64 `json:"score"`
	VerifiedBot         bool    `json:"verifiedBot"`
	VerifiedBotCategory string  `json:"verifiedBotCategory"`
	JSON                struct {
		CorporateProxy      respjson.Field
		Score               respjson.Field
		VerifiedBot         respjson.Field
		VerifiedBotCategory respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataSessionNetworkBotManagement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataSessionScreen is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataSessionScreen
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataSessionScreen struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	JSON   struct {
		Height respjson.Field
		Width  respjson.Field
		raw    string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataSessionScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataSessionUtm is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataSessionUtm
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataSessionUtm struct {
	Campaign string `json:"campaign"`
	Content  string `json:"content"`
	Medium   string `json:"medium"`
	Source   string `json:"source"`
	Term     string `json:"term"`
	JSON     struct {
		Campaign respjson.Field
		Content  respjson.Field
		Medium   respjson.Field
		Source   respjson.Field
		Term     respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataSessionUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichment is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion]. UnwrapUnsafeWebhookEventUnionDataEnrichment
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataEnrichment struct {
	// This field is a union of [[]IdentityEnrichedWebhookEventDataEnrichmentCompany],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentCompany]
	Companies UnwrapUnsafeWebhookEventUnionDataEnrichmentCompanies `json:"companies"`
	// This field is a union of [IdentityEnrichedWebhookEventDataEnrichmentPerson],
	// [IdentitySessionFinalizedWebhookEventDataEnrichmentPerson]
	Person UnwrapUnsafeWebhookEventUnionDataEnrichmentPerson `json:"person"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentEmailValidation],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation]
	EmailValidations UnwrapUnsafeWebhookEventUnionDataEnrichmentEmailValidations `json:"emailValidations"`
	JSON             struct {
		Companies        respjson.Field
		Person           respjson.Field
		EmailValidations respjson.Field
		raw              string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentCompanies is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentCompanies provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentCompanies
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentCompanies]
type UnwrapUnsafeWebhookEventUnionDataEnrichmentCompanies struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentCompany] instead of an object.
	OfIdentityEnrichedWebhookEventDataEnrichmentCompanies []IdentityEnrichedWebhookEventDataEnrichmentCompany `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentCompany] instead of an
	// object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentCompanies []IdentitySessionFinalizedWebhookEventDataEnrichmentCompany `json:",inline"`
	JSON                                                          struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentCompanies         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentCompanies respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentCompanies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentPerson is an implicit subunion of
// [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentPerson provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataEnrichmentPerson struct {
	EnrichedAt  string  `json:"enrichedAt"`
	LinkedinURL string  `json:"linkedinUrl"`
	Provider    string  `json:"provider"`
	Connections float64 `json:"connections"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEducation],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation]
	Education UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEducation `json:"education"`
	Email     string                                                     `json:"email"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment]
	Employments UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEmployments `json:"employments"`
	// This field is a union of
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonExperience],
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience]
	Experience UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonExperience `json:"experience"`
	FirstName  string                                                      `json:"firstName"`
	FullName   string                                                      `json:"fullName"`
	Headline   string                                                      `json:"headline"`
	Languages  []string                                                    `json:"languages"`
	LastName   string                                                      `json:"lastName"`
	LinkedinID string                                                      `json:"linkedinId"`
	// This field is a union of
	// [IdentityEnrichedWebhookEventDataEnrichmentPersonLocation],
	// [IdentitySessionFinalizedWebhookEventDataEnrichmentPersonLocation]
	Location          UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonLocation `json:"location"`
	ProfilePictureURL string                                                    `json:"profilePictureUrl"`
	Skills            []string                                                  `json:"skills"`
	Summary           string                                                    `json:"summary"`
	JSON              struct {
		EnrichedAt        respjson.Field
		LinkedinURL       respjson.Field
		Provider          respjson.Field
		Connections       respjson.Field
		Education         respjson.Field
		Email             respjson.Field
		Employments       respjson.Field
		Experience        respjson.Field
		FirstName         respjson.Field
		FullName          respjson.Field
		Headline          respjson.Field
		Languages         respjson.Field
		LastName          respjson.Field
		LinkedinID        respjson.Field
		Location          respjson.Field
		ProfilePictureURL respjson.Field
		Skills            respjson.Field
		Summary           respjson.Field
		raw               string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEducation is an implicit
// subunion of [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEducation provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentPersonEducationArray
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducationArray]
type UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEducation struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEducation] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentPersonEducationArray []IdentityEnrichedWebhookEventDataEnrichmentPersonEducation `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation] instead of
	// an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducationArray []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducation `json:",inline"`
	JSON                                                                     struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentPersonEducationArray         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEducationArray respjson.Field
		raw                                                                      string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEducation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEmployments is an implicit
// subunion of [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEmployments provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentPersonEmployments
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployments]
type UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEmployments struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentPersonEmployments []IdentityEnrichedWebhookEventDataEnrichmentPersonEmployment `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment] instead
	// of an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployments []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployment `json:",inline"`
	JSON                                                                  struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentPersonEmployments         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonEmployments respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonEmployments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonExperience is an implicit
// subunion of [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonExperience provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentPersonExperienceArray
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceArray]
type UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonExperience struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentPersonExperience] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentPersonExperienceArray []IdentityEnrichedWebhookEventDataEnrichmentPersonExperience `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience] instead
	// of an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceArray []IdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperience `json:",inline"`
	JSON                                                                      struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentPersonExperienceArray         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentPersonExperienceArray respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonExperience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonLocation is an implicit
// subunion of [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonLocation provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
type UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonLocation struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Locality    string `json:"locality"`
	PostalCode  string `json:"postalCode"`
	Raw         string `json:"raw"`
	Region      string `json:"region"`
	JSON        struct {
		Country     respjson.Field
		CountryCode respjson.Field
		Locality    respjson.Field
		PostalCode  respjson.Field
		Raw         respjson.Field
		Region      respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentPersonLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapUnsafeWebhookEventUnionDataEnrichmentEmailValidations is an implicit
// subunion of [UnwrapUnsafeWebhookEventUnion].
// UnwrapUnsafeWebhookEventUnionDataEnrichmentEmailValidations provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapUnsafeWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIdentityEnrichedWebhookEventDataEnrichmentEmailValidations
// OfIdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidations]
type UnwrapUnsafeWebhookEventUnionDataEnrichmentEmailValidations struct {
	// This field will be present if the value is a
	// [[]IdentityEnrichedWebhookEventDataEnrichmentEmailValidation] instead of an
	// object.
	OfIdentityEnrichedWebhookEventDataEnrichmentEmailValidations []IdentityEnrichedWebhookEventDataEnrichmentEmailValidation `json:",inline"`
	// This field will be present if the value is a
	// [[]IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation] instead of
	// an object.
	OfIdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidations []IdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidation `json:",inline"`
	JSON                                                                 struct {
		OfIdentityEnrichedWebhookEventDataEnrichmentEmailValidations         respjson.Field
		OfIdentitySessionFinalizedWebhookEventDataEnrichmentEmailValidations respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *UnwrapUnsafeWebhookEventUnionDataEnrichmentEmailValidations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
