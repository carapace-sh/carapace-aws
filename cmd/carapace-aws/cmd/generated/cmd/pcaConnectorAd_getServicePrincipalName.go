package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_getServicePrincipalNameCmd = &cobra.Command{
	Use:   "get-service-principal-name",
	Short: "Lists the service principal name that the connector uses to authenticate with Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_getServicePrincipalNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_getServicePrincipalNameCmd).Standalone()

		pcaConnectorAd_getServicePrincipalNameCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
		pcaConnectorAd_getServicePrincipalNameCmd.Flags().String("directory-registration-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).")
		pcaConnectorAd_getServicePrincipalNameCmd.MarkFlagRequired("connector-arn")
		pcaConnectorAd_getServicePrincipalNameCmd.MarkFlagRequired("directory-registration-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_getServicePrincipalNameCmd)
}
