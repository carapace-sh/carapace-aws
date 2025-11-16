package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_cancelPracticeRunCmd = &cobra.Command{
	Use:   "cancel-practice-run",
	Short: "Cancel an in-progress practice run zonal shift in Amazon Application Recovery Controller.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_cancelPracticeRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_cancelPracticeRunCmd).Standalone()

		arcZonalShift_cancelPracticeRunCmd.Flags().String("zonal-shift-id", "", "The identifier of a practice run zonal shift in Amazon Application Recovery Controller that you want to cancel.")
		arcZonalShift_cancelPracticeRunCmd.MarkFlagRequired("zonal-shift-id")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_cancelPracticeRunCmd)
}
