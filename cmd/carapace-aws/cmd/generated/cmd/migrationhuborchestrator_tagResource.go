package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhuborchestrator_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag a resource by specifying its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhuborchestrator_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhuborchestrator_tagResourceCmd).Standalone()

		migrationhuborchestrator_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which you want to add tags.")
		migrationhuborchestrator_tagResourceCmd.Flags().String("tags", "", "A collection of labels, in the form of key:value pairs, that apply to this resource.")
		migrationhuborchestrator_tagResourceCmd.MarkFlagRequired("resource-arn")
		migrationhuborchestrator_tagResourceCmd.MarkFlagRequired("tags")
	})
	migrationhuborchestratorCmd.AddCommand(migrationhuborchestrator_tagResourceCmd)
}
