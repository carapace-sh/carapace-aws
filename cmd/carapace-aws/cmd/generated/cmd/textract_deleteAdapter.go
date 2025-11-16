package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_deleteAdapterCmd = &cobra.Command{
	Use:   "delete-adapter",
	Short: "Deletes an Amazon Textract adapter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_deleteAdapterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_deleteAdapterCmd).Standalone()

		textract_deleteAdapterCmd.Flags().String("adapter-id", "", "A string containing a unique ID for the adapter to be deleted.")
		textract_deleteAdapterCmd.MarkFlagRequired("adapter-id")
	})
	textractCmd.AddCommand(textract_deleteAdapterCmd)
}
