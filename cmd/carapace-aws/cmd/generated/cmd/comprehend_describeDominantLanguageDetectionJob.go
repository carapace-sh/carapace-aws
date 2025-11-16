package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeDominantLanguageDetectionJobCmd = &cobra.Command{
	Use:   "describe-dominant-language-detection-job",
	Short: "Gets the properties associated with a dominant language detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeDominantLanguageDetectionJobCmd).Standalone()

	comprehend_describeDominantLanguageDetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
	comprehend_describeDominantLanguageDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_describeDominantLanguageDetectionJobCmd)
}
