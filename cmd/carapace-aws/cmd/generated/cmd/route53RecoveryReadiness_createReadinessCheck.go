package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_createReadinessCheckCmd = &cobra.Command{
	Use:   "create-readiness-check",
	Short: "Creates a readiness check in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_createReadinessCheckCmd).Standalone()

	route53RecoveryReadiness_createReadinessCheckCmd.Flags().String("readiness-check-name", "", "The name of the readiness check to create.")
	route53RecoveryReadiness_createReadinessCheckCmd.Flags().String("resource-set-name", "", "The name of the resource set to check.")
	route53RecoveryReadiness_createReadinessCheckCmd.Flags().String("tags", "", "")
	route53RecoveryReadiness_createReadinessCheckCmd.MarkFlagRequired("readiness-check-name")
	route53RecoveryReadiness_createReadinessCheckCmd.MarkFlagRequired("resource-set-name")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_createReadinessCheckCmd)
}
