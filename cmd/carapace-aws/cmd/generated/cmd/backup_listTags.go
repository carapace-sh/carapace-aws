package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Returns the tags assigned to the resource, such as a target recovery point, backup plan, or backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listTagsCmd).Standalone()

		backup_listTagsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listTagsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
		backup_listTagsCmd.Flags().String("resource-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a resource.")
		backup_listTagsCmd.MarkFlagRequired("resource-arn")
	})
	backupCmd.AddCommand(backup_listTagsCmd)
}
