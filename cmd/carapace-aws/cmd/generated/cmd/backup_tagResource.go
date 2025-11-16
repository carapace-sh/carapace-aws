package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns a set of key-value pairs to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_tagResourceCmd).Standalone()

		backup_tagResourceCmd.Flags().String("resource-arn", "", "The ARN that uniquely identifies the resource.")
		backup_tagResourceCmd.Flags().String("tags", "", "Key-value pairs that are used to help organize your resources.")
		backup_tagResourceCmd.MarkFlagRequired("resource-arn")
		backup_tagResourceCmd.MarkFlagRequired("tags")
	})
	backupCmd.AddCommand(backup_tagResourceCmd)
}
