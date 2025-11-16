package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteTenantResourceAssociationCmd = &cobra.Command{
	Use:   "delete-tenant-resource-association",
	Short: "Delete an association between a tenant and a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteTenantResourceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_deleteTenantResourceAssociationCmd).Standalone()

		sesv2_deleteTenantResourceAssociationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove from the tenant association.")
		sesv2_deleteTenantResourceAssociationCmd.Flags().String("tenant-name", "", "The name of the tenant to remove the resource association from.")
		sesv2_deleteTenantResourceAssociationCmd.MarkFlagRequired("resource-arn")
		sesv2_deleteTenantResourceAssociationCmd.MarkFlagRequired("tenant-name")
	})
	sesv2Cmd.AddCommand(sesv2_deleteTenantResourceAssociationCmd)
}
