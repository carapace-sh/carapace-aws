package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_updateResourceSetCmd = &cobra.Command{
	Use:   "update-resource-set",
	Short: "Updates a resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_updateResourceSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_updateResourceSetCmd).Standalone()

		route53RecoveryReadiness_updateResourceSetCmd.Flags().String("resource-set-name", "", "Name of a resource set.")
		route53RecoveryReadiness_updateResourceSetCmd.Flags().String("resource-set-type", "", "The resource type of the resources in the resource set.")
		route53RecoveryReadiness_updateResourceSetCmd.Flags().String("resources", "", "A list of resource objects.")
		route53RecoveryReadiness_updateResourceSetCmd.MarkFlagRequired("resource-set-name")
		route53RecoveryReadiness_updateResourceSetCmd.MarkFlagRequired("resource-set-type")
		route53RecoveryReadiness_updateResourceSetCmd.MarkFlagRequired("resources")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_updateResourceSetCmd)
}
