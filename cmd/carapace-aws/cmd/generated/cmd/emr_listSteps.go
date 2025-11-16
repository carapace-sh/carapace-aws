package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listStepsCmd = &cobra.Command{
	Use:   "list-steps",
	Short: "Provides a list of steps for the cluster in reverse order unless you specify `stepIds` with the request or filter by `StepStates`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_listStepsCmd).Standalone()

		emr_listStepsCmd.Flags().String("cluster-id", "", "The identifier of the cluster for which to list the steps.")
		emr_listStepsCmd.Flags().String("marker", "", "The maximum number of steps that a single `ListSteps` action returns is 50. To return a longer list of steps, use multiple `ListSteps` actions along with the `Marker` parameter, which is a pagination token that indicates the next set of results to retrieve.")
		emr_listStepsCmd.Flags().String("step-ids", "", "The filter to limit the step list based on the identifier of the steps.")
		emr_listStepsCmd.Flags().String("step-states", "", "The filter to limit the step list based on certain states.")
		emr_listStepsCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_listStepsCmd)
}
