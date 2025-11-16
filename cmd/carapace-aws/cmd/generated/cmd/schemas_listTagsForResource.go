package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Get tags for resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_listTagsForResourceCmd).Standalone()

	schemas_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	schemas_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	schemasCmd.AddCommand(schemas_listTagsForResourceCmd)
}
