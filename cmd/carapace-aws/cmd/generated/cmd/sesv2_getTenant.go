package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getTenantCmd = &cobra.Command{
	Use:   "get-tenant",
	Short: "Get information about a specific tenant, including the tenant's name, ID, ARN, creation timestamp, tags, and sending status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getTenantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getTenantCmd).Standalone()

		sesv2_getTenantCmd.Flags().String("tenant-name", "", "The name of the tenant to retrieve information about.")
		sesv2_getTenantCmd.MarkFlagRequired("tenant-name")
	})
	sesv2Cmd.AddCommand(sesv2_getTenantCmd)
}
