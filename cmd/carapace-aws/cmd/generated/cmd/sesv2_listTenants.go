package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listTenantsCmd = &cobra.Command{
	Use:   "list-tenants",
	Short: "List all tenants associated with your account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listTenantsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listTenantsCmd).Standalone()

		sesv2_listTenantsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListTenants` to indicate the position in the list of tenants.")
		sesv2_listTenantsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListTenants`.")
	})
	sesv2Cmd.AddCommand(sesv2_listTenantsCmd)
}
