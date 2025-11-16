package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectPiiEntitiesCmd = &cobra.Command{
	Use:   "detect-pii-entities",
	Short: "Inspects the input text for entities that contain personally identifiable information (PII) and returns information about them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectPiiEntitiesCmd).Standalone()

	comprehend_detectPiiEntitiesCmd.Flags().String("language-code", "", "The language of the input text.")
	comprehend_detectPiiEntitiesCmd.Flags().String("text", "", "A UTF-8 text string.")
	comprehend_detectPiiEntitiesCmd.MarkFlagRequired("language-code")
	comprehend_detectPiiEntitiesCmd.MarkFlagRequired("text")
	comprehendCmd.AddCommand(comprehend_detectPiiEntitiesCmd)
}
