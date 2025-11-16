package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags attached to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_listTagsForResourceCmd).Standalone()

	rolesanywhere_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	rolesanywhere_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	rolesanywhereCmd.AddCommand(rolesanywhere_listTagsForResourceCmd)
}
