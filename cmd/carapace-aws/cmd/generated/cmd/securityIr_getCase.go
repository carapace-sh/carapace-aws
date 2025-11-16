package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_getCaseCmd = &cobra.Command{
	Use:   "get-case",
	Short: "Returns the attributes of a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_getCaseCmd).Standalone()

	securityIr_getCaseCmd.Flags().String("case-id", "", "Required element for GetCase to identify the requested case ID.")
	securityIr_getCaseCmd.MarkFlagRequired("case-id")
	securityIrCmd.AddCommand(securityIr_getCaseCmd)
}
