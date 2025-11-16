package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_describeTextTranslationJobCmd = &cobra.Command{
	Use:   "describe-text-translation-job",
	Short: "Gets the properties associated with an asynchronous batch translation job including name, ID, status, source and target languages, input/output S3 buckets, and so on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_describeTextTranslationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_describeTextTranslationJobCmd).Standalone()

		translate_describeTextTranslationJobCmd.Flags().String("job-id", "", "The identifier that Amazon Translate generated for the job.")
		translate_describeTextTranslationJobCmd.MarkFlagRequired("job-id")
	})
	translateCmd.AddCommand(translate_describeTextTranslationJobCmd)
}
