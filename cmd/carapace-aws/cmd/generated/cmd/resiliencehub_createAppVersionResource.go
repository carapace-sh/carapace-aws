package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_createAppVersionResourceCmd = &cobra.Command{
	Use:   "create-app-version-resource",
	Short: "Adds a resource to the Resilience Hub application and assigns it to the specified Application Components.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_createAppVersionResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_createAppVersionResourceCmd).Standalone()

		resiliencehub_createAppVersionResourceCmd.Flags().String("additional-info", "", "Currently, there is no supported additional information for resources.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("app-components", "", "List of Application Components that this resource belongs to.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("aws-account-id", "", "Amazon Web Services account that owns the physical resource.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("aws-region", "", "Amazon Web Services region that owns the physical resource.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("logical-resource-id", "", "Logical identifier of the resource.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("physical-resource-id", "", "Physical identifier of the resource.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("resource-name", "", "Name of the resource.")
		resiliencehub_createAppVersionResourceCmd.Flags().String("resource-type", "", "Type of resource.")
		resiliencehub_createAppVersionResourceCmd.MarkFlagRequired("app-arn")
		resiliencehub_createAppVersionResourceCmd.MarkFlagRequired("app-components")
		resiliencehub_createAppVersionResourceCmd.MarkFlagRequired("logical-resource-id")
		resiliencehub_createAppVersionResourceCmd.MarkFlagRequired("physical-resource-id")
		resiliencehub_createAppVersionResourceCmd.MarkFlagRequired("resource-type")
	})
	resiliencehubCmd.AddCommand(resiliencehub_createAppVersionResourceCmd)
}
