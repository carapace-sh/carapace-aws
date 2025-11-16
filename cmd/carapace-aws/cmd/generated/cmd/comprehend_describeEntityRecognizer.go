package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeEntityRecognizerCmd = &cobra.Command{
	Use:   "describe-entity-recognizer",
	Short: "Provides details about an entity recognizer including status, S3 buckets containing training data, recognizer metadata, metrics, and so on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeEntityRecognizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeEntityRecognizerCmd).Standalone()

		comprehend_describeEntityRecognizerCmd.Flags().String("entity-recognizer-arn", "", "The Amazon Resource Name (ARN) that identifies the entity recognizer.")
		comprehend_describeEntityRecognizerCmd.MarkFlagRequired("entity-recognizer-arn")
	})
	comprehendCmd.AddCommand(comprehend_describeEntityRecognizerCmd)
}
