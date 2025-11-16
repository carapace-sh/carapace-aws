package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_listChallengeMetadataCmd = &cobra.Command{
	Use:   "list-challenge-metadata",
	Short: "Retrieves the challenge metadata for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_listChallengeMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorScep_listChallengeMetadataCmd).Standalone()

		pcaConnectorScep_listChallengeMetadataCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector.")
		pcaConnectorScep_listChallengeMetadataCmd.Flags().String("max-results", "", "The maximum number of objects that you want Connector for SCEP to return for this request.")
		pcaConnectorScep_listChallengeMetadataCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Connector for SCEP returns a `NextToken` value in the response.")
		pcaConnectorScep_listChallengeMetadataCmd.MarkFlagRequired("connector-arn")
	})
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_listChallengeMetadataCmd)
}
