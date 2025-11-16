package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectDominantLanguageCmd = &cobra.Command{
	Use:   "detect-dominant-language",
	Short: "Determines the dominant language of the input text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectDominantLanguageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_detectDominantLanguageCmd).Standalone()

		comprehend_detectDominantLanguageCmd.Flags().String("text", "", "A UTF-8 text string.")
		comprehend_detectDominantLanguageCmd.MarkFlagRequired("text")
	})
	comprehendCmd.AddCommand(comprehend_detectDominantLanguageCmd)
}
