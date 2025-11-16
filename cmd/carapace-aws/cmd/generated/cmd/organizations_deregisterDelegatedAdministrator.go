package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_deregisterDelegatedAdministratorCmd = &cobra.Command{
	Use:   "deregister-delegated-administrator",
	Short: "Removes the specified member Amazon Web Services account as a delegated administrator for the specified Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_deregisterDelegatedAdministratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_deregisterDelegatedAdministratorCmd).Standalone()

		organizations_deregisterDelegatedAdministratorCmd.Flags().String("account-id", "", "The account ID number of the member account in the organization that you want to deregister as a delegated administrator.")
		organizations_deregisterDelegatedAdministratorCmd.Flags().String("service-principal", "", "The service principal name of an Amazon Web Services service for which the account is a delegated administrator.")
		organizations_deregisterDelegatedAdministratorCmd.MarkFlagRequired("account-id")
		organizations_deregisterDelegatedAdministratorCmd.MarkFlagRequired("service-principal")
	})
	organizationsCmd.AddCommand(organizations_deregisterDelegatedAdministratorCmd)
}
