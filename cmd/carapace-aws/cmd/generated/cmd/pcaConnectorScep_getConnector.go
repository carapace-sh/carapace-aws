package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_getConnectorCmd = &cobra.Command{
	Use:   "get-connector",
	Short: "Retrieves details about the specified [Connector](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_Connector.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_getConnectorCmd).Standalone()

	pcaConnectorScep_getConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector.")
	pcaConnectorScep_getConnectorCmd.MarkFlagRequired("connector-arn")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_getConnectorCmd)
}
