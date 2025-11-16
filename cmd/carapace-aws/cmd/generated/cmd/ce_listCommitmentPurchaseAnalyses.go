package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_listCommitmentPurchaseAnalysesCmd = &cobra.Command{
	Use:   "list-commitment-purchase-analyses",
	Short: "Lists the commitment purchase analyses for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_listCommitmentPurchaseAnalysesCmd).Standalone()

	ce_listCommitmentPurchaseAnalysesCmd.Flags().String("analysis-ids", "", "The analysis IDs associated with the commitment purchase analyses.")
	ce_listCommitmentPurchaseAnalysesCmd.Flags().String("analysis-status", "", "The status of the analysis.")
	ce_listCommitmentPurchaseAnalysesCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_listCommitmentPurchaseAnalysesCmd.Flags().String("page-size", "", "The number of analyses that you want returned in a single response object.")
	ceCmd.AddCommand(ce_listCommitmentPurchaseAnalysesCmd)
}
