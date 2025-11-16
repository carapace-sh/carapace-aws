package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_deleteConnectorCmd = &cobra.Command{
	Use:   "delete-connector",
	Short: "Deletes the specified [Connector](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_Connector.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_deleteConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorScep_deleteConnectorCmd).Standalone()

		pcaConnectorScep_deleteConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector to delete.")
		pcaConnectorScep_deleteConnectorCmd.MarkFlagRequired("connector-arn")
	})
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_deleteConnectorCmd)
}
