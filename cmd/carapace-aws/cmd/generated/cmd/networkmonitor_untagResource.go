package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a key-value pair from a monitor or probe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_untagResourceCmd).Standalone()

	networkmonitor_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the monitor or probe that the tag should be removed from.")
	networkmonitor_untagResourceCmd.Flags().String("tag-keys", "", "The key-value pa")
	networkmonitor_untagResourceCmd.MarkFlagRequired("resource-arn")
	networkmonitor_untagResourceCmd.MarkFlagRequired("tag-keys")
	networkmonitorCmd.AddCommand(networkmonitor_untagResourceCmd)
}
