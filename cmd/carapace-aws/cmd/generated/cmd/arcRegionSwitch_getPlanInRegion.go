package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_getPlanInRegionCmd = &cobra.Command{
	Use:   "get-plan-in-region",
	Short: "Retrieves information about a Region switch plan in a specific Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_getPlanInRegionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_getPlanInRegionCmd).Standalone()

		arcRegionSwitch_getPlanInRegionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the plan in Region.")
		arcRegionSwitch_getPlanInRegionCmd.MarkFlagRequired("arn")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_getPlanInRegionCmd)
}
