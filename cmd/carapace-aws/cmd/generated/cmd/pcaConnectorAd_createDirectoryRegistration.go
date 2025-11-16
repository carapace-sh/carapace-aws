package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_createDirectoryRegistrationCmd = &cobra.Command{
	Use:   "create-directory-registration",
	Short: "Creates a directory registration that authorizes communication between Amazon Web Services Private CA and an Active Directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_createDirectoryRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_createDirectoryRegistrationCmd).Standalone()

		pcaConnectorAd_createDirectoryRegistrationCmd.Flags().String("client-token", "", "Idempotency token.")
		pcaConnectorAd_createDirectoryRegistrationCmd.Flags().String("directory-id", "", "The identifier of the Active Directory.")
		pcaConnectorAd_createDirectoryRegistrationCmd.Flags().String("tags", "", "Metadata assigned to a directory registration consisting of a key-value pair.")
		pcaConnectorAd_createDirectoryRegistrationCmd.MarkFlagRequired("directory-id")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_createDirectoryRegistrationCmd)
}
