package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listAccountPermissionsCmd = &cobra.Command{
	Use:   "list-account-permissions",
	Short: "Lists the permissions an account has to configure Amazon Inspector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listAccountPermissionsCmd).Standalone()

	inspector2_listAccountPermissionsCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
	inspector2_listAccountPermissionsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	inspector2_listAccountPermissionsCmd.Flags().String("service", "", "The service scan type to check permissions for.")
	inspector2Cmd.AddCommand(inspector2_listAccountPermissionsCmd)
}
