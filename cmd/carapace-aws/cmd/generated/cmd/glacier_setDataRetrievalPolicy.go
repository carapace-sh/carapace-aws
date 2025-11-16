package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_setDataRetrievalPolicyCmd = &cobra.Command{
	Use:   "set-data-retrieval-policy",
	Short: "This operation sets and then enacts a data retrieval policy in the region specified in the PUT request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_setDataRetrievalPolicyCmd).Standalone()

	glacier_setDataRetrievalPolicyCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
	glacier_setDataRetrievalPolicyCmd.Flags().String("policy", "", "The data retrieval policy in JSON format.")
	glacier_setDataRetrievalPolicyCmd.MarkFlagRequired("account-id")
	glacierCmd.AddCommand(glacier_setDataRetrievalPolicyCmd)
}
