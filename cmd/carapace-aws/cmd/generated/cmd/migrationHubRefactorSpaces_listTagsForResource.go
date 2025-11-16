package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags of a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_listTagsForResourceCmd).Standalone()

	migrationHubRefactorSpaces_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	migrationHubRefactorSpaces_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_listTagsForResourceCmd)
}
