package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_deleteServicePrincipalNameCmd = &cobra.Command{
	Use:   "delete-service-principal-name",
	Short: "Deletes the service principal name (SPN) used by a connector to authenticate with your Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_deleteServicePrincipalNameCmd).Standalone()

	pcaConnectorAd_deleteServicePrincipalNameCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
	pcaConnectorAd_deleteServicePrincipalNameCmd.Flags().String("directory-registration-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).")
	pcaConnectorAd_deleteServicePrincipalNameCmd.MarkFlagRequired("connector-arn")
	pcaConnectorAd_deleteServicePrincipalNameCmd.MarkFlagRequired("directory-registration-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_deleteServicePrincipalNameCmd)
}
