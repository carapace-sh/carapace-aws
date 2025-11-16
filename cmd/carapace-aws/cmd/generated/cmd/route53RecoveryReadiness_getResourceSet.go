package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getResourceSetCmd = &cobra.Command{
	Use:   "get-resource-set",
	Short: "Displays the details about a resource set, including a list of the resources in the set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getResourceSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_getResourceSetCmd).Standalone()

		route53RecoveryReadiness_getResourceSetCmd.Flags().String("resource-set-name", "", "Name of a resource set.")
		route53RecoveryReadiness_getResourceSetCmd.MarkFlagRequired("resource-set-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getResourceSetCmd)
}
