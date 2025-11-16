package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_listAutoshiftsCmd = &cobra.Command{
	Use:   "list-autoshifts",
	Short: "Returns the autoshifts for an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_listAutoshiftsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_listAutoshiftsCmd).Standalone()

		arcZonalShift_listAutoshiftsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		arcZonalShift_listAutoshiftsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		arcZonalShift_listAutoshiftsCmd.Flags().String("status", "", "The status of the autoshift.")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_listAutoshiftsCmd)
}
