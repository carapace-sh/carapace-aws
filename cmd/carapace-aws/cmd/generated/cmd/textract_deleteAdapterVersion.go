package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_deleteAdapterVersionCmd = &cobra.Command{
	Use:   "delete-adapter-version",
	Short: "Deletes an Amazon Textract adapter version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_deleteAdapterVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_deleteAdapterVersionCmd).Standalone()

		textract_deleteAdapterVersionCmd.Flags().String("adapter-id", "", "A string containing a unique ID for the adapter version that will be deleted.")
		textract_deleteAdapterVersionCmd.Flags().String("adapter-version", "", "Specifies the adapter version to be deleted.")
		textract_deleteAdapterVersionCmd.MarkFlagRequired("adapter-id")
		textract_deleteAdapterVersionCmd.MarkFlagRequired("adapter-version")
	})
	textractCmd.AddCommand(textract_deleteAdapterVersionCmd)
}
