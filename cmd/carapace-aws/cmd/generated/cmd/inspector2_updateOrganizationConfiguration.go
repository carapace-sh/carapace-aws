package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateOrganizationConfigurationCmd = &cobra.Command{
	Use:   "update-organization-configuration",
	Short: "Updates the configurations for your Amazon Inspector organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateOrganizationConfigurationCmd).Standalone()

	inspector2_updateOrganizationConfigurationCmd.Flags().String("auto-enable", "", "Defines which scan types are enabled automatically for new members of your Amazon Inspector organization.")
	inspector2_updateOrganizationConfigurationCmd.MarkFlagRequired("auto-enable")
	inspector2Cmd.AddCommand(inspector2_updateOrganizationConfigurationCmd)
}
