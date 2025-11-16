package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_listZonalShiftsCmd = &cobra.Command{
	Use:   "list-zonal-shifts",
	Short: "Lists all active and completed zonal shifts in Amazon Application Recovery Controller in your Amazon Web Services account in this Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_listZonalShiftsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_listZonalShiftsCmd).Standalone()

		arcZonalShift_listZonalShiftsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		arcZonalShift_listZonalShiftsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		arcZonalShift_listZonalShiftsCmd.Flags().String("resource-identifier", "", "The identifier for the resource that you want to list zonal shifts for.")
		arcZonalShift_listZonalShiftsCmd.Flags().String("status", "", "A status for a zonal shift.")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_listZonalShiftsCmd)
}
