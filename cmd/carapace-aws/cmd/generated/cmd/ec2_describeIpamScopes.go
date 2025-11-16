package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamScopesCmd = &cobra.Command{
	Use:   "describe-ipam-scopes",
	Short: "Get information about your IPAM scopes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamScopesCmd).Standalone()

	ec2_describeIpamScopesCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeIpamScopesCmd.Flags().String("filters", "", "One or more filters for the request.")
	ec2_describeIpamScopesCmd.Flags().String("ipam-scope-ids", "", "The IDs of the scopes you want information on.")
	ec2_describeIpamScopesCmd.Flags().String("max-results", "", "The maximum number of results to return in the request.")
	ec2_describeIpamScopesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeIpamScopesCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeIpamScopesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeIpamScopesCmd)
}
