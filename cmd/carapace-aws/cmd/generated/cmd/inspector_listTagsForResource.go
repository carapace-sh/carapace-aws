package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with an assessment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listTagsForResourceCmd).Standalone()

	inspector_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN that specifies the assessment template whose tags you want to list.")
	inspector_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	inspectorCmd.AddCommand(inspector_listTagsForResourceCmd)
}
