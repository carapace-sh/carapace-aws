package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteDataLakeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "delete-data-lake-organization-configuration",
	Short: "Turns off automatic enablement of Amazon Security Lake for member accounts that are added to an organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteDataLakeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_deleteDataLakeOrganizationConfigurationCmd).Standalone()

		securitylake_deleteDataLakeOrganizationConfigurationCmd.Flags().String("auto-enable-new-account", "", "Turns off automatic enablement of Security Lake for member accounts that are added to an organization.")
	})
	securitylakeCmd.AddCommand(securitylake_deleteDataLakeOrganizationConfigurationCmd)
}
