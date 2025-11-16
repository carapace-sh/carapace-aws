package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEventsForOrganizationCmd = &cobra.Command{
	Use:   "describe-events-for-organization",
	Short: "Returns information about events across your organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEventsForOrganizationCmd).Standalone()

	health_describeEventsForOrganizationCmd.Flags().String("filter", "", "Values to narrow the results returned.")
	health_describeEventsForOrganizationCmd.Flags().String("locale", "", "The locale (language) to return information in.")
	health_describeEventsForOrganizationCmd.Flags().String("max-results", "", "The maximum number of items to return in one batch, between 10 and 100, inclusive.")
	health_describeEventsForOrganizationCmd.Flags().String("next-token", "", "If the results of a search are large, only a portion of the results are returned, and a `nextToken` pagination token is returned in the response.")
	healthCmd.AddCommand(health_describeEventsForOrganizationCmd)
}
