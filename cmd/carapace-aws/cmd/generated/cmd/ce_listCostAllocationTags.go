package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_listCostAllocationTagsCmd = &cobra.Command{
	Use:   "list-cost-allocation-tags",
	Short: "Get a list of cost allocation tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_listCostAllocationTagsCmd).Standalone()

	ce_listCostAllocationTagsCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for this request.")
	ce_listCostAllocationTagsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	ce_listCostAllocationTagsCmd.Flags().String("status", "", "The status of cost allocation tag keys that are returned for this request.")
	ce_listCostAllocationTagsCmd.Flags().String("tag-keys", "", "The list of cost allocation tag keys that are returned for this request.")
	ce_listCostAllocationTagsCmd.Flags().String("type", "", "The type of `CostAllocationTag` object that are returned for this request.")
	ceCmd.AddCommand(ce_listCostAllocationTagsCmd)
}
