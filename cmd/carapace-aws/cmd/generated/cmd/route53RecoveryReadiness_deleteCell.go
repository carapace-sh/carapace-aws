package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_deleteCellCmd = &cobra.Command{
	Use:   "delete-cell",
	Short: "Delete a cell.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_deleteCellCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_deleteCellCmd).Standalone()

		route53RecoveryReadiness_deleteCellCmd.Flags().String("cell-name", "", "The name of the cell.")
		route53RecoveryReadiness_deleteCellCmd.MarkFlagRequired("cell-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_deleteCellCmd)
}
