package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createEgressOnlyInternetGatewayCmd = &cobra.Command{
	Use:   "create-egress-only-internet-gateway",
	Short: "\\[IPv6 only] Creates an egress-only internet gateway for your VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createEgressOnlyInternetGatewayCmd).Standalone()

	ec2_createEgressOnlyInternetGatewayCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createEgressOnlyInternetGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createEgressOnlyInternetGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createEgressOnlyInternetGatewayCmd.Flags().String("tag-specifications", "", "The tags to assign to the egress-only internet gateway.")
	ec2_createEgressOnlyInternetGatewayCmd.Flags().String("vpc-id", "", "The ID of the VPC for which to create the egress-only internet gateway.")
	ec2_createEgressOnlyInternetGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2_createEgressOnlyInternetGatewayCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_createEgressOnlyInternetGatewayCmd)
}
