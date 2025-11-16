package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_startCommitmentPurchaseAnalysisCmd = &cobra.Command{
	Use:   "start-commitment-purchase-analysis",
	Short: "Specifies the parameters of a planned commitment purchase and starts the generation of the analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_startCommitmentPurchaseAnalysisCmd).Standalone()

	ce_startCommitmentPurchaseAnalysisCmd.Flags().String("commitment-purchase-analysis-configuration", "", "The configuration for the commitment purchase analysis.")
	ce_startCommitmentPurchaseAnalysisCmd.MarkFlagRequired("commitment-purchase-analysis-configuration")
	ceCmd.AddCommand(ce_startCommitmentPurchaseAnalysisCmd)
}
