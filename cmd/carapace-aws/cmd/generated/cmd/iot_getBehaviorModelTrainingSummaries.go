package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getBehaviorModelTrainingSummariesCmd = &cobra.Command{
	Use:   "get-behavior-model-training-summaries",
	Short: "Returns a Device Defender's ML Detect Security Profile training model's status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getBehaviorModelTrainingSummariesCmd).Standalone()

	iot_getBehaviorModelTrainingSummariesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_getBehaviorModelTrainingSummariesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_getBehaviorModelTrainingSummariesCmd.Flags().String("security-profile-name", "", "The name of the security profile.")
	iotCmd.AddCommand(iot_getBehaviorModelTrainingSummariesCmd)
}
