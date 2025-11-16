package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sso_listAccountsCmd = &cobra.Command{
	Use:   "list-accounts",
	Short: "Lists all AWS accounts assigned to the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sso_listAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sso_listAccountsCmd).Standalone()

		sso_listAccountsCmd.Flags().String("access-token", "", "The token issued by the `CreateToken` API call.")
		sso_listAccountsCmd.Flags().String("max-results", "", "This is the number of items clients can request per page.")
		sso_listAccountsCmd.Flags().String("next-token", "", "(Optional) When requesting subsequent pages, this is the page token from the previous response output.")
		sso_listAccountsCmd.MarkFlagRequired("access-token")
	})
	ssoCmd.AddCommand(sso_listAccountsCmd)
}
