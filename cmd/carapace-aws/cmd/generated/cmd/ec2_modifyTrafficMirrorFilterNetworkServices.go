package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyTrafficMirrorFilterNetworkServicesCmd = &cobra.Command{
	Use:   "modify-traffic-mirror-filter-network-services",
	Short: "Allows or restricts mirroring network services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyTrafficMirrorFilterNetworkServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyTrafficMirrorFilterNetworkServicesCmd).Standalone()

		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.Flags().String("add-network-services", "", "The network service, for example Amazon DNS, that you want to mirror.")
		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.Flags().String("remove-network-services", "", "The network service, for example Amazon DNS, that you no longer want to mirror.")
		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.Flags().String("traffic-mirror-filter-id", "", "The ID of the Traffic Mirror filter.")
		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyTrafficMirrorFilterNetworkServicesCmd.MarkFlagRequired("traffic-mirror-filter-id")
	})
	ec2Cmd.AddCommand(ec2_modifyTrafficMirrorFilterNetworkServicesCmd)
}
