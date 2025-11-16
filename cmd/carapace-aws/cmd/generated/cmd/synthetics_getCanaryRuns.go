package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_getCanaryRunsCmd = &cobra.Command{
	Use:   "get-canary-runs",
	Short: "Retrieves a list of runs for a specified canary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_getCanaryRunsCmd).Standalone()

	synthetics_getCanaryRunsCmd.Flags().String("dry-run-id", "", "The DryRunId associated with an existing canary’s dry run.")
	synthetics_getCanaryRunsCmd.Flags().String("max-results", "", "Specify this parameter to limit how many runs are returned each time you use the `GetCanaryRuns` operation.")
	synthetics_getCanaryRunsCmd.Flags().String("name", "", "The name of the canary that you want to see runs for.")
	synthetics_getCanaryRunsCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	synthetics_getCanaryRunsCmd.Flags().String("run-type", "", "- When you provide `RunType=CANARY_RUN` and `dryRunId`, you will get an exception")
	synthetics_getCanaryRunsCmd.MarkFlagRequired("name")
	syntheticsCmd.AddCommand(synthetics_getCanaryRunsCmd)
}
