package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_tagResourceCmd).Standalone()

		glue_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the Glue resource to which to add the tags.")
		glue_tagResourceCmd.Flags().String("tags-to-add", "", "Tags to add to this resource.")
		glue_tagResourceCmd.MarkFlagRequired("resource-arn")
		glue_tagResourceCmd.MarkFlagRequired("tags-to-add")
	})
	glueCmd.AddCommand(glue_tagResourceCmd)
}
