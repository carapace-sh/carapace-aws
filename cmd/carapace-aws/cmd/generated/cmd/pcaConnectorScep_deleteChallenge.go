package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_deleteChallengeCmd = &cobra.Command{
	Use:   "delete-challenge",
	Short: "Deletes the specified [Challenge](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_Challenge.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_deleteChallengeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorScep_deleteChallengeCmd).Standalone()

		pcaConnectorScep_deleteChallengeCmd.Flags().String("challenge-arn", "", "The Amazon Resource Name (ARN) of the challenge password to delete.")
		pcaConnectorScep_deleteChallengeCmd.MarkFlagRequired("challenge-arn")
	})
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_deleteChallengeCmd)
}
