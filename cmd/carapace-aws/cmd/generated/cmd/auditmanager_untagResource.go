package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_untagResourceCmd).Standalone()

		auditmanager_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the specified resource.")
		auditmanager_untagResourceCmd.Flags().String("tag-keys", "", "The name or key of the tag.")
		auditmanager_untagResourceCmd.MarkFlagRequired("resource-arn")
		auditmanager_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	auditmanagerCmd.AddCommand(auditmanager_untagResourceCmd)
}
