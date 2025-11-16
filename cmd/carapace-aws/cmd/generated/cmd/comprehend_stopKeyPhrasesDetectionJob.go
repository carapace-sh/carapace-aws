package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopKeyPhrasesDetectionJobCmd = &cobra.Command{
	Use:   "stop-key-phrases-detection-job",
	Short: "Stops a key phrases detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopKeyPhrasesDetectionJobCmd).Standalone()

	comprehend_stopKeyPhrasesDetectionJobCmd.Flags().String("job-id", "", "The identifier of the key phrases detection job to stop.")
	comprehend_stopKeyPhrasesDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_stopKeyPhrasesDetectionJobCmd)
}
