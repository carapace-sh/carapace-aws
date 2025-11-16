package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_createConnectorCmd = &cobra.Command{
	Use:   "create-connector",
	Short: "Creates a connector between Amazon Web Services Private CA and an Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_createConnectorCmd).Standalone()

	pcaConnectorAd_createConnectorCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of the certificate authority being used.")
	pcaConnectorAd_createConnectorCmd.Flags().String("client-token", "", "Idempotency token.")
	pcaConnectorAd_createConnectorCmd.Flags().String("directory-id", "", "The identifier of the Active Directory.")
	pcaConnectorAd_createConnectorCmd.Flags().String("tags", "", "Metadata assigned to a connector consisting of a key-value pair.")
	pcaConnectorAd_createConnectorCmd.Flags().String("vpc-information", "", "Information about your VPC and security groups used with the connector.")
	pcaConnectorAd_createConnectorCmd.MarkFlagRequired("certificate-authority-arn")
	pcaConnectorAd_createConnectorCmd.MarkFlagRequired("directory-id")
	pcaConnectorAd_createConnectorCmd.MarkFlagRequired("vpc-information")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_createConnectorCmd)
}
