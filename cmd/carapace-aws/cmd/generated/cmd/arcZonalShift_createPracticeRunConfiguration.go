package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_createPracticeRunConfigurationCmd = &cobra.Command{
	Use:   "create-practice-run-configuration",
	Short: "A practice run configuration for zonal autoshift is required when you enable zonal autoshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_createPracticeRunConfigurationCmd).Standalone()

	arcZonalShift_createPracticeRunConfigurationCmd.Flags().String("allowed-windows", "", "Optionally, you can allow ARC to start practice runs for specific windows of days and times.")
	arcZonalShift_createPracticeRunConfigurationCmd.Flags().String("blocked-dates", "", "Optionally, you can block ARC from starting practice runs for a resource on specific calendar dates.")
	arcZonalShift_createPracticeRunConfigurationCmd.Flags().String("blocked-windows", "", "Optionally, you can block ARC from starting practice runs for specific windows of days and times.")
	arcZonalShift_createPracticeRunConfigurationCmd.Flags().String("blocking-alarms", "", "*Blocking alarms* for practice runs are optional alarms that you can specify that block practice runs when one or more of the alarms is in an `ALARM` state.")
	arcZonalShift_createPracticeRunConfigurationCmd.Flags().String("outcome-alarms", "", "*Outcome alarms* for practice runs are alarms that you specify that end a practice run when one or more of the alarms is in an `ALARM` state.")
	arcZonalShift_createPracticeRunConfigurationCmd.Flags().String("resource-identifier", "", "The identifier of the resource that Amazon Web Services shifts traffic for with a practice run zonal shift.")
	arcZonalShift_createPracticeRunConfigurationCmd.MarkFlagRequired("outcome-alarms")
	arcZonalShift_createPracticeRunConfigurationCmd.MarkFlagRequired("resource-identifier")
	arcZonalShiftCmd.AddCommand(arcZonalShift_createPracticeRunConfigurationCmd)
}
