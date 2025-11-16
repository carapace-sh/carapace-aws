package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listResourceTenantsCmd = &cobra.Command{
	Use:   "list-resource-tenants",
	Short: "List all tenants associated with a specific resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listResourceTenantsCmd).Standalone()

	sesv2_listResourceTenantsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListResourceTenants` to indicate the position in the list of resource tenants.")
	sesv2_listResourceTenantsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListResourceTenants`.")
	sesv2_listResourceTenantsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to list associated tenants for.")
	sesv2_listResourceTenantsCmd.MarkFlagRequired("resource-arn")
	sesv2Cmd.AddCommand(sesv2_listResourceTenantsCmd)
}
