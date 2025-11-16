package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getMltransformsCmd = &cobra.Command{
	Use:   "get-mltransforms",
	Short: "Gets a sortable, filterable list of existing Glue machine learning transforms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getMltransformsCmd).Standalone()

	glue_getMltransformsCmd.Flags().String("filter", "", "The filter transformation criteria.")
	glue_getMltransformsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	glue_getMltransformsCmd.Flags().String("next-token", "", "A paginated token to offset the results.")
	glue_getMltransformsCmd.Flags().String("sort", "", "The sorting criteria.")
	glueCmd.AddCommand(glue_getMltransformsCmd)
}
