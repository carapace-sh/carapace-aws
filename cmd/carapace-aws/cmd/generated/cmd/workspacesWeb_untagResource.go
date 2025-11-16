package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_untagResourceCmd).Standalone()

		workspacesWeb_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		workspacesWeb_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
		workspacesWeb_untagResourceCmd.MarkFlagRequired("resource-arn")
		workspacesWeb_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_untagResourceCmd)
}
