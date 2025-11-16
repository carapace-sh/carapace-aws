package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamsCmd = &cobra.Command{
	Use:   "describe-ipams",
	Short: "Get information about your IPAM pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamsCmd).Standalone()

	ec2_describeIpamsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeIpamsCmd.Flags().String("filters", "", "One or more filters for the request.")
	ec2_describeIpamsCmd.Flags().String("ipam-ids", "", "The IDs of the IPAMs you want information on.")
	ec2_describeIpamsCmd.Flags().String("max-results", "", "The maximum number of results to return in the request.")
	ec2_describeIpamsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeIpamsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_describeIpamsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeIpamsCmd)
}
