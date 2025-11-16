package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_tagResourceCmd).Standalone()

	internetmonitor_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a tag that you add to a resource.")
	internetmonitor_tagResourceCmd.Flags().String("tags", "", "Tags that you add to a resource.")
	internetmonitor_tagResourceCmd.MarkFlagRequired("resource-arn")
	internetmonitor_tagResourceCmd.MarkFlagRequired("tags")
	internetmonitorCmd.AddCommand(internetmonitor_tagResourceCmd)
}
