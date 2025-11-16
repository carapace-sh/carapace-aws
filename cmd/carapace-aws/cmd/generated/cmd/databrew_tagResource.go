package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds metadata tags to a DataBrew resource, such as a dataset, project, recipe, job, or schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_tagResourceCmd).Standalone()

	databrew_tagResourceCmd.Flags().String("resource-arn", "", "The DataBrew resource to which tags should be added.")
	databrew_tagResourceCmd.Flags().String("tags", "", "One or more tags to be assigned to the resource.")
	databrew_tagResourceCmd.MarkFlagRequired("resource-arn")
	databrew_tagResourceCmd.MarkFlagRequired("tags")
	databrewCmd.AddCommand(databrew_tagResourceCmd)
}
