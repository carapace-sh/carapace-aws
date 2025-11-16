package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getReadinessCheckCmd = &cobra.Command{
	Use:   "get-readiness-check",
	Short: "Gets details about a readiness check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getReadinessCheckCmd).Standalone()

	route53RecoveryReadiness_getReadinessCheckCmd.Flags().String("readiness-check-name", "", "Name of a readiness check.")
	route53RecoveryReadiness_getReadinessCheckCmd.MarkFlagRequired("readiness-check-name")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getReadinessCheckCmd)
}
