package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeMacHostsCmd = &cobra.Command{
	Use:   "describe-mac-hosts",
	Short: "Describes the specified EC2 Mac Dedicated Host or all of your EC2 Mac Dedicated Hosts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeMacHostsCmd).Standalone()

	ec2_describeMacHostsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeMacHostsCmd.Flags().String("host-ids", "", "The IDs of the EC2 Mac Dedicated Hosts.")
	ec2_describeMacHostsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeMacHostsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ec2Cmd.AddCommand(ec2_describeMacHostsCmd)
}
