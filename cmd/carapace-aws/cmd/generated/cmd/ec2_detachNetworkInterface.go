package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_detachNetworkInterfaceCmd = &cobra.Command{
	Use:   "detach-network-interface",
	Short: "Detaches a network interface from an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_detachNetworkInterfaceCmd).Standalone()

	ec2_detachNetworkInterfaceCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
	ec2_detachNetworkInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_detachNetworkInterfaceCmd.Flags().Bool("force", false, "Specifies whether to force a detachment.")
	ec2_detachNetworkInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_detachNetworkInterfaceCmd.Flags().Bool("no-force", false, "Specifies whether to force a detachment.")
	ec2_detachNetworkInterfaceCmd.MarkFlagRequired("attachment-id")
	ec2_detachNetworkInterfaceCmd.Flag("no-dry-run").Hidden = true
	ec2_detachNetworkInterfaceCmd.Flag("no-force").Hidden = true
	ec2Cmd.AddCommand(ec2_detachNetworkInterfaceCmd)
}
