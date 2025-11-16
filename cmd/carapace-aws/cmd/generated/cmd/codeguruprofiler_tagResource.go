package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Use to assign one or more tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_tagResourceCmd).Standalone()

	codeguruprofiler_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that the tags are added to.")
	codeguruprofiler_tagResourceCmd.Flags().String("tags", "", "The list of tags that are added to the specified resource.")
	codeguruprofiler_tagResourceCmd.MarkFlagRequired("resource-arn")
	codeguruprofiler_tagResourceCmd.MarkFlagRequired("tags")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_tagResourceCmd)
}
