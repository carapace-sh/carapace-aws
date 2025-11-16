package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createTenantCmd = &cobra.Command{
	Use:   "create-tenant",
	Short: "Create a tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createTenantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createTenantCmd).Standalone()

		sesv2_createTenantCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) to associate with the tenant")
		sesv2_createTenantCmd.Flags().String("tenant-name", "", "The name of the tenant to create.")
		sesv2_createTenantCmd.MarkFlagRequired("tenant-name")
	})
	sesv2Cmd.AddCommand(sesv2_createTenantCmd)
}
