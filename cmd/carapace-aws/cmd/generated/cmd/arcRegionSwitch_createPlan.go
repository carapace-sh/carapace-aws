package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_createPlanCmd = &cobra.Command{
	Use:   "create-plan",
	Short: "Creates a new Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_createPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_createPlanCmd).Standalone()

		arcRegionSwitch_createPlanCmd.Flags().String("associated-alarms", "", "The alarms associated with a Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("description", "", "The description of a Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("execution-role", "", "An execution role is a way to categorize a Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("name", "", "The name of a Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("primary-region", "", "The primary Amazon Web Services Region for the application.")
		arcRegionSwitch_createPlanCmd.Flags().String("recovery-approach", "", "The recovery approach for a Region switch plan, which can be active/active (activeActive) or active/passive (activePassive).")
		arcRegionSwitch_createPlanCmd.Flags().String("recovery-time-objective-minutes", "", "Optionally, you can specify an recovery time objective for a Region switch plan, in minutes.")
		arcRegionSwitch_createPlanCmd.Flags().String("regions", "", "An array that specifies the Amazon Web Services Regions for a Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("tags", "", "The tags to apply to the Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("triggers", "", "The triggers associated with a Region switch plan.")
		arcRegionSwitch_createPlanCmd.Flags().String("workflows", "", "An array of workflows included in a Region switch plan.")
		arcRegionSwitch_createPlanCmd.MarkFlagRequired("execution-role")
		arcRegionSwitch_createPlanCmd.MarkFlagRequired("name")
		arcRegionSwitch_createPlanCmd.MarkFlagRequired("recovery-approach")
		arcRegionSwitch_createPlanCmd.MarkFlagRequired("regions")
		arcRegionSwitch_createPlanCmd.MarkFlagRequired("workflows")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_createPlanCmd)
}
