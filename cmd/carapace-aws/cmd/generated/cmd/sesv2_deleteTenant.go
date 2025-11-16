package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteTenantCmd = &cobra.Command{
	Use:   "delete-tenant",
	Short: "Delete an existing tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteTenantCmd).Standalone()

	sesv2_deleteTenantCmd.Flags().String("tenant-name", "", "The name of the tenant to delete.")
	sesv2_deleteTenantCmd.MarkFlagRequired("tenant-name")
	sesv2Cmd.AddCommand(sesv2_deleteTenantCmd)
}
