package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listSecurityConfigsCmd = &cobra.Command{
	Use:   "list-security-configs",
	Short: "Returns information about configured OpenSearch Serverless security configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listSecurityConfigsCmd).Standalone()

	opensearchserverless_listSecurityConfigsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearchserverless_listSecurityConfigsCmd.Flags().String("next-token", "", "If your initial `ListSecurityConfigs` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListSecurityConfigs` operations, which returns results in the next page.")
	opensearchserverless_listSecurityConfigsCmd.Flags().String("type", "", "The type of security configuration.")
	opensearchserverless_listSecurityConfigsCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_listSecurityConfigsCmd)
}
