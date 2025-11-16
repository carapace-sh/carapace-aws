package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_untagResourceCmd).Standalone()

	resiliencehub_untagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource.")
	resiliencehub_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags you want to remove.")
	resiliencehub_untagResourceCmd.MarkFlagRequired("resource-arn")
	resiliencehub_untagResourceCmd.MarkFlagRequired("tag-keys")
	resiliencehubCmd.AddCommand(resiliencehub_untagResourceCmd)
}
