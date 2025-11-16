package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createMilestoneCmd = &cobra.Command{
	Use:   "create-milestone",
	Short: "Create a milestone for an existing workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createMilestoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_createMilestoneCmd).Standalone()

		wellarchitected_createMilestoneCmd.Flags().String("client-request-token", "", "")
		wellarchitected_createMilestoneCmd.Flags().String("milestone-name", "", "")
		wellarchitected_createMilestoneCmd.Flags().String("workload-id", "", "")
		wellarchitected_createMilestoneCmd.MarkFlagRequired("client-request-token")
		wellarchitected_createMilestoneCmd.MarkFlagRequired("milestone-name")
		wellarchitected_createMilestoneCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_createMilestoneCmd)
}
