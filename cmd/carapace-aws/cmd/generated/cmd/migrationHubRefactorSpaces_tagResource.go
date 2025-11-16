package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Removes the tags of a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_tagResourceCmd).Standalone()

		migrationHubRefactorSpaces_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		migrationHubRefactorSpaces_tagResourceCmd.Flags().String("tags", "", "The new or modified tags for the resource.")
		migrationHubRefactorSpaces_tagResourceCmd.MarkFlagRequired("resource-arn")
		migrationHubRefactorSpaces_tagResourceCmd.MarkFlagRequired("tags")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_tagResourceCmd)
}
