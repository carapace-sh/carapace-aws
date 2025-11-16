package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateOrganizationConfigurationCmd = &cobra.Command{
	Use:   "update-organization-configuration",
	Short: "Configures the delegated administrator account with the provided values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateOrganizationConfigurationCmd).Standalone()

	guardduty_updateOrganizationConfigurationCmd.Flags().Bool("auto-enable", false, "Represents whether to automatically enable member accounts in the organization.")
	guardduty_updateOrganizationConfigurationCmd.Flags().String("auto-enable-organization-members", "", "Indicates the auto-enablement configuration of GuardDuty for the member accounts in the organization.")
	guardduty_updateOrganizationConfigurationCmd.Flags().String("data-sources", "", "Describes which data sources will be updated.")
	guardduty_updateOrganizationConfigurationCmd.Flags().String("detector-id", "", "The ID of the detector that configures the delegated administrator.")
	guardduty_updateOrganizationConfigurationCmd.Flags().String("features", "", "A list of features that will be configured for the organization.")
	guardduty_updateOrganizationConfigurationCmd.Flags().Bool("no-auto-enable", false, "Represents whether to automatically enable member accounts in the organization.")
	guardduty_updateOrganizationConfigurationCmd.MarkFlagRequired("detector-id")
	guardduty_updateOrganizationConfigurationCmd.Flag("no-auto-enable").Hidden = true
	guarddutyCmd.AddCommand(guardduty_updateOrganizationConfigurationCmd)
}
