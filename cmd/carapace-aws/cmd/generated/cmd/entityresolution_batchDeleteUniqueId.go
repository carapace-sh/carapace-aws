package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_batchDeleteUniqueIdCmd = &cobra.Command{
	Use:   "batch-delete-unique-id",
	Short: "Deletes multiple unique IDs in a matching workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_batchDeleteUniqueIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_batchDeleteUniqueIdCmd).Standalone()

		entityresolution_batchDeleteUniqueIdCmd.Flags().String("input-source", "", "The input source for the batch delete unique ID operation.")
		entityresolution_batchDeleteUniqueIdCmd.Flags().String("unique-ids", "", "The unique IDs to delete.")
		entityresolution_batchDeleteUniqueIdCmd.Flags().String("workflow-name", "", "The name of the workflow.")
		entityresolution_batchDeleteUniqueIdCmd.MarkFlagRequired("unique-ids")
		entityresolution_batchDeleteUniqueIdCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_batchDeleteUniqueIdCmd)
}
