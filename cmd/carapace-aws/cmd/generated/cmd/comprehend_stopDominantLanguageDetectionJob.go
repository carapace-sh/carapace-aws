package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopDominantLanguageDetectionJobCmd = &cobra.Command{
	Use:   "stop-dominant-language-detection-job",
	Short: "Stops a dominant language detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopDominantLanguageDetectionJobCmd).Standalone()

	comprehend_stopDominantLanguageDetectionJobCmd.Flags().String("job-id", "", "The identifier of the dominant language detection job to stop.")
	comprehend_stopDominantLanguageDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_stopDominantLanguageDetectionJobCmd)
}
