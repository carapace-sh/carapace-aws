package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeAffectedAccountsForOrganizationCmd = &cobra.Command{
	Use:   "describe-affected-accounts-for-organization",
	Short: "Returns a list of accounts in the organization from Organizations that are affected by the provided event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeAffectedAccountsForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(health_describeAffectedAccountsForOrganizationCmd).Standalone()

		health_describeAffectedAccountsForOrganizationCmd.Flags().String("event-arn", "", "The unique identifier for the event.")
		health_describeAffectedAccountsForOrganizationCmd.Flags().String("max-results", "", "The maximum number of items to return in one batch, between 10 and 100, inclusive.")
		health_describeAffectedAccountsForOrganizationCmd.Flags().String("next-token", "", "If the results of a search are large, only a portion of the results are returned, and a `nextToken` pagination token is returned in the response.")
		health_describeAffectedAccountsForOrganizationCmd.MarkFlagRequired("event-arn")
	})
	healthCmd.AddCommand(health_describeAffectedAccountsForOrganizationCmd)
}
