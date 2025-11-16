package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_listTagsForResourceCmd).Standalone()

		neptuneGraph_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		neptuneGraph_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_listTagsForResourceCmd)
}
