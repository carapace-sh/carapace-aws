package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "Lists the details of the organization and account instances of IAM Identity Center that were created in or visible to the account calling this API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listInstancesCmd).Standalone()

		ssoAdmin_listInstancesCmd.Flags().String("max-results", "", "The maximum number of results to display for the instance.")
		ssoAdmin_listInstancesCmd.Flags().String("next-token", "", "The pagination token for the list API.")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listInstancesCmd)
}
