package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVerifiedAccessEndpointCmd = &cobra.Command{
	Use:   "delete-verified-access-endpoint",
	Short: "Delete an Amazon Web Services Verified Access endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVerifiedAccessEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVerifiedAccessEndpointCmd).Standalone()

		ec2_deleteVerifiedAccessEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_deleteVerifiedAccessEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVerifiedAccessEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVerifiedAccessEndpointCmd.Flags().String("verified-access-endpoint-id", "", "The ID of the Verified Access endpoint.")
		ec2_deleteVerifiedAccessEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteVerifiedAccessEndpointCmd.MarkFlagRequired("verified-access-endpoint-id")
	})
	ec2Cmd.AddCommand(ec2_deleteVerifiedAccessEndpointCmd)
}
