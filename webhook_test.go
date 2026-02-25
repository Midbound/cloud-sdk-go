// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package midboundcloud_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/Midbound/cloud-sdk-go"
	"github.com/Midbound/cloud-sdk-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookUnwrap(t *testing.T) {
	client := midboundcloud.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
	)
	payload := []byte(`{"id":"evt_01KH5FW03RQB28JZ4F9D2ZC8ME","created":1739571908123,"data":{"attribution":{"pixelId":"pixelId","prid":"jNNdd","sessionId":"sessionId","type":"website","vid":"vid"},"identity":{"demographics":{"firstName":"firstName","hasChildren":true,"isHomeowner":true,"isMarried":true,"lastName":"lastName"},"emails":[{"type":"personal","value":"value"}],"linkedinUrl":"linkedinUrl","location":{"city":"city","country":"country","state":"state"},"phones":["string"],"professional":{"companyName":"companyName","jobTitle":"jobTitle"}},"session":{"createdAt":0,"endedAt":0,"fbclid":"fbclid","gclid":"gclid","landingPage":"landingPage","landingTitle":"landingTitle","network":{"asn":{"code":0,"name":"name"},"botManagement":{"corporateProxy":true,"score":0,"verifiedBot":true,"verifiedBotCategory":"verifiedBotCategory"},"city":"city","colo":"colo","continent":"continent","country":"country","ip":"ip","isEU":true,"latitude":"latitude","longitude":"longitude","metroCode":"metroCode","postalCode":"postalCode","region":"region","regionCode":"regionCode"},"pid":"pid","referrer":"referrer","screen":{"height":0,"width":0},"sid":"sid","tenant":"tenant","timezone":"timezone","userAgent":"userAgent","utm":{"campaign":"campaign","content":"content","medium":"medium","source":"source","term":"term"},"vid":"vid","options":{"foo":"bar"}}},"type":"identity.resolved"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Error("Failed to sign test webhook message")
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Error("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}
}
