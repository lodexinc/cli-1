package v2action_test

import (
	. "code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/v2actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Space Summary Actions", func() {
	var (
		actor                     Actor
		fakeCloudControllerClient *v2actionfakes.FakeCloudControllerClient
		spaceSummary              SpaceSummary
		warnings                  Warnings
		err                       error
		// expectedErr               error
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v2actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil)
	})

	JustBeforeEach(func() {
		spaceSummary, warnings, err = actor.GetSpaceSummaryByOrganizationAndName("some-org-guid", "some-space")
	})

	Describe("GetOrganizationSummaryByName", func() {
		Context("when no errors are encountered", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetOrganizationReturns(
					ccv2.Organization{
						GUID: "some-org-guid",
						Name: "some-org",
					},
					ccv2.Warnings{"warning-1", "warning-2"},
					nil)

				fakeCloudControllerClient.GetSpacesReturns(
					[]ccv2.Space{
						{
							GUID: "some-space-guid",
							Name: "some-space",
							SpaceQuotaDefinitionGUID: "some-space-quota-guid",
						},
					},
					ccv2.Warnings{"warning-3", "warning-4"},
					nil,
				)

				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv2.Application{
						{
							Name: "some-app-2",
						},
						{
							Name: "some-app-1",
						},
					},
					ccv2.Warnings{"warning-5", "warning-6"},
					nil,
				)

				fakeCloudControllerClient.GetSpaceServiceInstancesReturns(
					[]ccv2.ServiceInstance{
						{
							GUID: "some-service-instance-guid-2",
							Name: "some-service-instance-2",
						},
						{
							GUID: "some-service-instance-guid-1",
							Name: "some-service-instance-1",
						},
					},
					ccv2.Warnings{"warning-7", "warning-8"},
					nil,
				)

				fakeCloudControllerClient.GetSpaceQuotaReturns(
					ccv2.SpaceQuota{
						GUID: "some-space-quota-guid",
						Name: "some-space-quota",
					},
					ccv2.Warnings{"warning-9", "warning-10"},
					nil,
				)

				fakeCloudControllerClient.GetSpaceRunningSecurityGroupsBySpaceReturns(
					[]ccv2.SecurityGroup{
						{
							Name: "some-shared-security-group",
						},
						{
							Name: "some-running-security-group",
						},
					},
					ccv2.Warnings{"warning-11", "warning-12"},
					nil,
				)

				fakeCloudControllerClient.GetSpaceStagingSecurityGroupsBySpaceReturns(
					[]ccv2.SecurityGroup{
						{
							Name: "some-staging-security-group",
						},
						{
							Name: "some-shared-security-group",
						},
					},
					ccv2.Warnings{"warning-13", "warning-14"},
					nil,
				)
			})

			It("returns the organization summary and all warnings", func() {
				Expect(err).NotTo(HaveOccurred())

				Expect(warnings).To(ConsistOf([]string{
					"warning-1",
					"warning-2",
					"warning-3",
					"warning-4",
					"warning-5",
					"warning-6",
					"warning-7",
					"warning-8",
					"warning-9",
					"warning-10",
					"warning-11",
					"warning-12",
					"warning-13",
					"warning-14",
				}))

				Expect(spaceSummary).To(Equal(SpaceSummary{
					SpaceName:            "some-space",
					OrgName:              "some-org",
					AppNames:             []string{"some-app-1", "some-app-2"},
					ServiceInstanceNames: []string{"some-service-instance-1", "some-service-instance-2"},
					SpaceQuotaName:       "some-space-quota",
					SecurityGroupNames:   []string{"some-running-security-group", "some-shared-security-group", "some-staging-security-group"},
				}))

				Expect(fakeCloudControllerClient.GetOrganizationCallCount()).To(Equal(1))
				Expect(fakeCloudControllerClient.GetOrganizationArgsForCall(0)).To(Equal("some-org-guid"))

				Expect(fakeCloudControllerClient.GetSpacesCallCount()).To(Equal(1))
				query := fakeCloudControllerClient.GetSpacesArgsForCall(0)
				Expect(query).To(ConsistOf(
					ccv2.Query{
						Filter:   ccv2.NameFilter,
						Operator: ccv2.EqualOperator,
						Value:    "some-space",
					},
					ccv2.Query{
						Filter:   ccv2.OrganizationGUIDFilter,
						Operator: ccv2.EqualOperator,
						Value:    "some-org-guid",
					},
				))

				Expect(fakeCloudControllerClient.GetApplicationsCallCount()).To(Equal(1))
				query = fakeCloudControllerClient.GetApplicationsArgsForCall(0)
				Expect(query).To(ConsistOf(
					ccv2.Query{
						Filter:   ccv2.SpaceGUIDFilter,
						Operator: ccv2.EqualOperator,
						Value:    "some-space-guid",
					},
				))

				Expect(fakeCloudControllerClient.GetSpaceServiceInstancesCallCount()).To(Equal(1))
				spaceGUID, includeUserProvidedServices, query := fakeCloudControllerClient.GetSpaceServiceInstancesArgsForCall(0)
				Expect(spaceGUID).To(Equal("some-space-guid"))
				Expect(includeUserProvidedServices).To(BeTrue())
				Expect(query).To(BeNil())

				Expect(fakeCloudControllerClient.GetSpaceQuotaCallCount()).To(Equal(1))
				spaceQuotaGUID := fakeCloudControllerClient.GetSpaceQuotaArgsForCall(0)
				Expect(spaceQuotaGUID).To(Equal("some-space-quota-guid"))

				Expect(fakeCloudControllerClient.GetSpaceRunningSecurityGroupsBySpaceCallCount()).To(Equal(1))
				spaceGUID = fakeCloudControllerClient.GetSpaceRunningSecurityGroupsBySpaceArgsForCall(0)
				Expect(spaceGUID).To(Equal("some-space-guid"))

				Expect(fakeCloudControllerClient.GetSpaceStagingSecurityGroupsBySpaceCallCount()).To(Equal(1))
				spaceGUID = fakeCloudControllerClient.GetSpaceStagingSecurityGroupsBySpaceArgsForCall(0)
				Expect(spaceGUID).To(Equal("some-space-guid"))
			})

			// })

			// Context("when an error is encountered getting the organization", func() {
			// BeforeEach(func() {
			// 	expectedErr = errors.New("get-orgs-error")
			// 	fakeCloudControllerClient.GetOrganizationsReturns(
			// 		[]ccv2.Organization{},
			// 		ccv2.Warnings{
			// 			"warning-1",
			// 			"warning-2",
			// 		},
			// 		expectedErr,
			// 	)
			// })

			// It("returns the error and all warnings", func() {
			// 	Expect(err).To(MatchError(expectedErr))
			// 	Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
			// })
			// })

			// Context("when the organization exists", func() {
			// BeforeEach(func() {
			// 	fakeCloudControllerClient.GetOrganizationsReturns(
			// 		[]ccv2.Organization{
			// 			{GUID: "some-org-guid"},
			// 		},
			// 		ccv2.Warnings{
			// 			"warning-1",
			// 			"warning-2",
			// 		},
			// 		nil,
			// 	)
			// })

			// Context("when an error is encountered getting the organization domains", func() {
			// 	BeforeEach(func() {
			// 		expectedErr = errors.New("shared domains error")
			// 		fakeCloudControllerClient.GetSharedDomainsReturns([]ccv2.Domain{}, ccv2.Warnings{"shared domains warning"}, expectedErr)
			// 	})

			// 	It("returns that error and all warnings", func() {
			// 		Expect(err).To(MatchError(expectedErr))
			// 		Expect(warnings).To(ConsistOf("warning-1", "warning-2", "shared domains warning"))
			// 	})
			// })

			// Context("when an error is encountered getting the organization quota", func() {
			// 	BeforeEach(func() {
			// 		expectedErr = errors.New("some org quota error")
			// 		fakeCloudControllerClient.GetOrganizationQuotaReturns(ccv2.OrganizationQuota{}, ccv2.Warnings{"quota warning"}, expectedErr)
			// 	})

			// 	It("returns the error and all warnings", func() {
			// 		Expect(err).To(MatchError(expectedErr))
			// 		Expect(warnings).To(ConsistOf("warning-1", "warning-2", "quota warning"))
			// 	})
			// })

			// Context("when an error is encountered getting the organization spaces", func() {
			// 	BeforeEach(func() {
			// 		expectedErr = errors.New("cc-get-spaces-error")
			// 		fakeCloudControllerClient.GetSpacesReturns(
			// 			[]ccv2.Space{},
			// 			ccv2.Warnings{"spaces warning"},
			// 			expectedErr,
			// 		)
			// 	})

			// 	It("returns the error and all warnings", func() {
			// 		Expect(err).To(MatchError(expectedErr))
			// 		Expect(warnings).To(ConsistOf("warning-1", "warning-2", "spaces warning"))
			// 	})
			// })
		})
	})
})
