package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_untagResourceCmd).Standalone()

	internetmonitor_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a tag you remove a resource from.")
	internetmonitor_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys that you remove from a resource.")
	internetmonitor_untagResourceCmd.MarkFlagRequired("resource-arn")
	internetmonitor_untagResourceCmd.MarkFlagRequired("tag-keys")
	internetmonitorCmd.AddCommand(internetmonitor_untagResourceCmd)
}
