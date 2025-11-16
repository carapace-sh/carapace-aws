package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_updateAppVersionResourceCmd = &cobra.Command{
	Use:   "update-app-version-resource",
	Short: "Updates the resource details in the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_updateAppVersionResourceCmd).Standalone()

	resiliencehub_updateAppVersionResourceCmd.Flags().String("additional-info", "", "Currently, there is no supported additional information for resources.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("app-components", "", "List of Application Components that this resource belongs to.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("aws-account-id", "", "Amazon Web Services account that owns the physical resource.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("aws-region", "", "Amazon Web Services region that owns the physical resource.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("excluded", "", "Indicates if a resource is excluded from an Resilience Hub application.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("logical-resource-id", "", "Logical identifier of the resource.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("physical-resource-id", "", "Physical identifier of the resource.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("resource-name", "", "Name of the resource.")
	resiliencehub_updateAppVersionResourceCmd.Flags().String("resource-type", "", "Type of resource.")
	resiliencehub_updateAppVersionResourceCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_updateAppVersionResourceCmd)
}
