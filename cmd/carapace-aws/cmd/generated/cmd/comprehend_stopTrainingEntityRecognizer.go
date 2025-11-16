package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopTrainingEntityRecognizerCmd = &cobra.Command{
	Use:   "stop-training-entity-recognizer",
	Short: "Stops an entity recognizer training job while in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopTrainingEntityRecognizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_stopTrainingEntityRecognizerCmd).Standalone()

		comprehend_stopTrainingEntityRecognizerCmd.Flags().String("entity-recognizer-arn", "", "The Amazon Resource Name (ARN) that identifies the entity recognizer currently being trained.")
		comprehend_stopTrainingEntityRecognizerCmd.MarkFlagRequired("entity-recognizer-arn")
	})
	comprehendCmd.AddCommand(comprehend_stopTrainingEntityRecognizerCmd)
}
