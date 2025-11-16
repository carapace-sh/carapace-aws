package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to this resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_listTagsForResourceCmd).Standalone()

		networkmonitor_listTagsForResourceCmd.Flags().String("resource-arn", "", "The")
		networkmonitor_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	networkmonitorCmd.AddCommand(networkmonitor_listTagsForResourceCmd)
}
