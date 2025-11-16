package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deregisterDataLakeDelegatedAdministratorCmd = &cobra.Command{
	Use:   "deregister-data-lake-delegated-administrator",
	Short: "Deletes the Amazon Security Lake delegated administrator account for the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deregisterDataLakeDelegatedAdministratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_deregisterDataLakeDelegatedAdministratorCmd).Standalone()

	})
	securitylakeCmd.AddCommand(securitylake_deregisterDataLakeDelegatedAdministratorCmd)
}
