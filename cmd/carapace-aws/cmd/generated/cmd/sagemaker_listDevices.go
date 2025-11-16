package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "A list of devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listDevicesCmd).Standalone()

		sagemaker_listDevicesCmd.Flags().String("device-fleet-name", "", "Filter for fleets containing this name in their device fleet name.")
		sagemaker_listDevicesCmd.Flags().String("latest-heartbeat-after", "", "Select fleets where the job was updated after X")
		sagemaker_listDevicesCmd.Flags().String("max-results", "", "Maximum number of results to select.")
		sagemaker_listDevicesCmd.Flags().String("model-name", "", "A filter that searches devices that contains this name in any of their models.")
		sagemaker_listDevicesCmd.Flags().String("next-token", "", "The response from the last list when returning a list large enough to need tokening.")
	})
	sagemakerCmd.AddCommand(sagemaker_listDevicesCmd)
}
