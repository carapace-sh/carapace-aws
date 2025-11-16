package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamExternalResourceVerificationTokenCmd = &cobra.Command{
	Use:   "delete-ipam-external-resource-verification-token",
	Short: "Delete a verification token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamExternalResourceVerificationTokenCmd).Standalone()

	ec2_deleteIpamExternalResourceVerificationTokenCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteIpamExternalResourceVerificationTokenCmd.Flags().String("ipam-external-resource-verification-token-id", "", "The token ID.")
	ec2_deleteIpamExternalResourceVerificationTokenCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteIpamExternalResourceVerificationTokenCmd.MarkFlagRequired("ipam-external-resource-verification-token-id")
	ec2_deleteIpamExternalResourceVerificationTokenCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteIpamExternalResourceVerificationTokenCmd)
}
