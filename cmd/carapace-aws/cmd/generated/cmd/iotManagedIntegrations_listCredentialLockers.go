package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listCredentialLockersCmd = &cobra.Command{
	Use:   "list-credential-lockers",
	Short: "List information on an existing credential locker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listCredentialLockersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listCredentialLockersCmd).Standalone()

		iotManagedIntegrations_listCredentialLockersCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listCredentialLockersCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listCredentialLockersCmd)
}
