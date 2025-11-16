package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_updatePracticeRunConfigurationCmd = &cobra.Command{
	Use:   "update-practice-run-configuration",
	Short: "Update a practice run configuration to change one or more of the following: add, change, or remove the blocking alarm; change the outcome alarm; or add, change, or remove blocking dates or time windows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_updatePracticeRunConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_updatePracticeRunConfigurationCmd).Standalone()

		arcZonalShift_updatePracticeRunConfigurationCmd.Flags().String("allowed-windows", "", "Add, change, or remove windows of days and times for when you can, optionally, allow ARC to start a practice run for a resource.")
		arcZonalShift_updatePracticeRunConfigurationCmd.Flags().String("blocked-dates", "", "Add, change, or remove blocked dates for a practice run in zonal autoshift.")
		arcZonalShift_updatePracticeRunConfigurationCmd.Flags().String("blocked-windows", "", "Add, change, or remove windows of days and times for when you can, optionally, block ARC from starting a practice run for a resource.")
		arcZonalShift_updatePracticeRunConfigurationCmd.Flags().String("blocking-alarms", "", "Add, change, or remove the Amazon CloudWatch alarms that you optionally specify as the blocking alarms for practice runs.")
		arcZonalShift_updatePracticeRunConfigurationCmd.Flags().String("outcome-alarms", "", "Specify one or more Amazon CloudWatch alarms as the outcome alarms for practice runs.")
		arcZonalShift_updatePracticeRunConfigurationCmd.Flags().String("resource-identifier", "", "The identifier for the resource that you want to update the practice run configuration for.")
		arcZonalShift_updatePracticeRunConfigurationCmd.MarkFlagRequired("resource-identifier")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_updatePracticeRunConfigurationCmd)
}
