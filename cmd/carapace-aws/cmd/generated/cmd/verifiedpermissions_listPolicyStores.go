package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_listPolicyStoresCmd = &cobra.Command{
	Use:   "list-policy-stores",
	Short: "Returns a paginated list of all policy stores in the calling Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_listPolicyStoresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_listPolicyStoresCmd).Standalone()

		verifiedpermissions_listPolicyStoresCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		verifiedpermissions_listPolicyStoresCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_listPolicyStoresCmd)
}
