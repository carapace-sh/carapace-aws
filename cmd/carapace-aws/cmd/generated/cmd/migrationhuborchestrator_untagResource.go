package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_untagResourceCmd).Standalone()

	migrationhuborchestrator_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which you want to remove tags.")
	migrationhuborchestrator_untagResourceCmd.Flags().String("tag-keys", "", "One or more tag keys.")
	migrationhuborchestrator_untagResourceCmd.MarkFlagRequired("resource-arn")
	migrationhuborchestrator_untagResourceCmd.MarkFlagRequired("tag-keys")
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_untagResourceCmd)
}
