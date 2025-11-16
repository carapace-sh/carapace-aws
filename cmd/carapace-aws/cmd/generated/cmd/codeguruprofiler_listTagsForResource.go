package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags that are assigned to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_listTagsForResourceCmd).Standalone()

	codeguruprofiler_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that contains the tags to return.")
	codeguruprofiler_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_listTagsForResourceCmd)
}
