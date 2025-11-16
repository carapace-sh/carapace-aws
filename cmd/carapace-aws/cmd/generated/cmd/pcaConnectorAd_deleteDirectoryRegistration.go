package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_deleteDirectoryRegistrationCmd = &cobra.Command{
	Use:   "delete-directory-registration",
	Short: "Deletes a directory registration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_deleteDirectoryRegistrationCmd).Standalone()

	pcaConnectorAd_deleteDirectoryRegistrationCmd.Flags().String("directory-registration-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).")
	pcaConnectorAd_deleteDirectoryRegistrationCmd.MarkFlagRequired("directory-registration-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_deleteDirectoryRegistrationCmd)
}
