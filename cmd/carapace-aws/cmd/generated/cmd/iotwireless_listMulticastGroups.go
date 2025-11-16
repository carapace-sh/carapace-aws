package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listMulticastGroupsCmd = &cobra.Command{
	Use:   "list-multicast-groups",
	Short: "Lists the multicast groups registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listMulticastGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listMulticastGroupsCmd).Standalone()

		iotwireless_listMulticastGroupsCmd.Flags().String("max-results", "", "")
		iotwireless_listMulticastGroupsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listMulticastGroupsCmd)
}
