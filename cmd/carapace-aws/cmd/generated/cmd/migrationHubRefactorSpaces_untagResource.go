package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_untagResourceCmd).Standalone()

	migrationHubRefactorSpaces_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	migrationHubRefactorSpaces_untagResourceCmd.Flags().String("tag-keys", "", "The list of keys of the tags to be removed from the resource.")
	migrationHubRefactorSpaces_untagResourceCmd.MarkFlagRequired("resource-arn")
	migrationHubRefactorSpaces_untagResourceCmd.MarkFlagRequired("tag-keys")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_untagResourceCmd)
}
