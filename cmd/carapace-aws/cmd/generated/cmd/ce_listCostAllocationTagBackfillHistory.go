package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_listCostAllocationTagBackfillHistoryCmd = &cobra.Command{
	Use:   "list-cost-allocation-tag-backfill-history",
	Short: "Retrieves a list of your historical cost allocation tag backfill requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_listCostAllocationTagBackfillHistoryCmd).Standalone()

	ce_listCostAllocationTagBackfillHistoryCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for this request.")
	ce_listCostAllocationTagBackfillHistoryCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	ceCmd.AddCommand(ce_listCostAllocationTagBackfillHistoryCmd)
}
