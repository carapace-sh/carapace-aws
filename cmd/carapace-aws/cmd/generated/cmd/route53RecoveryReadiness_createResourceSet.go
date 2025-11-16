package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_createResourceSetCmd = &cobra.Command{
	Use:   "create-resource-set",
	Short: "Creates a resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_createResourceSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_createResourceSetCmd).Standalone()

		route53RecoveryReadiness_createResourceSetCmd.Flags().String("resource-set-name", "", "The name of the resource set to create.")
		route53RecoveryReadiness_createResourceSetCmd.Flags().String("resource-set-type", "", "The resource type of the resources in the resource set.")
		route53RecoveryReadiness_createResourceSetCmd.Flags().String("resources", "", "A list of resource objects in the resource set.")
		route53RecoveryReadiness_createResourceSetCmd.Flags().String("tags", "", "A tag to associate with the parameters for a resource set.")
		route53RecoveryReadiness_createResourceSetCmd.MarkFlagRequired("resource-set-name")
		route53RecoveryReadiness_createResourceSetCmd.MarkFlagRequired("resource-set-type")
		route53RecoveryReadiness_createResourceSetCmd.MarkFlagRequired("resources")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_createResourceSetCmd)
}
