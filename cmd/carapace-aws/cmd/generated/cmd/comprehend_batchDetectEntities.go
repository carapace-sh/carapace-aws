package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_batchDetectEntitiesCmd = &cobra.Command{
	Use:   "batch-detect-entities",
	Short: "Inspects the text of a batch of documents for named entities and returns information about them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_batchDetectEntitiesCmd).Standalone()

	comprehend_batchDetectEntitiesCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehend_batchDetectEntitiesCmd.Flags().String("text-list", "", "A list containing the UTF-8 encoded text of the input documents.")
	comprehend_batchDetectEntitiesCmd.MarkFlagRequired("language-code")
	comprehend_batchDetectEntitiesCmd.MarkFlagRequired("text-list")
	comprehendCmd.AddCommand(comprehend_batchDetectEntitiesCmd)
}
