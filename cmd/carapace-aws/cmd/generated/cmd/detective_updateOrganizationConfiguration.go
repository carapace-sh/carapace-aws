package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_updateOrganizationConfigurationCmd = &cobra.Command{
	Use:   "update-organization-configuration",
	Short: "Updates the configuration for the Organizations integration in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_updateOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_updateOrganizationConfigurationCmd).Standalone()

		detective_updateOrganizationConfigurationCmd.Flags().Bool("auto-enable", false, "Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.")
		detective_updateOrganizationConfigurationCmd.Flags().String("graph-arn", "", "The ARN of the organization behavior graph.")
		detective_updateOrganizationConfigurationCmd.Flags().Bool("no-auto-enable", false, "Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.")
		detective_updateOrganizationConfigurationCmd.MarkFlagRequired("graph-arn")
		detective_updateOrganizationConfigurationCmd.Flag("no-auto-enable").Hidden = true
	})
	detectiveCmd.AddCommand(detective_updateOrganizationConfigurationCmd)
}
