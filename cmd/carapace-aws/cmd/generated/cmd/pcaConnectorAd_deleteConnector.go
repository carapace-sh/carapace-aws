package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_deleteConnectorCmd = &cobra.Command{
	Use:   "delete-connector",
	Short: "Deletes a connector for Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_deleteConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_deleteConnectorCmd).Standalone()

		pcaConnectorAd_deleteConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
		pcaConnectorAd_deleteConnectorCmd.MarkFlagRequired("connector-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_deleteConnectorCmd)
}
