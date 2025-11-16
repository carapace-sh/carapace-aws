package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_getChallengeMetadataCmd = &cobra.Command{
	Use:   "get-challenge-metadata",
	Short: "Retrieves the metadata for the specified [Challenge](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_Challenge.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_getChallengeMetadataCmd).Standalone()

	pcaConnectorScep_getChallengeMetadataCmd.Flags().String("challenge-arn", "", "The Amazon Resource Name (ARN) of the challenge.")
	pcaConnectorScep_getChallengeMetadataCmd.MarkFlagRequired("challenge-arn")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_getChallengeMetadataCmd)
}
