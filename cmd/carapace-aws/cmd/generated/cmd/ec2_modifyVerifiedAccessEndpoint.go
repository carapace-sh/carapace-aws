package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessEndpointCmd = &cobra.Command{
	Use:   "modify-verified-access-endpoint",
	Short: "Modifies the configuration of the specified Amazon Web Services Verified Access endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVerifiedAccessEndpointCmd).Standalone()

		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("cidr-options", "", "The CIDR options.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("description", "", "A description for the Verified Access endpoint.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("load-balancer-options", "", "The load balancer details if creating the Verified Access endpoint as `load-balancer`type.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("network-interface-options", "", "The network interface options.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("rds-options", "", "The RDS options.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("verified-access-endpoint-id", "", "The ID of the Verified Access endpoint.")
		ec2_modifyVerifiedAccessEndpointCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group.")
		ec2_modifyVerifiedAccessEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVerifiedAccessEndpointCmd.MarkFlagRequired("verified-access-endpoint-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessEndpointCmd)
}
