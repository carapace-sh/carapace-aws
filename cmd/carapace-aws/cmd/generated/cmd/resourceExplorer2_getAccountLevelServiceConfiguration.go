package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getAccountLevelServiceConfigurationCmd = &cobra.Command{
	Use:   "get-account-level-service-configuration",
	Short: "Retrieves the status of your account's Amazon Web Services service access, and validates the service linked role required to access the multi-account search feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getAccountLevelServiceConfigurationCmd).Standalone()

	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getAccountLevelServiceConfigurationCmd)
}
