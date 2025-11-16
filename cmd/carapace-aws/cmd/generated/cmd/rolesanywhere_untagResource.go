package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_untagResourceCmd).Standalone()

		rolesanywhere_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		rolesanywhere_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys.")
		rolesanywhere_untagResourceCmd.MarkFlagRequired("resource-arn")
		rolesanywhere_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_untagResourceCmd)
}
