package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_createCellCmd = &cobra.Command{
	Use:   "create-cell",
	Short: "Creates a cell in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_createCellCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_createCellCmd).Standalone()

		route53RecoveryReadiness_createCellCmd.Flags().String("cell-name", "", "The name of the cell to create.")
		route53RecoveryReadiness_createCellCmd.Flags().String("cells", "", "A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells.")
		route53RecoveryReadiness_createCellCmd.Flags().String("tags", "", "")
		route53RecoveryReadiness_createCellCmd.MarkFlagRequired("cell-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_createCellCmd)
}
