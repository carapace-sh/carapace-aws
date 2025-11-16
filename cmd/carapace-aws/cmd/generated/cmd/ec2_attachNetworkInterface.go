package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_attachNetworkInterfaceCmd = &cobra.Command{
	Use:   "attach-network-interface",
	Short: "Attaches a network interface to an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_attachNetworkInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_attachNetworkInterfaceCmd).Standalone()

		ec2_attachNetworkInterfaceCmd.Flags().String("device-index", "", "The index of the device for the network interface attachment.")
		ec2_attachNetworkInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_attachNetworkInterfaceCmd.Flags().String("ena-queue-count", "", "The number of ENA queues to be created with the instance.")
		ec2_attachNetworkInterfaceCmd.Flags().String("ena-srd-specification", "", "Configures ENA Express for the network interface that this action attaches to the instance.")
		ec2_attachNetworkInterfaceCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_attachNetworkInterfaceCmd.Flags().String("network-card-index", "", "The index of the network card.")
		ec2_attachNetworkInterfaceCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
		ec2_attachNetworkInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_attachNetworkInterfaceCmd.MarkFlagRequired("device-index")
		ec2_attachNetworkInterfaceCmd.MarkFlagRequired("instance-id")
		ec2_attachNetworkInterfaceCmd.MarkFlagRequired("network-interface-id")
		ec2_attachNetworkInterfaceCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_attachNetworkInterfaceCmd)
}
