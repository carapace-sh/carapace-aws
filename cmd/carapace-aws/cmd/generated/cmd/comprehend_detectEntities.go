package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectEntitiesCmd = &cobra.Command{
	Use:   "detect-entities",
	Short: "Detects named entities in input text when you use the pre-trained model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_detectEntitiesCmd).Standalone()

		comprehend_detectEntitiesCmd.Flags().String("bytes", "", "This field applies only when you use a custom entity recognition model that was trained with PDF annotations.")
		comprehend_detectEntitiesCmd.Flags().String("document-reader-config", "", "Provides configuration parameters to override the default actions for extracting text from PDF documents and image files.")
		comprehend_detectEntitiesCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name of an endpoint that is associated with a custom entity recognition model.")
		comprehend_detectEntitiesCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_detectEntitiesCmd.Flags().String("text", "", "A UTF-8 text string.")
	})
	comprehendCmd.AddCommand(comprehend_detectEntitiesCmd)
}
