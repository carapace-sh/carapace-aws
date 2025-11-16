package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_deleteResourceSetCmd = &cobra.Command{
	Use:   "delete-resource-set",
	Short: "Deletes a resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_deleteResourceSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_deleteResourceSetCmd).Standalone()

		route53RecoveryReadiness_deleteResourceSetCmd.Flags().String("resource-set-name", "", "Name of a resource set.")
		route53RecoveryReadiness_deleteResourceSetCmd.MarkFlagRequired("resource-set-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_deleteResourceSetCmd)
}
