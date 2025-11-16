package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getLensReviewCmd = &cobra.Command{
	Use:   "get-lens-review",
	Short: "Get lens review.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getLensReviewCmd).Standalone()

	wellarchitected_getLensReviewCmd.Flags().String("lens-alias", "", "")
	wellarchitected_getLensReviewCmd.Flags().String("milestone-number", "", "")
	wellarchitected_getLensReviewCmd.Flags().String("workload-id", "", "")
	wellarchitected_getLensReviewCmd.MarkFlagRequired("lens-alias")
	wellarchitected_getLensReviewCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_getLensReviewCmd)
}
