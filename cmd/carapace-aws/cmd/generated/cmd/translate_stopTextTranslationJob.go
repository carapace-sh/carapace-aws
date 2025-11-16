package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_stopTextTranslationJobCmd = &cobra.Command{
	Use:   "stop-text-translation-job",
	Short: "Stops an asynchronous batch translation job that is in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_stopTextTranslationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_stopTextTranslationJobCmd).Standalone()

		translate_stopTextTranslationJobCmd.Flags().String("job-id", "", "The job ID of the job to be stopped.")
		translate_stopTextTranslationJobCmd.MarkFlagRequired("job-id")
	})
	translateCmd.AddCommand(translate_stopTextTranslationJobCmd)
}
