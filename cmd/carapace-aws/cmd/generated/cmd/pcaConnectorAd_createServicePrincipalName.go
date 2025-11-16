package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_createServicePrincipalNameCmd = &cobra.Command{
	Use:   "create-service-principal-name",
	Short: "Creates a service principal name (SPN) for the service account in Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_createServicePrincipalNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_createServicePrincipalNameCmd).Standalone()

		pcaConnectorAd_createServicePrincipalNameCmd.Flags().String("client-token", "", "Idempotency token.")
		pcaConnectorAd_createServicePrincipalNameCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
		pcaConnectorAd_createServicePrincipalNameCmd.Flags().String("directory-registration-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html).")
		pcaConnectorAd_createServicePrincipalNameCmd.MarkFlagRequired("connector-arn")
		pcaConnectorAd_createServicePrincipalNameCmd.MarkFlagRequired("directory-registration-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_createServicePrincipalNameCmd)
}
