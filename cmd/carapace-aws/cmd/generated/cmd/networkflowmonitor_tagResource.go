package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_tagResourceCmd).Standalone()

	networkflowmonitor_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	networkflowmonitor_tagResourceCmd.Flags().String("tags", "", "The tags for a resource.")
	networkflowmonitor_tagResourceCmd.MarkFlagRequired("resource-arn")
	networkflowmonitor_tagResourceCmd.MarkFlagRequired("tags")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_tagResourceCmd)
}
