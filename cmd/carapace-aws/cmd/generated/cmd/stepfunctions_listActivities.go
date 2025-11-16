package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listActivitiesCmd = &cobra.Command{
	Use:   "list-activities",
	Short: "Lists the existing activities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listActivitiesCmd).Standalone()

	stepfunctions_listActivitiesCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	stepfunctions_listActivitiesCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	stepfunctionsCmd.AddCommand(stepfunctions_listActivitiesCmd)
}
