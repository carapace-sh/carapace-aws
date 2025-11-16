package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableVgwRoutePropagationCmd = &cobra.Command{
	Use:   "enable-vgw-route-propagation",
	Short: "Enables a virtual private gateway (VGW) to propagate routes to the specified route table of a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableVgwRoutePropagationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableVgwRoutePropagationCmd).Standalone()

		ec2_enableVgwRoutePropagationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableVgwRoutePropagationCmd.Flags().String("gateway-id", "", "The ID of the virtual private gateway that is attached to a VPC.")
		ec2_enableVgwRoutePropagationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableVgwRoutePropagationCmd.Flags().String("route-table-id", "", "The ID of the route table.")
		ec2_enableVgwRoutePropagationCmd.MarkFlagRequired("gateway-id")
		ec2_enableVgwRoutePropagationCmd.Flag("no-dry-run").Hidden = true
		ec2_enableVgwRoutePropagationCmd.MarkFlagRequired("route-table-id")
	})
	ec2Cmd.AddCommand(ec2_enableVgwRoutePropagationCmd)
}
