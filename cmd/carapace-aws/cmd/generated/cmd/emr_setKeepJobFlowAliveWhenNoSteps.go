package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_setKeepJobFlowAliveWhenNoStepsCmd = &cobra.Command{
	Use:   "set-keep-job-flow-alive-when-no-steps",
	Short: "You can use the `SetKeepJobFlowAliveWhenNoSteps` to configure a cluster (job flow) to terminate after the step execution, i.e., all your steps are executed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_setKeepJobFlowAliveWhenNoStepsCmd).Standalone()

	emr_setKeepJobFlowAliveWhenNoStepsCmd.Flags().String("job-flow-ids", "", "A list of strings that uniquely identify the clusters to protect.")
	emr_setKeepJobFlowAliveWhenNoStepsCmd.Flags().Bool("keep-job-flow-alive-when-no-steps", false, "A Boolean that indicates whether to terminate the cluster after all steps are executed.")
	emr_setKeepJobFlowAliveWhenNoStepsCmd.Flags().Bool("no-keep-job-flow-alive-when-no-steps", false, "A Boolean that indicates whether to terminate the cluster after all steps are executed.")
	emr_setKeepJobFlowAliveWhenNoStepsCmd.MarkFlagRequired("job-flow-ids")
	emr_setKeepJobFlowAliveWhenNoStepsCmd.MarkFlagRequired("keep-job-flow-alive-when-no-steps")
	emr_setKeepJobFlowAliveWhenNoStepsCmd.Flag("no-keep-job-flow-alive-when-no-steps").Hidden = true
	emr_setKeepJobFlowAliveWhenNoStepsCmd.MarkFlagRequired("no-keep-job-flow-alive-when-no-steps")
	emrCmd.AddCommand(emr_setKeepJobFlowAliveWhenNoStepsCmd)
}
