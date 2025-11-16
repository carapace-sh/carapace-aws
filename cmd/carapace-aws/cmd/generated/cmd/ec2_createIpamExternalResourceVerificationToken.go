package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createIpamExternalResourceVerificationTokenCmd = &cobra.Command{
	Use:   "create-ipam-external-resource-verification-token",
	Short: "Create a verification token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createIpamExternalResourceVerificationTokenCmd).Standalone()

	ec2_createIpamExternalResourceVerificationTokenCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createIpamExternalResourceVerificationTokenCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createIpamExternalResourceVerificationTokenCmd.Flags().String("ipam-id", "", "The ID of the IPAM that will create the token.")
	ec2_createIpamExternalResourceVerificationTokenCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createIpamExternalResourceVerificationTokenCmd.Flags().String("tag-specifications", "", "Token tags.")
	ec2_createIpamExternalResourceVerificationTokenCmd.MarkFlagRequired("ipam-id")
	ec2_createIpamExternalResourceVerificationTokenCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createIpamExternalResourceVerificationTokenCmd)
}
