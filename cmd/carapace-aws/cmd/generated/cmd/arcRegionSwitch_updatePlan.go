package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_updatePlanCmd = &cobra.Command{
	Use:   "update-plan",
	Short: "Updates an existing Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_updatePlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_updatePlanCmd).Standalone()

		arcRegionSwitch_updatePlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the plan.")
		arcRegionSwitch_updatePlanCmd.Flags().String("associated-alarms", "", "The updated CloudWatch alarms associated with the plan.")
		arcRegionSwitch_updatePlanCmd.Flags().String("description", "", "The updated description for the Region switch plan.")
		arcRegionSwitch_updatePlanCmd.Flags().String("execution-role", "", "The updated IAM role ARN that grants Region switch the permissions needed to execute the plan steps.")
		arcRegionSwitch_updatePlanCmd.Flags().String("recovery-time-objective-minutes", "", "The updated target recovery time objective (RTO) in minutes for the plan.")
		arcRegionSwitch_updatePlanCmd.Flags().String("triggers", "", "The updated conditions that can automatically trigger the execution of the plan.")
		arcRegionSwitch_updatePlanCmd.Flags().String("workflows", "", "The updated workflows for the Region switch plan.")
		arcRegionSwitch_updatePlanCmd.MarkFlagRequired("arn")
		arcRegionSwitch_updatePlanCmd.MarkFlagRequired("execution-role")
		arcRegionSwitch_updatePlanCmd.MarkFlagRequired("workflows")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_updatePlanCmd)
}
