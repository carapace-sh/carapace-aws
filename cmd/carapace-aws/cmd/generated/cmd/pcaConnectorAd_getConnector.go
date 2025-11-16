package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_getConnectorCmd = &cobra.Command{
	Use:   "get-connector",
	Short: "Lists information about your connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_getConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_getConnectorCmd).Standalone()

		pcaConnectorAd_getConnectorCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
		pcaConnectorAd_getConnectorCmd.MarkFlagRequired("connector-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_getConnectorCmd)
}
