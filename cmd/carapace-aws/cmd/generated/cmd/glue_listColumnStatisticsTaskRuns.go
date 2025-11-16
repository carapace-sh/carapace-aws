package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listColumnStatisticsTaskRunsCmd = &cobra.Command{
	Use:   "list-column-statistics-task-runs",
	Short: "List all task runs for a particular account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listColumnStatisticsTaskRunsCmd).Standalone()

	glue_listColumnStatisticsTaskRunsCmd.Flags().String("max-results", "", "The maximum size of the response.")
	glue_listColumnStatisticsTaskRunsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	glueCmd.AddCommand(glue_listColumnStatisticsTaskRunsCmd)
}
