package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_executeFastResetCmd = &cobra.Command{
	Use:   "execute-fast-reset",
	Short: "The fast reset REST API lets you reset a Neptune graph quicky and easily, removing all of its data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_executeFastResetCmd).Standalone()

	neptunedata_executeFastResetCmd.Flags().String("action", "", "The fast reset action.")
	neptunedata_executeFastResetCmd.Flags().String("token", "", "The fast-reset token to initiate the reset.")
	neptunedata_executeFastResetCmd.MarkFlagRequired("action")
	neptunedataCmd.AddCommand(neptunedata_executeFastResetCmd)
}
