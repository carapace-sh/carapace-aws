package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCommitmentPurchaseAnalysisCmd = &cobra.Command{
	Use:   "get-commitment-purchase-analysis",
	Short: "Retrieves a commitment purchase analysis result based on the `AnalysisId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCommitmentPurchaseAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getCommitmentPurchaseAnalysisCmd).Standalone()

		ce_getCommitmentPurchaseAnalysisCmd.Flags().String("analysis-id", "", "The analysis ID that's associated with the commitment purchase analysis.")
		ce_getCommitmentPurchaseAnalysisCmd.MarkFlagRequired("analysis-id")
	})
	ceCmd.AddCommand(ce_getCommitmentPurchaseAnalysisCmd)
}
