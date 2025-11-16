package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableVgwRoutePropagationCmd = &cobra.Command{
	Use:   "disable-vgw-route-propagation",
	Short: "Disables a virtual private gateway (VGW) from propagating routes to a specified route table of a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableVgwRoutePropagationCmd).Standalone()

	ec2_disableVgwRoutePropagationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableVgwRoutePropagationCmd.Flags().String("gateway-id", "", "The ID of the virtual private gateway.")
	ec2_disableVgwRoutePropagationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableVgwRoutePropagationCmd.Flags().String("route-table-id", "", "The ID of the route table.")
	ec2_disableVgwRoutePropagationCmd.MarkFlagRequired("gateway-id")
	ec2_disableVgwRoutePropagationCmd.Flag("no-dry-run").Hidden = true
	ec2_disableVgwRoutePropagationCmd.MarkFlagRequired("route-table-id")
	ec2Cmd.AddCommand(ec2_disableVgwRoutePropagationCmd)
}
