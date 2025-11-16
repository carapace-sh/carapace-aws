package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns all the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_listTagsForResourceCmd).Standalone()

	networkflowmonitor_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	networkflowmonitor_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_listTagsForResourceCmd)
}
