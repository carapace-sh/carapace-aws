package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listReviewTemplatesCmd = &cobra.Command{
	Use:   "list-review-templates",
	Short: "List review templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listReviewTemplatesCmd).Standalone()

	wellarchitected_listReviewTemplatesCmd.Flags().String("max-results", "", "")
	wellarchitected_listReviewTemplatesCmd.Flags().String("next-token", "", "")
	wellarchitectedCmd.AddCommand(wellarchitected_listReviewTemplatesCmd)
}
