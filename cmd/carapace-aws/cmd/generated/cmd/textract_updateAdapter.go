package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_updateAdapterCmd = &cobra.Command{
	Use:   "update-adapter",
	Short: "Update the configuration for an adapter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_updateAdapterCmd).Standalone()

	textract_updateAdapterCmd.Flags().String("adapter-id", "", "A string containing a unique ID for the adapter that will be updated.")
	textract_updateAdapterCmd.Flags().String("adapter-name", "", "The new name to be applied to the adapter.")
	textract_updateAdapterCmd.Flags().String("auto-update", "", "The new auto-update status to be applied to the adapter.")
	textract_updateAdapterCmd.Flags().String("description", "", "The new description to be applied to the adapter.")
	textract_updateAdapterCmd.MarkFlagRequired("adapter-id")
	textractCmd.AddCommand(textract_updateAdapterCmd)
}
