package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_upgradeLensReviewCmd = &cobra.Command{
	Use:   "upgrade-lens-review",
	Short: "Upgrade lens review for a particular workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_upgradeLensReviewCmd).Standalone()

	wellarchitected_upgradeLensReviewCmd.Flags().String("client-request-token", "", "")
	wellarchitected_upgradeLensReviewCmd.Flags().String("lens-alias", "", "")
	wellarchitected_upgradeLensReviewCmd.Flags().String("milestone-name", "", "")
	wellarchitected_upgradeLensReviewCmd.Flags().String("workload-id", "", "")
	wellarchitected_upgradeLensReviewCmd.MarkFlagRequired("lens-alias")
	wellarchitected_upgradeLensReviewCmd.MarkFlagRequired("milestone-name")
	wellarchitected_upgradeLensReviewCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_upgradeLensReviewCmd)
}
