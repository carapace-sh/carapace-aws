package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_createConnectorCmd = &cobra.Command{
	Use:   "create-connector",
	Short: "Creates a SCEP connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_createConnectorCmd).Standalone()

	pcaConnectorScep_createConnectorCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services Private Certificate Authority certificate authority to use with this connector.")
	pcaConnectorScep_createConnectorCmd.Flags().String("client-token", "", "Custom string that can be used to distinguish between calls to the [CreateChallenge](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_CreateChallenge.html) action.")
	pcaConnectorScep_createConnectorCmd.Flags().String("mobile-device-management", "", "If you don't supply a value, by default Connector for SCEP creates a connector for general-purpose use.")
	pcaConnectorScep_createConnectorCmd.Flags().String("tags", "", "The key-value pairs to associate with the resource.")
	pcaConnectorScep_createConnectorCmd.MarkFlagRequired("certificate-authority-arn")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_createConnectorCmd)
}
