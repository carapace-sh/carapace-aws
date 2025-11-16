package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_containsPiiEntitiesCmd = &cobra.Command{
	Use:   "contains-pii-entities",
	Short: "Analyzes input text for the presence of personally identifiable information (PII) and returns the labels of identified PII entity types such as name, address, bank account number, or phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_containsPiiEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_containsPiiEntitiesCmd).Standalone()

		comprehend_containsPiiEntitiesCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_containsPiiEntitiesCmd.Flags().String("text", "", "A UTF-8 text string.")
		comprehend_containsPiiEntitiesCmd.MarkFlagRequired("language-code")
		comprehend_containsPiiEntitiesCmd.MarkFlagRequired("text")
	})
	comprehendCmd.AddCommand(comprehend_containsPiiEntitiesCmd)
}
