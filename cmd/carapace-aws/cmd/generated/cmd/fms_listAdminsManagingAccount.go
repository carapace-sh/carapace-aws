package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listAdminsManagingAccountCmd = &cobra.Command{
	Use:   "list-admins-managing-account",
	Short: "Lists the accounts that are managing the specified Organizations member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listAdminsManagingAccountCmd).Standalone()

	fms_listAdminsManagingAccountCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
	fms_listAdminsManagingAccountCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Firewall Manager returns a `NextToken` value in the response.")
	fmsCmd.AddCommand(fms_listAdminsManagingAccountCmd)
}
