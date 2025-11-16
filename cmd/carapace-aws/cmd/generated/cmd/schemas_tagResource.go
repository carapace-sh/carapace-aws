package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_tagResourceCmd).Standalone()

		schemas_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		schemas_tagResourceCmd.Flags().String("tags", "", "Tags associated with the resource.")
		schemas_tagResourceCmd.MarkFlagRequired("resource-arn")
		schemas_tagResourceCmd.MarkFlagRequired("tags")
	})
	schemasCmd.AddCommand(schemas_tagResourceCmd)
}
