package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Retrieves a list of tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getTagsCmd).Standalone()

		glue_getTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which to retrieve tags.")
		glue_getTagsCmd.MarkFlagRequired("resource-arn")
	})
	glueCmd.AddCommand(glue_getTagsCmd)
}
