package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds key-value pairs to a monitor or probe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_tagResourceCmd).Standalone()

		networkmonitor_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the monitor or probe to tag.")
		networkmonitor_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs assigned to the monitor or probe.")
		networkmonitor_tagResourceCmd.MarkFlagRequired("resource-arn")
		networkmonitor_tagResourceCmd.MarkFlagRequired("tags")
	})
	networkmonitorCmd.AddCommand(networkmonitor_tagResourceCmd)
}
