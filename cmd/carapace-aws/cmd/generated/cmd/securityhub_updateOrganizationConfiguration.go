package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateOrganizationConfigurationCmd = &cobra.Command{
	Use:   "update-organization-configuration",
	Short: "Updates the configuration of your organization in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateOrganizationConfigurationCmd).Standalone()

		securityhub_updateOrganizationConfigurationCmd.Flags().Bool("auto-enable", false, "Whether to automatically enable Security Hub in new member accounts when they join the organization.")
		securityhub_updateOrganizationConfigurationCmd.Flags().String("auto-enable-standards", "", "Whether to automatically enable Security Hub [default standards](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-enable-disable.html) in new member accounts when they join the organization.")
		securityhub_updateOrganizationConfigurationCmd.Flags().Bool("no-auto-enable", false, "Whether to automatically enable Security Hub in new member accounts when they join the organization.")
		securityhub_updateOrganizationConfigurationCmd.Flags().String("organization-configuration", "", "")
		securityhub_updateOrganizationConfigurationCmd.MarkFlagRequired("auto-enable")
		securityhub_updateOrganizationConfigurationCmd.Flag("no-auto-enable").Hidden = true
		securityhub_updateOrganizationConfigurationCmd.MarkFlagRequired("no-auto-enable")
	})
	securityhubCmd.AddCommand(securityhub_updateOrganizationConfigurationCmd)
}
