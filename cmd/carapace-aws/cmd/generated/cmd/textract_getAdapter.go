package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_getAdapterCmd = &cobra.Command{
	Use:   "get-adapter",
	Short: "Gets configuration information for an adapter specified by an AdapterId, returning information on AdapterName, Description, CreationTime, AutoUpdate status, and FeatureTypes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_getAdapterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_getAdapterCmd).Standalone()

		textract_getAdapterCmd.Flags().String("adapter-id", "", "A string containing a unique ID for the adapter.")
		textract_getAdapterCmd.MarkFlagRequired("adapter-id")
	})
	textractCmd.AddCommand(textract_getAdapterCmd)
}
