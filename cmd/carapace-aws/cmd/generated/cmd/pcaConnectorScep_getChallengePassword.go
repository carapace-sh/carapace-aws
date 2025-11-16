package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_getChallengePasswordCmd = &cobra.Command{
	Use:   "get-challenge-password",
	Short: "Retrieves the challenge password for the specified [Challenge](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_Challenge.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_getChallengePasswordCmd).Standalone()

	pcaConnectorScep_getChallengePasswordCmd.Flags().String("challenge-arn", "", "The Amazon Resource Name (ARN) of the challenge.")
	pcaConnectorScep_getChallengePasswordCmd.MarkFlagRequired("challenge-arn")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_getChallengePasswordCmd)
}
