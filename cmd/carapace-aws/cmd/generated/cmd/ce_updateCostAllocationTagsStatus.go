package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_updateCostAllocationTagsStatusCmd = &cobra.Command{
	Use:   "update-cost-allocation-tags-status",
	Short: "Updates status for cost allocation tags in bulk, with maximum batch size of 20. If the tag status that's updated is the same as the existing tag status, the request doesn't fail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_updateCostAllocationTagsStatusCmd).Standalone()

	ce_updateCostAllocationTagsStatusCmd.Flags().String("cost-allocation-tags-status", "", "The list of `CostAllocationTagStatusEntry` objects that are used to update cost allocation tags status for this request.")
	ce_updateCostAllocationTagsStatusCmd.MarkFlagRequired("cost-allocation-tags-status")
	ceCmd.AddCommand(ce_updateCostAllocationTagsStatusCmd)
}
