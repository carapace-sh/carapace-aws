package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_listDirectoryRegistrationsCmd = &cobra.Command{
	Use:   "list-directory-registrations",
	Short: "Lists the directory registrations that you created by using the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API\\_CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration) action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_listDirectoryRegistrationsCmd).Standalone()

	pcaConnectorAd_listDirectoryRegistrationsCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response on each page.")
	pcaConnectorAd_listDirectoryRegistrationsCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_listDirectoryRegistrationsCmd)
}
