package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listTenantResourcesCmd = &cobra.Command{
	Use:   "list-tenant-resources",
	Short: "List all resources associated with a specific tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listTenantResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listTenantResourcesCmd).Standalone()

		sesv2_listTenantResourcesCmd.Flags().String("filter", "", "A map of filter keys and values for filtering the list of tenant resources.")
		sesv2_listTenantResourcesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListTenantResources` to indicate the position in the list of tenant resources.")
		sesv2_listTenantResourcesCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListTenantResources`.")
		sesv2_listTenantResourcesCmd.Flags().String("tenant-name", "", "The name of the tenant to list resources for.")
		sesv2_listTenantResourcesCmd.MarkFlagRequired("tenant-name")
	})
	sesv2Cmd.AddCommand(sesv2_listTenantResourcesCmd)
}
