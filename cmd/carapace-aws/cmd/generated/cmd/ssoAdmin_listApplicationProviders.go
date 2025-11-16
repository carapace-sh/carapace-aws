package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationProvidersCmd = &cobra.Command{
	Use:   "list-application-providers",
	Short: "Lists the application providers configured in the IAM Identity Center identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationProvidersCmd).Standalone()

	ssoAdmin_listApplicationProvidersCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
	ssoAdmin_listApplicationProvidersCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationProvidersCmd)
}
