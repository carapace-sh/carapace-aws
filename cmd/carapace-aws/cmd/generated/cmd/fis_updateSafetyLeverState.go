package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_updateSafetyLeverStateCmd = &cobra.Command{
	Use:   "update-safety-lever-state",
	Short: "Updates the specified safety lever state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_updateSafetyLeverStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_updateSafetyLeverStateCmd).Standalone()

		fis_updateSafetyLeverStateCmd.Flags().String("id", "", "The ID of the safety lever.")
		fis_updateSafetyLeverStateCmd.Flags().String("state", "", "The state of the safety lever.")
		fis_updateSafetyLeverStateCmd.MarkFlagRequired("id")
		fis_updateSafetyLeverStateCmd.MarkFlagRequired("state")
	})
	fisCmd.AddCommand(fis_updateSafetyLeverStateCmd)
}
