package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or set of tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_untagResourceCmd).Standalone()

		redshiftServerless_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
		redshiftServerless_untagResourceCmd.Flags().String("tag-keys", "", "The tag or set of tags to remove from the resource.")
		redshiftServerless_untagResourceCmd.MarkFlagRequired("resource-arn")
		redshiftServerless_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_untagResourceCmd)
}
