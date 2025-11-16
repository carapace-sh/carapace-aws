package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listDeviceFleetsCmd = &cobra.Command{
	Use:   "list-device-fleets",
	Short: "Returns a list of devices in the fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listDeviceFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listDeviceFleetsCmd).Standalone()

		sagemaker_listDeviceFleetsCmd.Flags().String("creation-time-after", "", "Filter fleets where packaging job was created after specified time.")
		sagemaker_listDeviceFleetsCmd.Flags().String("creation-time-before", "", "Filter fleets where the edge packaging job was created before specified time.")
		sagemaker_listDeviceFleetsCmd.Flags().String("last-modified-time-after", "", "Select fleets where the job was updated after X")
		sagemaker_listDeviceFleetsCmd.Flags().String("last-modified-time-before", "", "Select fleets where the job was updated before X")
		sagemaker_listDeviceFleetsCmd.Flags().String("max-results", "", "The maximum number of results to select.")
		sagemaker_listDeviceFleetsCmd.Flags().String("name-contains", "", "Filter for fleets containing this name in their fleet device name.")
		sagemaker_listDeviceFleetsCmd.Flags().String("next-token", "", "The response from the last list when returning a list large enough to need tokening.")
		sagemaker_listDeviceFleetsCmd.Flags().String("sort-by", "", "The column to sort by.")
		sagemaker_listDeviceFleetsCmd.Flags().String("sort-order", "", "What direction to sort in.")
	})
	sagemakerCmd.AddCommand(sagemaker_listDeviceFleetsCmd)
}
