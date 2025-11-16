package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateRouteTableCmd = &cobra.Command{
	Use:   "disassociate-route-table",
	Short: "Disassociates a subnet or gateway from a route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateRouteTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateRouteTableCmd).Standalone()

		ec2_disassociateRouteTableCmd.Flags().String("association-id", "", "The association ID representing the current association between the route table and subnet or gateway.")
		ec2_disassociateRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateRouteTableCmd.MarkFlagRequired("association-id")
		ec2_disassociateRouteTableCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disassociateRouteTableCmd)
}
