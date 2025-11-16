package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createDataLakeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "create-data-lake-organization-configuration",
	Short: "Automatically enables Amazon Security Lake for new member accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createDataLakeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_createDataLakeOrganizationConfigurationCmd).Standalone()

		securitylake_createDataLakeOrganizationConfigurationCmd.Flags().String("auto-enable-new-account", "", "Enable Security Lake with the specified configuration settings, to begin collecting security data for new accounts in your organization.")
	})
	securitylakeCmd.AddCommand(securitylake_createDataLakeOrganizationConfigurationCmd)
}
