package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listLanguageModelsCmd = &cobra.Command{
	Use:   "list-language-models",
	Short: "Provides a list of custom language models that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listLanguageModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_listLanguageModelsCmd).Standalone()

		transcribe_listLanguageModelsCmd.Flags().String("max-results", "", "The maximum number of custom language models to return in each page of results.")
		transcribe_listLanguageModelsCmd.Flags().String("name-contains", "", "Returns only the custom language models that contain the specified string.")
		transcribe_listLanguageModelsCmd.Flags().String("next-token", "", "If your `ListLanguageModels` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
		transcribe_listLanguageModelsCmd.Flags().String("status-equals", "", "Returns only custom language models with the specified status.")
	})
	transcribeCmd.AddCommand(transcribe_listLanguageModelsCmd)
}
