package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_untagResourceCmd).Standalone()

		networkflowmonitor_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		networkflowmonitor_untagResourceCmd.Flags().String("tag-keys", "", "Keys that you specified when you tagged a resource.")
		networkflowmonitor_untagResourceCmd.MarkFlagRequired("resource-arn")
		networkflowmonitor_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_untagResourceCmd)
}
