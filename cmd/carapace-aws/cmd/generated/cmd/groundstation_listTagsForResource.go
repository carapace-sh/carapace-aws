package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_listTagsForResourceCmd).Standalone()

		groundstation_listTagsForResourceCmd.Flags().String("resource-arn", "", "ARN of a resource.")
		groundstation_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	groundstationCmd.AddCommand(groundstation_listTagsForResourceCmd)
}
