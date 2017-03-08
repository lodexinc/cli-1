package v2action

import (
	"sort"

	"code.cloudfoundry.org/cli/util/generic"
)

type SpaceSummary struct {
	SpaceName            string
	OrgName              string
	AppNames             []string
	ServiceInstanceNames []string
	SpaceQuotaName       string
	SecurityGroupNames   []string
	// SecurityGroupRules   []SecurityGroupRule
}

// type SecurityGroupRule struct {
// 	Description string
// 	Destination string
// 	Lifecycle   string
// 	Name        string
// 	Port        int
// 	Protocol    string
// }

func (actor Actor) GetSpaceSummaryByOrganizationAndName(orgGUID string, name string) (SpaceSummary, Warnings, error) {
	var allWarnings Warnings

	org, warnings, _ := actor.GetOrganization(orgGUID)
	allWarnings = append(allWarnings, warnings...)
	// if err != nil {
	// 	return SpaceSummary{}, nil, nil
	// }

	space, warnings, _ := actor.GetSpaceByOrganizationAndName(org.GUID, name)
	allWarnings = append(allWarnings, warnings...)
	// if err != nil {
	// 	return SpaceSummary{}, nil, nil
	// }

	apps, warnings, _ := actor.GetApplicationsBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)

	appNames := make([]string, len(apps))
	for i, app := range apps {
		appNames[i] = app.Name
	}
	sort.Strings(appNames)

	serviceInstances, warnings, _ := actor.GetServiceInstancesBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)

	serviceInstanceNames := make([]string, len(serviceInstances))
	for i, serviceInstance := range serviceInstances {
		serviceInstanceNames[i] = serviceInstance.Name
	}
	sort.Strings(serviceInstanceNames)

	spaceQuota, warnings, _ := actor.GetSpaceQuota(space.SpaceQuotaDefinitionGUID)
	allWarnings = append(allWarnings, warnings...)

	securityGroups, warnings, _ := actor.GetSpaceRunningSecurityGroupsBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)
	var securityGroupNames []string
	for _, securityGroup := range securityGroups {
		securityGroupNames = append(securityGroupNames, securityGroup.Name)
	}

	securityGroups, warnings, _ = actor.GetSpaceStagingSecurityGroupsBySpace(space.GUID)
	allWarnings = append(allWarnings, warnings...)
	for _, securityGroup := range securityGroups {
		securityGroupNames = append(securityGroupNames, securityGroup.Name)
	}

	securityGroupNames = generic.SortAndUniquifyStringSlice(securityGroupNames)

	spaceSummary := SpaceSummary{
		SpaceName:            name,
		OrgName:              org.Name,
		AppNames:             appNames,
		ServiceInstanceNames: serviceInstanceNames,
		SpaceQuotaName:       spaceQuota.Name,
		SecurityGroupNames:   securityGroupNames,
	}

	return spaceSummary, allWarnings, nil
}
