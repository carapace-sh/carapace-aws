package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_startTextTranslationJobCmd = &cobra.Command{
	Use:   "start-text-translation-job",
	Short: "Starts an asynchronous batch translation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_startTextTranslationJobCmd).Standalone()

	translate_startTextTranslationJobCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	translate_startTextTranslationJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM) role that grants Amazon Translate read access to your input data.")
	translate_startTextTranslationJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input documents for the translation job.")
	translate_startTextTranslationJobCmd.Flags().String("job-name", "", "The name of the batch translation job to be performed.")
	translate_startTextTranslationJobCmd.Flags().String("output-data-config", "", "Specifies the S3 folder to which your job output will be saved.")
	translate_startTextTranslationJobCmd.Flags().String("parallel-data-names", "", "The name of a parallel data resource to add to the translation job.")
	translate_startTextTranslationJobCmd.Flags().String("settings", "", "Settings to configure your translation output.")
	translate_startTextTranslationJobCmd.Flags().String("source-language-code", "", "The language code of the input language.")
	translate_startTextTranslationJobCmd.Flags().String("target-language-codes", "", "The target languages of the translation job.")
	translate_startTextTranslationJobCmd.Flags().String("terminology-names", "", "The name of a custom terminology resource to add to the translation job.")
	translate_startTextTranslationJobCmd.MarkFlagRequired("client-token")
	translate_startTextTranslationJobCmd.MarkFlagRequired("data-access-role-arn")
	translate_startTextTranslationJobCmd.MarkFlagRequired("input-data-config")
	translate_startTextTranslationJobCmd.MarkFlagRequired("output-data-config")
	translate_startTextTranslationJobCmd.MarkFlagRequired("source-language-code")
	translate_startTextTranslationJobCmd.MarkFlagRequired("target-language-codes")
	translateCmd.AddCommand(translate_startTextTranslationJobCmd)
}
