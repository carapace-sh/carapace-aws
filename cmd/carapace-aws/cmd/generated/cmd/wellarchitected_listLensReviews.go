package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listLensReviewsCmd = &cobra.Command{
	Use:   "list-lens-reviews",
	Short: "List lens reviews for a particular workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listLensReviewsCmd).Standalone()

	wellarchitected_listLensReviewsCmd.Flags().String("max-results", "", "")
	wellarchitected_listLensReviewsCmd.Flags().String("milestone-number", "", "")
	wellarchitected_listLensReviewsCmd.Flags().String("next-token", "", "")
	wellarchitected_listLensReviewsCmd.Flags().String("workload-id", "", "")
	wellarchitected_listLensReviewsCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_listLensReviewsCmd)
}
