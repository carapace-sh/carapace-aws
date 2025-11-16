package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_listPlansCmd = &cobra.Command{
	Use:   "list-plans",
	Short: "Lists all Region switch plans in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_listPlansCmd).Standalone()

	arcRegionSwitch_listPlansCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	arcRegionSwitch_listPlansCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_listPlansCmd)
}
