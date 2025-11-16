package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_describeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "describe-organization-configuration",
	Short: "Returns information about the account selected as the delegated administrator for GuardDuty.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_describeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_describeOrganizationConfigurationCmd).Standalone()

		guardduty_describeOrganizationConfigurationCmd.Flags().String("detector-id", "", "The detector ID of the delegated administrator for which you need to retrieve the information.")
		guardduty_describeOrganizationConfigurationCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
		guardduty_describeOrganizationConfigurationCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
		guardduty_describeOrganizationConfigurationCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_describeOrganizationConfigurationCmd)
}
