package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeOrganizationOverviewCmd = &cobra.Command{
	Use:   "describe-organization-overview",
	Short: "Returns an overview of your organization's history based on the specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeOrganizationOverviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeOrganizationOverviewCmd).Standalone()

		devopsGuru_describeOrganizationOverviewCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account.")
		devopsGuru_describeOrganizationOverviewCmd.Flags().String("from-time", "", "The start of the time range passed in.")
		devopsGuru_describeOrganizationOverviewCmd.Flags().String("organizational-unit-ids", "", "The ID of the organizational unit.")
		devopsGuru_describeOrganizationOverviewCmd.Flags().String("to-time", "", "The end of the time range passed in.")
		devopsGuru_describeOrganizationOverviewCmd.MarkFlagRequired("from-time")
	})
	devopsGuruCmd.AddCommand(devopsGuru_describeOrganizationOverviewCmd)
}
