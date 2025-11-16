package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateOrganizationConfigurationCmd = &cobra.Command{
	Use:   "update-organization-configuration",
	Short: "Updates the Amazon Macie configuration settings for an organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateOrganizationConfigurationCmd).Standalone()

		macie2_updateOrganizationConfigurationCmd.Flags().String("auto-enable", "", "Specifies whether to enable Amazon Macie automatically for accounts that are added to the organization in Organizations.")
		macie2_updateOrganizationConfigurationCmd.MarkFlagRequired("auto-enable")
	})
	macie2Cmd.AddCommand(macie2_updateOrganizationConfigurationCmd)
}
