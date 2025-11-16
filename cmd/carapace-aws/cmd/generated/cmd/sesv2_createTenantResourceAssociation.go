package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createTenantResourceAssociationCmd = &cobra.Command{
	Use:   "create-tenant-resource-association",
	Short: "Associate a resource with a tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createTenantResourceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createTenantResourceAssociationCmd).Standalone()

		sesv2_createTenantResourceAssociationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to associate with the tenant.")
		sesv2_createTenantResourceAssociationCmd.Flags().String("tenant-name", "", "The name of the tenant to associate the resource with.")
		sesv2_createTenantResourceAssociationCmd.MarkFlagRequired("resource-arn")
		sesv2_createTenantResourceAssociationCmd.MarkFlagRequired("tenant-name")
	})
	sesv2Cmd.AddCommand(sesv2_createTenantResourceAssociationCmd)
}
