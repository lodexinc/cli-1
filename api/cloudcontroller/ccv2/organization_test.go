package ccv2_test

import (
	"net/http"

	. "code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Organization", func() {
	var client *Client

	BeforeEach(func() {
		client = NewTestClient()
	})

	Describe("GetOrganizations", func() {
		Context("when no errors are encountered", func() {
			Context("when results are paginated", func() {
				BeforeEach(func() {
					response1 := `{
					"next_url": "/v2/organizations?q=some-query:some-value&page=2",
					"resources": [
						{
							"metadata": {
								"guid": "org-guid-1"
							},
							"entity": {
								"name": "org-1",
								"quota_definition_guid": "some-quota-guid"
							}
						},
						{
							"metadata": {
								"guid": "org-guid-2"
							},
							"entity": {
								"name": "org-2",
								"quota_definition_guid": "some-quota-guid"
							}
						}
					]
				}`
					response2 := `{
					"next_url": null,
					"resources": [
						{
							"metadata": {
								"guid": "org-guid-3"
							},
							"entity": {
								"name": "org-3",
								"quota_definition_guid": "some-quota-guid"
							}
						},
						{
							"metadata": {
								"guid": "org-guid-4"
							},
							"entity": {
								"name": "org-4",
								"quota_definition_guid": "some-quota-guid"
							}
						}
					]
				}`
					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v2/organizations", "q=some-query:some-value"),
							RespondWith(http.StatusOK, response1, http.Header{"X-Cf-Warnings": {"warning-1"}}),
						))
					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v2/organizations", "q=some-query:some-value&page=2"),
							RespondWith(http.StatusOK, response2, http.Header{"X-Cf-Warnings": {"warning-2"}}),
						))
				})

				It("returns paginated results and all warnings", func() {
					orgs, warnings, err := client.GetOrganizations([]Query{{
						Filter:   "some-query",
						Operator: EqualOperator,
						Value:    "some-value",
					}})

					Expect(err).NotTo(HaveOccurred())
					Expect(orgs).To(Equal([]Organization{
						{GUID: "org-guid-1", Name: "org-1", QuotaDefinitionGUID: "some-quota-guid"},
						{GUID: "org-guid-2", Name: "org-2", QuotaDefinitionGUID: "some-quota-guid"},
						{GUID: "org-guid-3", Name: "org-3", QuotaDefinitionGUID: "some-quota-guid"},
						{GUID: "org-guid-4", Name: "org-4", QuotaDefinitionGUID: "some-quota-guid"},
					}))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})
		})

		Context("when an error is encountered", func() {
			BeforeEach(func() {
				response := `{
  "code": 10001,
  "description": "Some Error",
  "error_code": "CF-SomeError"
}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v2/organizations"),
						RespondWith(http.StatusTeapot, response, http.Header{"X-Cf-Warnings": {"warning-1, warning-2"}}),
					))
			})

			It("returns an error and all warnings", func() {
				_, warnings, err := client.GetOrganizations(nil)

				Expect(err).To(MatchError(UnexpectedResponseError{
					ResponseCode: http.StatusTeapot,
					CCErrorResponse: CCErrorResponse{
						Code:        10001,
						Description: "Some Error",
						ErrorCode:   "CF-SomeError",
					},
				}))
				Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
			})
		})
	})
})
