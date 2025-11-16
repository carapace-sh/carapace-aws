package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listMilestonesCmd = &cobra.Command{
	Use:   "list-milestones",
	Short: "List all milestones for an existing workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listMilestonesCmd).Standalone()

	wellarchitected_listMilestonesCmd.Flags().String("max-results", "", "")
	wellarchitected_listMilestonesCmd.Flags().String("next-token", "", "")
	wellarchitected_listMilestonesCmd.Flags().String("workload-id", "", "")
	wellarchitected_listMilestonesCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_listMilestonesCmd)
}
