package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listLensReviewImprovementsCmd = &cobra.Command{
	Use:   "list-lens-review-improvements",
	Short: "List the improvements of a particular lens review.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listLensReviewImprovementsCmd).Standalone()

	wellarchitected_listLensReviewImprovementsCmd.Flags().String("lens-alias", "", "")
	wellarchitected_listLensReviewImprovementsCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
	wellarchitected_listLensReviewImprovementsCmd.Flags().String("milestone-number", "", "")
	wellarchitected_listLensReviewImprovementsCmd.Flags().String("next-token", "", "")
	wellarchitected_listLensReviewImprovementsCmd.Flags().String("pillar-id", "", "")
	wellarchitected_listLensReviewImprovementsCmd.Flags().String("question-priority", "", "The priority of the question.")
	wellarchitected_listLensReviewImprovementsCmd.Flags().String("workload-id", "", "")
	wellarchitected_listLensReviewImprovementsCmd.MarkFlagRequired("lens-alias")
	wellarchitected_listLensReviewImprovementsCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_listLensReviewImprovementsCmd)
}
