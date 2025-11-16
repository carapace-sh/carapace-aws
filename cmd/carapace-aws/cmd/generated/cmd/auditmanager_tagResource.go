package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags the specified resource in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_tagResourceCmd).Standalone()

	auditmanager_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	auditmanager_tagResourceCmd.Flags().String("tags", "", "The tags that are associated with the resource.")
	auditmanager_tagResourceCmd.MarkFlagRequired("resource-arn")
	auditmanager_tagResourceCmd.MarkFlagRequired("tags")
	auditmanagerCmd.AddCommand(auditmanager_tagResourceCmd)
}
