package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "describe-organization-configuration",
	Short: "Returns information about the way your organization is configured in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_describeOrganizationConfigurationCmd).Standalone()

	})
	securityhubCmd.AddCommand(securityhub_describeOrganizationConfigurationCmd)
}
