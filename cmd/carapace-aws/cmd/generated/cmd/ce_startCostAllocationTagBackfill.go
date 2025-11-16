package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_startCostAllocationTagBackfillCmd = &cobra.Command{
	Use:   "start-cost-allocation-tag-backfill",
	Short: "Request a cost allocation tag backfill.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_startCostAllocationTagBackfillCmd).Standalone()

	ce_startCostAllocationTagBackfillCmd.Flags().String("backfill-from", "", "The date you want the backfill to start from.")
	ce_startCostAllocationTagBackfillCmd.MarkFlagRequired("backfill-from")
	ceCmd.AddCommand(ce_startCostAllocationTagBackfillCmd)
}
