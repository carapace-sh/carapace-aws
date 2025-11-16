package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeByoipCidrsCmd = &cobra.Command{
	Use:   "describe-byoip-cidrs",
	Short: "Describes the IP address ranges that were provisioned for use with Amazon Web Services resources through through bring your own IP addresses (BYOIP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeByoipCidrsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeByoipCidrsCmd).Standalone()

		ec2_describeByoipCidrsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeByoipCidrsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeByoipCidrsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeByoipCidrsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeByoipCidrsCmd.MarkFlagRequired("max-results")
		ec2_describeByoipCidrsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeByoipCidrsCmd)
}
