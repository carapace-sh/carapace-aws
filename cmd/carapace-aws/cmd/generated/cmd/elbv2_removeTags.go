package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes the specified tags from the specified Elastic Load Balancing resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_removeTagsCmd).Standalone()

	elbv2_removeTagsCmd.Flags().String("resource-arns", "", "The Amazon Resource Name (ARN) of the resource.")
	elbv2_removeTagsCmd.Flags().String("tag-keys", "", "The tag keys for the tags to remove.")
	elbv2_removeTagsCmd.MarkFlagRequired("resource-arns")
	elbv2_removeTagsCmd.MarkFlagRequired("tag-keys")
	elbv2Cmd.AddCommand(elbv2_removeTagsCmd)
}
