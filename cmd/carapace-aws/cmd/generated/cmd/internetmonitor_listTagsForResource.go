package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_listTagsForResourceCmd).Standalone()

		internetmonitor_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a resource.")
		internetmonitor_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	internetmonitorCmd.AddCommand(internetmonitor_listTagsForResourceCmd)
}
