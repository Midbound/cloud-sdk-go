// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package midboundcloud

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Midbound/cloud-sdk-go/internal/apijson"
	"github.com/Midbound/cloud-sdk-go/internal/requestconfig"
	"github.com/Midbound/cloud-sdk-go/option"
	"github.com/Midbound/cloud-sdk-go/packages/respjson"
)

// HealthService contains methods and other services that help with interacting
// with the midbound-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Returns the current health status of the API
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthCheckResponse struct {
	// Service health status
	//
	// Any of "ok", "degraded", "unhealthy".
	Status HealthCheckResponseStatus `json:"status" api:"required"`
	// Current server timestamp (ISO)
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// API version
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Service health status
type HealthCheckResponseStatus string

const (
	HealthCheckResponseStatusOk        HealthCheckResponseStatus = "ok"
	HealthCheckResponseStatusDegraded  HealthCheckResponseStatus = "degraded"
	HealthCheckResponseStatusUnhealthy HealthCheckResponseStatus = "unhealthy"
)
