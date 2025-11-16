package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_updateCaseStatusCmd = &cobra.Command{
	Use:   "update-case-status",
	Short: "Updates the state transitions for a designated cases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_updateCaseStatusCmd).Standalone()

	securityIr_updateCaseStatusCmd.Flags().String("case-id", "", "Required element for UpdateCaseStatus to identify the case to update.")
	securityIr_updateCaseStatusCmd.Flags().String("case-status", "", "Required element for UpdateCaseStatus to identify the status for a case.")
	securityIr_updateCaseStatusCmd.MarkFlagRequired("case-id")
	securityIr_updateCaseStatusCmd.MarkFlagRequired("case-status")
	securityIrCmd.AddCommand(securityIr_updateCaseStatusCmd)
}
