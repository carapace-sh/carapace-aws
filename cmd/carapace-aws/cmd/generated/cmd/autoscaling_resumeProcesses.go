package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_resumeProcessesCmd = &cobra.Command{
	Use:   "resume-processes",
	Short: "Resumes the specified suspended auto scaling processes, or all suspended process, for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_resumeProcessesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_resumeProcessesCmd).Standalone()

		autoscaling_resumeProcessesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_resumeProcessesCmd.Flags().String("scaling-processes", "", "One or more of the following processes:")
		autoscaling_resumeProcessesCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_resumeProcessesCmd)
}
