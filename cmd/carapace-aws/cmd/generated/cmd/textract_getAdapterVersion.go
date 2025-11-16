package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_getAdapterVersionCmd = &cobra.Command{
	Use:   "get-adapter-version",
	Short: "Gets configuration information for the specified adapter version, including: AdapterId, AdapterVersion, FeatureTypes, Status, StatusMessage, DatasetConfig, KMSKeyId, OutputConfig, Tags and EvaluationMetrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_getAdapterVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_getAdapterVersionCmd).Standalone()

		textract_getAdapterVersionCmd.Flags().String("adapter-id", "", "A string specifying a unique ID for the adapter version you want to retrieve information for.")
		textract_getAdapterVersionCmd.Flags().String("adapter-version", "", "A string specifying the adapter version you want to retrieve information for.")
		textract_getAdapterVersionCmd.MarkFlagRequired("adapter-id")
		textract_getAdapterVersionCmd.MarkFlagRequired("adapter-version")
	})
	textractCmd.AddCommand(textract_getAdapterVersionCmd)
}
