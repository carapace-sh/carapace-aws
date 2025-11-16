package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_closeCaseCmd = &cobra.Command{
	Use:   "close-case",
	Short: "Closes an existing case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_closeCaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_closeCaseCmd).Standalone()

		securityIr_closeCaseCmd.Flags().String("case-id", "", "Required element used in combination with CloseCase to identify the case ID to close.")
		securityIr_closeCaseCmd.MarkFlagRequired("case-id")
	})
	securityIrCmd.AddCommand(securityIr_closeCaseCmd)
}
