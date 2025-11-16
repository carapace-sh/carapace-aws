package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getMltaskRunsCmd = &cobra.Command{
	Use:   "get-mltask-runs",
	Short: "Gets a list of runs for a machine learning transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getMltaskRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getMltaskRunsCmd).Standalone()

		glue_getMltaskRunsCmd.Flags().String("filter", "", "The filter criteria, in the `TaskRunFilterCriteria` structure, for the task run.")
		glue_getMltaskRunsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		glue_getMltaskRunsCmd.Flags().String("next-token", "", "A token for pagination of the results.")
		glue_getMltaskRunsCmd.Flags().String("sort", "", "The sorting criteria, in the `TaskRunSortCriteria` structure, for the task run.")
		glue_getMltaskRunsCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_getMltaskRunsCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_getMltaskRunsCmd)
}
