package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_deleteEntityRecognizerCmd = &cobra.Command{
	Use:   "delete-entity-recognizer",
	Short: "Deletes an entity recognizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_deleteEntityRecognizerCmd).Standalone()

	comprehend_deleteEntityRecognizerCmd.Flags().String("entity-recognizer-arn", "", "The Amazon Resource Name (ARN) that identifies the entity recognizer.")
	comprehend_deleteEntityRecognizerCmd.MarkFlagRequired("entity-recognizer-arn")
	comprehendCmd.AddCommand(comprehend_deleteEntityRecognizerCmd)
}
