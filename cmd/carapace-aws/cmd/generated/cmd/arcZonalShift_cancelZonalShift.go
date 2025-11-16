package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_cancelZonalShiftCmd = &cobra.Command{
	Use:   "cancel-zonal-shift",
	Short: "Cancel a zonal shift in Amazon Application Recovery Controller.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_cancelZonalShiftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_cancelZonalShiftCmd).Standalone()

		arcZonalShift_cancelZonalShiftCmd.Flags().String("zonal-shift-id", "", "The internally-generated identifier of a zonal shift.")
		arcZonalShift_cancelZonalShiftCmd.MarkFlagRequired("zonal-shift-id")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_cancelZonalShiftCmd)
}
