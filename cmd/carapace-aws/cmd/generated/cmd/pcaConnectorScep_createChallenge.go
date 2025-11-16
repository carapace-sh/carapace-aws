package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_createChallengeCmd = &cobra.Command{
	Use:   "create-challenge",
	Short: "For general-purpose connectors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_createChallengeCmd).Standalone()

	pcaConnectorScep_createChallengeCmd.Flags().String("client-token", "", "Custom string that can be used to distinguish between calls to the [CreateChallenge](https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_CreateChallenge.html) action.")
	pcaConnectorScep_createChallengeCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) of the connector that you want to create a challenge for.")
	pcaConnectorScep_createChallengeCmd.Flags().String("tags", "", "The key-value pairs to associate with the resource.")
	pcaConnectorScep_createChallengeCmd.MarkFlagRequired("connector-arn")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_createChallengeCmd)
}
