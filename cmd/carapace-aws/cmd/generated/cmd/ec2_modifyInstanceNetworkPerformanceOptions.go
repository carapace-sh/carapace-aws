package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceNetworkPerformanceOptionsCmd = &cobra.Command{
	Use:   "modify-instance-network-performance-options",
	Short: "Change the configuration of the network performance options for an existing instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceNetworkPerformanceOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyInstanceNetworkPerformanceOptionsCmd).Standalone()

		ec2_modifyInstanceNetworkPerformanceOptionsCmd.Flags().String("bandwidth-weighting", "", "Specify the bandwidth weighting option to boost the associated type of baseline bandwidth, as follows:")
		ec2_modifyInstanceNetworkPerformanceOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyInstanceNetworkPerformanceOptionsCmd.Flags().String("instance-id", "", "The ID of the instance to update.")
		ec2_modifyInstanceNetworkPerformanceOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyInstanceNetworkPerformanceOptionsCmd.MarkFlagRequired("bandwidth-weighting")
		ec2_modifyInstanceNetworkPerformanceOptionsCmd.MarkFlagRequired("instance-id")
		ec2_modifyInstanceNetworkPerformanceOptionsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyInstanceNetworkPerformanceOptionsCmd)
}
