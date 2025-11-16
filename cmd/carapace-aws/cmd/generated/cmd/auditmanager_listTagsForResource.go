package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for the specified resource in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listTagsForResourceCmd).Standalone()

		auditmanager_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		auditmanager_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	auditmanagerCmd.AddCommand(auditmanager_listTagsForResourceCmd)
}
