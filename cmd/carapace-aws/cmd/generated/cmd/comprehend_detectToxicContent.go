package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectToxicContentCmd = &cobra.Command{
	Use:   "detect-toxic-content",
	Short: "Performs toxicity analysis on the list of text strings that you provide as input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectToxicContentCmd).Standalone()

	comprehend_detectToxicContentCmd.Flags().String("language-code", "", "The language of the input text.")
	comprehend_detectToxicContentCmd.Flags().String("text-segments", "", "A list of up to 10 text strings.")
	comprehend_detectToxicContentCmd.MarkFlagRequired("language-code")
	comprehend_detectToxicContentCmd.MarkFlagRequired("text-segments")
	comprehendCmd.AddCommand(comprehend_detectToxicContentCmd)
}
