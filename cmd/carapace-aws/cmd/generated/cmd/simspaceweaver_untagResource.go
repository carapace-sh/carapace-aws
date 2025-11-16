package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a SimSpace Weaver resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_untagResourceCmd).Standalone()

	simspaceweaver_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove tags from.")
	simspaceweaver_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the resource.")
	simspaceweaver_untagResourceCmd.MarkFlagRequired("resource-arn")
	simspaceweaver_untagResourceCmd.MarkFlagRequired("tag-keys")
	simspaceweaverCmd.AddCommand(simspaceweaver_untagResourceCmd)
}
