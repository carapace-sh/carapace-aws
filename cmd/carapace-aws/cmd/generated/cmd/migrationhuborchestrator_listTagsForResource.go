package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags added to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_listTagsForResourceCmd).Standalone()

	migrationhuborchestrator_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	migrationhuborchestrator_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_listTagsForResourceCmd)
}
