package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_batchDeleteReadSetCmd = &cobra.Command{
	Use:   "batch-delete-read-set",
	Short: "Deletes one or more read sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_batchDeleteReadSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_batchDeleteReadSetCmd).Standalone()

		omics_batchDeleteReadSetCmd.Flags().String("ids", "", "The read sets' IDs.")
		omics_batchDeleteReadSetCmd.Flags().String("sequence-store-id", "", "The read sets' sequence store ID.")
		omics_batchDeleteReadSetCmd.MarkFlagRequired("ids")
		omics_batchDeleteReadSetCmd.MarkFlagRequired("sequence-store-id")
	})
	omicsCmd.AddCommand(omics_batchDeleteReadSetCmd)
}
