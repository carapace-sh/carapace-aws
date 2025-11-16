package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getCellCmd = &cobra.Command{
	Use:   "get-cell",
	Short: "Gets information about a cell including cell name, cell Amazon Resource Name (ARN), ARNs of nested cells for this cell, and a list of those cell ARNs with their associated recovery group ARNs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getCellCmd).Standalone()

	route53RecoveryReadiness_getCellCmd.Flags().String("cell-name", "", "The name of the cell.")
	route53RecoveryReadiness_getCellCmd.MarkFlagRequired("cell-name")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getCellCmd)
}
