package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_batchPutPropertyValuesCmd = &cobra.Command{
	Use:   "batch-put-property-values",
	Short: "Sets values for multiple time series properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_batchPutPropertyValuesCmd).Standalone()

	iottwinmaker_batchPutPropertyValuesCmd.Flags().String("entries", "", "An object that maps strings to the property value entries to set.")
	iottwinmaker_batchPutPropertyValuesCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the properties to set.")
	iottwinmaker_batchPutPropertyValuesCmd.MarkFlagRequired("entries")
	iottwinmaker_batchPutPropertyValuesCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_batchPutPropertyValuesCmd)
}
