package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_listServicePrincipalNamesCmd = &cobra.Command{
	Use:   "list-service-principal-names",
	Short: "Lists the service principal names that the connector uses to authenticate with Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_listServicePrincipalNamesCmd).Standalone()

	pcaConnectorAd_listServicePrincipalNamesCmd.Flags().String("directory-registration-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).")
	pcaConnectorAd_listServicePrincipalNamesCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response on each page.")
	pcaConnectorAd_listServicePrincipalNamesCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
	pcaConnectorAd_listServicePrincipalNamesCmd.MarkFlagRequired("directory-registration-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_listServicePrincipalNamesCmd)
}
