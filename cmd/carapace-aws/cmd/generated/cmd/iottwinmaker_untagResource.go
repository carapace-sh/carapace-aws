package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_untagResourceCmd).Standalone()

	iottwinmaker_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iottwinmaker_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag key names to remove from the resource.")
	iottwinmaker_untagResourceCmd.MarkFlagRequired("resource-arn")
	iottwinmaker_untagResourceCmd.MarkFlagRequired("tag-keys")
	iottwinmakerCmd.AddCommand(iottwinmaker_untagResourceCmd)
}
