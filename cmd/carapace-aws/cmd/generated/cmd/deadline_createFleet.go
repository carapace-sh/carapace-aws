package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createFleetCmd = &cobra.Command{
	Use:   "create-fleet",
	Short: "Creates a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createFleetCmd).Standalone()

		deadline_createFleetCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_createFleetCmd.Flags().String("configuration", "", "The configuration settings for the fleet.")
		deadline_createFleetCmd.Flags().String("description", "", "The description of the fleet.")
		deadline_createFleetCmd.Flags().String("display-name", "", "The display name of the fleet.")
		deadline_createFleetCmd.Flags().String("farm-id", "", "The farm ID of the farm to connect to the fleet.")
		deadline_createFleetCmd.Flags().String("host-configuration", "", "Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.")
		deadline_createFleetCmd.Flags().String("max-worker-count", "", "The maximum number of workers for the fleet.")
		deadline_createFleetCmd.Flags().String("min-worker-count", "", "The minimum number of workers for the fleet.")
		deadline_createFleetCmd.Flags().String("role-arn", "", "The IAM role ARN for the role that the fleet's workers will use.")
		deadline_createFleetCmd.Flags().String("tags", "", "Each tag consists of a tag key and a tag value.")
		deadline_createFleetCmd.MarkFlagRequired("configuration")
		deadline_createFleetCmd.MarkFlagRequired("display-name")
		deadline_createFleetCmd.MarkFlagRequired("farm-id")
		deadline_createFleetCmd.MarkFlagRequired("max-worker-count")
		deadline_createFleetCmd.MarkFlagRequired("role-arn")
	})
	deadlineCmd.AddCommand(deadline_createFleetCmd)
}
