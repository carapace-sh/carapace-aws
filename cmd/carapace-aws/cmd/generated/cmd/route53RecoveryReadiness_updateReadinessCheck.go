package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_updateReadinessCheckCmd = &cobra.Command{
	Use:   "update-readiness-check",
	Short: "Updates a readiness check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_updateReadinessCheckCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_updateReadinessCheckCmd).Standalone()

		route53RecoveryReadiness_updateReadinessCheckCmd.Flags().String("readiness-check-name", "", "Name of a readiness check.")
		route53RecoveryReadiness_updateReadinessCheckCmd.Flags().String("resource-set-name", "", "The name of the resource set to be checked.")
		route53RecoveryReadiness_updateReadinessCheckCmd.MarkFlagRequired("readiness-check-name")
		route53RecoveryReadiness_updateReadinessCheckCmd.MarkFlagRequired("resource-set-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_updateReadinessCheckCmd)
}
