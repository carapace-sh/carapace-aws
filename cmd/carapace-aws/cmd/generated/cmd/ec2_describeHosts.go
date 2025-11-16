package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeHostsCmd = &cobra.Command{
	Use:   "describe-hosts",
	Short: "Describes the specified Dedicated Hosts or all your Dedicated Hosts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeHostsCmd).Standalone()

		ec2_describeHostsCmd.Flags().String("filter", "", "The filters.")
		ec2_describeHostsCmd.Flags().String("host-ids", "", "The IDs of the Dedicated Hosts.")
		ec2_describeHostsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
		ec2_describeHostsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	})
	ec2Cmd.AddCommand(ec2_describeHostsCmd)
}
