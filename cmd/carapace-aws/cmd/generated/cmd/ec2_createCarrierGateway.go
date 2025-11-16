package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCarrierGatewayCmd = &cobra.Command{
	Use:   "create-carrier-gateway",
	Short: "Creates a carrier gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCarrierGatewayCmd).Standalone()

	ec2_createCarrierGatewayCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createCarrierGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCarrierGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createCarrierGatewayCmd.Flags().String("tag-specifications", "", "The tags to associate with the carrier gateway.")
	ec2_createCarrierGatewayCmd.Flags().String("vpc-id", "", "The ID of the VPC to associate with the carrier gateway.")
	ec2_createCarrierGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2_createCarrierGatewayCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_createCarrierGatewayCmd)
}
