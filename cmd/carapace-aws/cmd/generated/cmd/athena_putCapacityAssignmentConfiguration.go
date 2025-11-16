package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_putCapacityAssignmentConfigurationCmd = &cobra.Command{
	Use:   "put-capacity-assignment-configuration",
	Short: "Puts a new capacity assignment configuration for a specified capacity reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_putCapacityAssignmentConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_putCapacityAssignmentConfigurationCmd).Standalone()

		athena_putCapacityAssignmentConfigurationCmd.Flags().String("capacity-assignments", "", "The list of assignments for the capacity assignment configuration.")
		athena_putCapacityAssignmentConfigurationCmd.Flags().String("capacity-reservation-name", "", "The name of the capacity reservation to put a capacity assignment configuration for.")
		athena_putCapacityAssignmentConfigurationCmd.MarkFlagRequired("capacity-assignments")
		athena_putCapacityAssignmentConfigurationCmd.MarkFlagRequired("capacity-reservation-name")
	})
	athenaCmd.AddCommand(athena_putCapacityAssignmentConfigurationCmd)
}
