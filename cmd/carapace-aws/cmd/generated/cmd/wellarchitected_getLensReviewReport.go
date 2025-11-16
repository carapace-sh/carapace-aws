package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getLensReviewReportCmd = &cobra.Command{
	Use:   "get-lens-review-report",
	Short: "Get lens review report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getLensReviewReportCmd).Standalone()

	wellarchitected_getLensReviewReportCmd.Flags().String("lens-alias", "", "")
	wellarchitected_getLensReviewReportCmd.Flags().String("milestone-number", "", "")
	wellarchitected_getLensReviewReportCmd.Flags().String("workload-id", "", "")
	wellarchitected_getLensReviewReportCmd.MarkFlagRequired("lens-alias")
	wellarchitected_getLensReviewReportCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_getLensReviewReportCmd)
}
