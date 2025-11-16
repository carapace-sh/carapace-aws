package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_registerDataLakeDelegatedAdministratorCmd = &cobra.Command{
	Use:   "register-data-lake-delegated-administrator",
	Short: "Designates the Amazon Security Lake delegated administrator account for the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_registerDataLakeDelegatedAdministratorCmd).Standalone()

	securitylake_registerDataLakeDelegatedAdministratorCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Security Lake delegated administrator.")
	securitylake_registerDataLakeDelegatedAdministratorCmd.MarkFlagRequired("account-id")
	securitylakeCmd.AddCommand(securitylake_registerDataLakeDelegatedAdministratorCmd)
}
