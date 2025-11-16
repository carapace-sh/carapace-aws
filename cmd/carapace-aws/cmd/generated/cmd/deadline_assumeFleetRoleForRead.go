package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_assumeFleetRoleForReadCmd = &cobra.Command{
	Use:   "assume-fleet-role-for-read",
	Short: "Get Amazon Web Services credentials from the fleet role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_assumeFleetRoleForReadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_assumeFleetRoleForReadCmd).Standalone()

		deadline_assumeFleetRoleForReadCmd.Flags().String("farm-id", "", "The farm ID for the fleet's farm.")
		deadline_assumeFleetRoleForReadCmd.Flags().String("fleet-id", "", "The fleet ID.")
		deadline_assumeFleetRoleForReadCmd.MarkFlagRequired("farm-id")
		deadline_assumeFleetRoleForReadCmd.MarkFlagRequired("fleet-id")
	})
	deadlineCmd.AddCommand(deadline_assumeFleetRoleForReadCmd)
}
