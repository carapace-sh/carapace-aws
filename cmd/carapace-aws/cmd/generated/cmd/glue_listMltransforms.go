package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listMltransformsCmd = &cobra.Command{
	Use:   "list-mltransforms",
	Short: "Retrieves a sortable, filterable list of existing Glue machine learning transforms in this Amazon Web Services account, or the resources with the specified tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listMltransformsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listMltransformsCmd).Standalone()

		glue_listMltransformsCmd.Flags().String("filter", "", "A `TransformFilterCriteria` used to filter the machine learning transforms.")
		glue_listMltransformsCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
		glue_listMltransformsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
		glue_listMltransformsCmd.Flags().String("sort", "", "A `TransformSortCriteria` used to sort the machine learning transforms.")
		glue_listMltransformsCmd.Flags().String("tags", "", "Specifies to return only these tagged resources.")
	})
	glueCmd.AddCommand(glue_listMltransformsCmd)
}
