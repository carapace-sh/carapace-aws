package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listMemberAccountsCmd = &cobra.Command{
	Use:   "list-member-accounts",
	Short: "Returns a `MemberAccounts` object that lists the member accounts in the administrator's Amazon Web Services organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listMemberAccountsCmd).Standalone()

	fms_listMemberAccountsCmd.Flags().String("max-results", "", "Specifies the number of member account IDs that you want Firewall Manager to return for this request.")
	fms_listMemberAccountsCmd.Flags().String("next-token", "", "If you specify a value for `MaxResults` and you have more account IDs than the number that you specify for `MaxResults`, Firewall Manager returns a `NextToken` value in the response that allows you to list another group of IDs.")
	fmsCmd.AddCommand(fms_listMemberAccountsCmd)
}
