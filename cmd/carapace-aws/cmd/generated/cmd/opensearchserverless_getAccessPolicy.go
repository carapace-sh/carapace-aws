package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_getAccessPolicyCmd = &cobra.Command{
	Use:   "get-access-policy",
	Short: "Returns an OpenSearch Serverless access policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_getAccessPolicyCmd).Standalone()

	opensearchserverless_getAccessPolicyCmd.Flags().String("name", "", "The name of the access policy.")
	opensearchserverless_getAccessPolicyCmd.Flags().String("type", "", "Tye type of policy.")
	opensearchserverless_getAccessPolicyCmd.MarkFlagRequired("name")
	opensearchserverless_getAccessPolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_getAccessPolicyCmd)
}
