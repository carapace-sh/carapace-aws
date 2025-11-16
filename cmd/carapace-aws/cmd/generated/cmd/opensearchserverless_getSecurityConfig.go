package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_getSecurityConfigCmd = &cobra.Command{
	Use:   "get-security-config",
	Short: "Returns information about an OpenSearch Serverless security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_getSecurityConfigCmd).Standalone()

	opensearchserverless_getSecurityConfigCmd.Flags().String("id", "", "The unique identifier of the security configuration.")
	opensearchserverless_getSecurityConfigCmd.MarkFlagRequired("id")
	opensearchserverlessCmd.AddCommand(opensearchserverless_getSecurityConfigCmd)
}
