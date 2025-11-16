package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEventDetailsForOrganizationCmd = &cobra.Command{
	Use:   "describe-event-details-for-organization",
	Short: "Returns detailed information about one or more specified events for one or more Amazon Web Services accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEventDetailsForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(health_describeEventDetailsForOrganizationCmd).Standalone()

		health_describeEventDetailsForOrganizationCmd.Flags().String("locale", "", "The locale (language) to return information in.")
		health_describeEventDetailsForOrganizationCmd.Flags().String("organization-event-detail-filters", "", "A set of JSON elements that includes the `awsAccountId` and the `eventArn`.")
		health_describeEventDetailsForOrganizationCmd.MarkFlagRequired("organization-event-detail-filters")
	})
	healthCmd.AddCommand(health_describeEventDetailsForOrganizationCmd)
}
