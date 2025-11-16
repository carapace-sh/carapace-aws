package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_startPracticeRunCmd = &cobra.Command{
	Use:   "start-practice-run",
	Short: "Start an on-demand practice run zonal shift in Amazon Application Recovery Controller.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_startPracticeRunCmd).Standalone()

	arcZonalShift_startPracticeRunCmd.Flags().String("away-from", "", "The Availability Zone (for example, `use1-az1`) that traffic is shifted away from for the resource that you specify for the practice run.")
	arcZonalShift_startPracticeRunCmd.Flags().String("comment", "", "The initial comment that you enter about the practice run.")
	arcZonalShift_startPracticeRunCmd.Flags().String("resource-identifier", "", "The identifier for the resource that you want to start a practice run zonal shift for.")
	arcZonalShift_startPracticeRunCmd.MarkFlagRequired("away-from")
	arcZonalShift_startPracticeRunCmd.MarkFlagRequired("comment")
	arcZonalShift_startPracticeRunCmd.MarkFlagRequired("resource-identifier")
	arcZonalShiftCmd.AddCommand(arcZonalShift_startPracticeRunCmd)
}
