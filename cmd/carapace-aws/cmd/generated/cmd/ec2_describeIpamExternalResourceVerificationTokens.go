package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamExternalResourceVerificationTokensCmd = &cobra.Command{
	Use:   "describe-ipam-external-resource-verification-tokens",
	Short: "Describe verification tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamExternalResourceVerificationTokensCmd).Standalone()

	ec2_describeIpamExternalResourceVerificationTokensCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeIpamExternalResourceVerificationTokensCmd.Flags().String("filters", "", "One or more filters for the request.")
	ec2_describeIpamExternalResourceVerificationTokensCmd.Flags().String("ipam-external-resource-verification-token-ids", "", "Verification token IDs.")
	ec2_describeIpamExternalResourceVerificationTokensCmd.Flags().String("max-results", "", "The maximum number of tokens to return in one page of results.")
	ec2_describeIpamExternalResourceVerificationTokensCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeIpamExternalResourceVerificationTokensCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeIpamExternalResourceVerificationTokensCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeIpamExternalResourceVerificationTokensCmd)
}
