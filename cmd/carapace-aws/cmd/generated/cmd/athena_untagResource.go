package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from an Athena resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_untagResourceCmd).Standalone()

		athena_untagResourceCmd.Flags().String("resource-arn", "", "Specifies the ARN of the resource from which tags are to be removed.")
		athena_untagResourceCmd.Flags().String("tag-keys", "", "A comma-separated list of one or more tag keys whose tags are to be removed from the specified resource.")
		athena_untagResourceCmd.MarkFlagRequired("resource-arn")
		athena_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	athenaCmd.AddCommand(athena_untagResourceCmd)
}
