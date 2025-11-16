package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_createAdapterCmd = &cobra.Command{
	Use:   "create-adapter",
	Short: "Creates an adapter, which can be fine-tuned for enhanced performance on user provided documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_createAdapterCmd).Standalone()

	textract_createAdapterCmd.Flags().String("adapter-name", "", "The name to be assigned to the adapter being created.")
	textract_createAdapterCmd.Flags().String("auto-update", "", "Controls whether or not the adapter should automatically update.")
	textract_createAdapterCmd.Flags().String("client-request-token", "", "Idempotent token is used to recognize the request.")
	textract_createAdapterCmd.Flags().String("description", "", "The description to be assigned to the adapter being created.")
	textract_createAdapterCmd.Flags().String("feature-types", "", "The type of feature that the adapter is being trained on.")
	textract_createAdapterCmd.Flags().String("tags", "", "A list of tags to be added to the adapter.")
	textract_createAdapterCmd.MarkFlagRequired("adapter-name")
	textract_createAdapterCmd.MarkFlagRequired("feature-types")
	textractCmd.AddCommand(textract_createAdapterCmd)
}
