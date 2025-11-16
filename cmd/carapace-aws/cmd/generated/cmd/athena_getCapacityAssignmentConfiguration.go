package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getCapacityAssignmentConfigurationCmd = &cobra.Command{
	Use:   "get-capacity-assignment-configuration",
	Short: "Gets the capacity assignment configuration for a capacity reservation, if one exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getCapacityAssignmentConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getCapacityAssignmentConfigurationCmd).Standalone()

		athena_getCapacityAssignmentConfigurationCmd.Flags().String("capacity-reservation-name", "", "The name of the capacity reservation to retrieve the capacity assignment configuration for.")
		athena_getCapacityAssignmentConfigurationCmd.MarkFlagRequired("capacity-reservation-name")
	})
	athenaCmd.AddCommand(athena_getCapacityAssignmentConfigurationCmd)
}
