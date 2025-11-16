package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(aiops_tagResourceCmd).Standalone()

		aiops_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to apply the tags to.")
		aiops_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
		aiops_tagResourceCmd.MarkFlagRequired("resource-arn")
		aiops_tagResourceCmd.MarkFlagRequired("tags")
	})
	aiopsCmd.AddCommand(aiops_tagResourceCmd)
}
