package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeOrganizationHealthCmd = &cobra.Command{
	Use:   "describe-organization-health",
	Short: "Returns active insights, predictive insights, and resource hours analyzed in last hour.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeOrganizationHealthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeOrganizationHealthCmd).Standalone()

		devopsGuru_describeOrganizationHealthCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account.")
		devopsGuru_describeOrganizationHealthCmd.Flags().String("organizational-unit-ids", "", "The ID of the organizational unit.")
	})
	devopsGuruCmd.AddCommand(devopsGuru_describeOrganizationHealthCmd)
}
