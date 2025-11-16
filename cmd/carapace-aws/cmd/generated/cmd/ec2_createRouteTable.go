package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createRouteTableCmd = &cobra.Command{
	Use:   "create-route-table",
	Short: "Creates a route table for the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createRouteTableCmd).Standalone()

	ec2_createRouteTableCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createRouteTableCmd.Flags().String("tag-specifications", "", "The tags to assign to the route table.")
	ec2_createRouteTableCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_createRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2_createRouteTableCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_createRouteTableCmd)
}
