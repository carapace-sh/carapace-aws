package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_batchGetMemberAccountDetailsCmd = &cobra.Command{
	Use:   "batch-get-member-account-details",
	Short: "Provides information on whether the supplied account IDs are associated with a membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_batchGetMemberAccountDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_batchGetMemberAccountDetailsCmd).Standalone()

		securityIr_batchGetMemberAccountDetailsCmd.Flags().String("account-ids", "", "Optional element to query the membership relationship status to a provided list of account IDs.")
		securityIr_batchGetMemberAccountDetailsCmd.Flags().String("membership-id", "", "Required element used in combination with BatchGetMemberAccountDetails to identify the membership ID to query.")
		securityIr_batchGetMemberAccountDetailsCmd.MarkFlagRequired("account-ids")
		securityIr_batchGetMemberAccountDetailsCmd.MarkFlagRequired("membership-id")
	})
	securityIrCmd.AddCommand(securityIr_batchGetMemberAccountDetailsCmd)
}
