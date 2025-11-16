package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_startZonalShiftCmd = &cobra.Command{
	Use:   "start-zonal-shift",
	Short: "You start a zonal shift to temporarily move load balancer traffic away from an Availability Zone in an Amazon Web Services Region, to help your application recover immediately, for example, from a developer's bad code deployment or from an Amazon Web Services infrastructure failure in a single Availability Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_startZonalShiftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_startZonalShiftCmd).Standalone()

		arcZonalShift_startZonalShiftCmd.Flags().String("away-from", "", "The Availability Zone (for example, `use1-az1`) that traffic is moved away from for a resource when you start a zonal shift.")
		arcZonalShift_startZonalShiftCmd.Flags().String("comment", "", "A comment that you enter about the zonal shift.")
		arcZonalShift_startZonalShiftCmd.Flags().String("expires-in", "", "The length of time that you want a zonal shift to be active, which ARC converts to an expiry time (expiration time).")
		arcZonalShift_startZonalShiftCmd.Flags().String("resource-identifier", "", "The identifier for the resource that Amazon Web Services shifts traffic for.")
		arcZonalShift_startZonalShiftCmd.MarkFlagRequired("away-from")
		arcZonalShift_startZonalShiftCmd.MarkFlagRequired("comment")
		arcZonalShift_startZonalShiftCmd.MarkFlagRequired("expires-in")
		arcZonalShift_startZonalShiftCmd.MarkFlagRequired("resource-identifier")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_startZonalShiftCmd)
}
