package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateFleetCmd = &cobra.Command{
	Use:   "update-fleet",
	Short: "Updates a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateFleetCmd).Standalone()

	deadline_updateFleetCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_updateFleetCmd.Flags().String("configuration", "", "The fleet configuration to update.")
	deadline_updateFleetCmd.Flags().String("description", "", "The description of the fleet to update.")
	deadline_updateFleetCmd.Flags().String("display-name", "", "The display name of the fleet to update.")
	deadline_updateFleetCmd.Flags().String("farm-id", "", "The farm ID to update.")
	deadline_updateFleetCmd.Flags().String("fleet-id", "", "The fleet ID to update.")
	deadline_updateFleetCmd.Flags().String("host-configuration", "", "Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.")
	deadline_updateFleetCmd.Flags().String("max-worker-count", "", "The maximum number of workers in the fleet.")
	deadline_updateFleetCmd.Flags().String("min-worker-count", "", "The minimum number of workers in the fleet.")
	deadline_updateFleetCmd.Flags().String("role-arn", "", "The IAM role ARN that the fleet's workers assume while running jobs.")
	deadline_updateFleetCmd.MarkFlagRequired("farm-id")
	deadline_updateFleetCmd.MarkFlagRequired("fleet-id")
	deadlineCmd.AddCommand(deadline_updateFleetCmd)
}
