package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Use to remove one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_untagResourceCmd).Standalone()

		codeguruprofiler_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that contains the tags to remove.")
		codeguruprofiler_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys.")
		codeguruprofiler_untagResourceCmd.MarkFlagRequired("resource-arn")
		codeguruprofiler_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_untagResourceCmd)
}
