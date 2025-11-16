package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateLensReviewCmd = &cobra.Command{
	Use:   "update-lens-review",
	Short: "Update lens review for a particular workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateLensReviewCmd).Standalone()

	wellarchitected_updateLensReviewCmd.Flags().String("jira-configuration", "", "Configuration of the Jira integration.")
	wellarchitected_updateLensReviewCmd.Flags().String("lens-alias", "", "")
	wellarchitected_updateLensReviewCmd.Flags().String("lens-notes", "", "")
	wellarchitected_updateLensReviewCmd.Flags().String("pillar-notes", "", "")
	wellarchitected_updateLensReviewCmd.Flags().String("workload-id", "", "")
	wellarchitected_updateLensReviewCmd.MarkFlagRequired("lens-alias")
	wellarchitected_updateLensReviewCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_updateLensReviewCmd)
}
