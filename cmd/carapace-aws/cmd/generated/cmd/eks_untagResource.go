package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from an Amazon EKS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_untagResourceCmd).Standalone()

		eks_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to delete tags from.")
		eks_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove.")
		eks_untagResourceCmd.MarkFlagRequired("resource-arn")
		eks_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	eksCmd.AddCommand(eks_untagResourceCmd)
}
