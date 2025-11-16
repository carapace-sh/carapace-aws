package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_updateCellCmd = &cobra.Command{
	Use:   "update-cell",
	Short: "Updates a cell to replace the list of nested cells with a new list of nested cells.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_updateCellCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_updateCellCmd).Standalone()

		route53RecoveryReadiness_updateCellCmd.Flags().String("cell-name", "", "The name of the cell.")
		route53RecoveryReadiness_updateCellCmd.Flags().String("cells", "", "A list of cell Amazon Resource Names (ARNs), which completely replaces the previous list.")
		route53RecoveryReadiness_updateCellCmd.MarkFlagRequired("cell-name")
		route53RecoveryReadiness_updateCellCmd.MarkFlagRequired("cells")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_updateCellCmd)
}
