package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_registerDelegatedAdministratorCmd = &cobra.Command{
	Use:   "register-delegated-administrator",
	Short: "Enables the specified member account to administer the Organizations features of the specified Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_registerDelegatedAdministratorCmd).Standalone()

	organizations_registerDelegatedAdministratorCmd.Flags().String("account-id", "", "The account ID number of the member account in the organization to register as a delegated administrator.")
	organizations_registerDelegatedAdministratorCmd.Flags().String("service-principal", "", "The service principal of the Amazon Web Services service for which you want to make the member account a delegated administrator.")
	organizations_registerDelegatedAdministratorCmd.MarkFlagRequired("account-id")
	organizations_registerDelegatedAdministratorCmd.MarkFlagRequired("service-principal")
	organizationsCmd.AddCommand(organizations_registerDelegatedAdministratorCmd)
}
