package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_listPlansInRegionCmd = &cobra.Command{
	Use:   "list-plans-in-region",
	Short: "Lists all Region switch plans in your Amazon Web Services account that are available in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_listPlansInRegionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_listPlansInRegionCmd).Standalone()

		arcRegionSwitch_listPlansInRegionCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		arcRegionSwitch_listPlansInRegionCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_listPlansInRegionCmd)
}
