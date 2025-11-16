package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeKeyPhrasesDetectionJobCmd = &cobra.Command{
	Use:   "describe-key-phrases-detection-job",
	Short: "Gets the properties associated with a key phrases detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeKeyPhrasesDetectionJobCmd).Standalone()

	comprehend_describeKeyPhrasesDetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
	comprehend_describeKeyPhrasesDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_describeKeyPhrasesDetectionJobCmd)
}
