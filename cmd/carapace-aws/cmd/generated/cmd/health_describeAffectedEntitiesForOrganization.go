package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeAffectedEntitiesForOrganizationCmd = &cobra.Command{
	Use:   "describe-affected-entities-for-organization",
	Short: "Returns a list of entities that have been affected by one or more events for one or more accounts in your organization in Organizations, based on the filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeAffectedEntitiesForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(health_describeAffectedEntitiesForOrganizationCmd).Standalone()

		health_describeAffectedEntitiesForOrganizationCmd.Flags().String("locale", "", "The locale (language) to return information in.")
		health_describeAffectedEntitiesForOrganizationCmd.Flags().String("max-results", "", "The maximum number of items to return in one batch, between 10 and 100, inclusive.")
		health_describeAffectedEntitiesForOrganizationCmd.Flags().String("next-token", "", "If the results of a search are large, only a portion of the results are returned, and a `nextToken` pagination token is returned in the response.")
		health_describeAffectedEntitiesForOrganizationCmd.Flags().String("organization-entity-account-filters", "", "A JSON set of elements including the `awsAccountId`, `eventArn` and a set of `statusCodes`.")
		health_describeAffectedEntitiesForOrganizationCmd.Flags().String("organization-entity-filters", "", "A JSON set of elements including the `awsAccountId` and the `eventArn`.")
	})
	healthCmd.AddCommand(health_describeAffectedEntitiesForOrganizationCmd)
}
