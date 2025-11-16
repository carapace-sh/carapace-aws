package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_updateZonalShiftCmd = &cobra.Command{
	Use:   "update-zonal-shift",
	Short: "Update an active zonal shift in Amazon Application Recovery Controller in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_updateZonalShiftCmd).Standalone()

	arcZonalShift_updateZonalShiftCmd.Flags().String("comment", "", "A comment that you enter about the zonal shift.")
	arcZonalShift_updateZonalShiftCmd.Flags().String("expires-in", "", "The length of time that you want a zonal shift to be active, which ARC converts to an expiry time (expiration time).")
	arcZonalShift_updateZonalShiftCmd.Flags().String("zonal-shift-id", "", "The identifier of a zonal shift.")
	arcZonalShift_updateZonalShiftCmd.MarkFlagRequired("zonal-shift-id")
	arcZonalShiftCmd.AddCommand(arcZonalShift_updateZonalShiftCmd)
}
