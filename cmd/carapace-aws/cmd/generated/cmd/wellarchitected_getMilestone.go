package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getMilestoneCmd = &cobra.Command{
	Use:   "get-milestone",
	Short: "Get a milestone for an existing workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getMilestoneCmd).Standalone()

	wellarchitected_getMilestoneCmd.Flags().String("milestone-number", "", "")
	wellarchitected_getMilestoneCmd.Flags().String("workload-id", "", "")
	wellarchitected_getMilestoneCmd.MarkFlagRequired("milestone-number")
	wellarchitected_getMilestoneCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_getMilestoneCmd)
}
