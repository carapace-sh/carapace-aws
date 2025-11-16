package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_untagResourceCmd).Standalone()

		schemas_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		schemas_untagResourceCmd.Flags().String("tag-keys", "", "Keys of key-value pairs.")
		schemas_untagResourceCmd.MarkFlagRequired("resource-arn")
		schemas_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	schemasCmd.AddCommand(schemas_untagResourceCmd)
}
