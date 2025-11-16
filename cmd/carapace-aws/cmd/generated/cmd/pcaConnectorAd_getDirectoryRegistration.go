package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_getDirectoryRegistrationCmd = &cobra.Command{
	Use:   "get-directory-registration",
	Short: "A structure that contains information about your directory registration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_getDirectoryRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_getDirectoryRegistrationCmd).Standalone()

		pcaConnectorAd_getDirectoryRegistrationCmd.Flags().String("directory-registration-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).")
		pcaConnectorAd_getDirectoryRegistrationCmd.MarkFlagRequired("directory-registration-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_getDirectoryRegistrationCmd)
}
