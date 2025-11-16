package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_getDataRetrievalPolicyCmd = &cobra.Command{
	Use:   "get-data-retrieval-policy",
	Short: "This operation returns the current data retrieval policy for the account and region specified in the GET request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_getDataRetrievalPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_getDataRetrievalPolicyCmd).Standalone()

		glacier_getDataRetrievalPolicyCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
		glacier_getDataRetrievalPolicyCmd.MarkFlagRequired("account-id")
	})
	glacierCmd.AddCommand(glacier_getDataRetrievalPolicyCmd)
}
