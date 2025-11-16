package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_getDataLakeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "get-data-lake-organization-configuration",
	Short: "Retrieves the configuration that will be automatically set up for accounts added to the organization after the organization has onboarded to Amazon Security Lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_getDataLakeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_getDataLakeOrganizationConfigurationCmd).Standalone()

	})
	securitylakeCmd.AddCommand(securitylake_getDataLakeOrganizationConfigurationCmd)
}
