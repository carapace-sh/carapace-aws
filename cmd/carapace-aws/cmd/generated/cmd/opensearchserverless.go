package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverlessCmd = &cobra.Command{
	Use:   "opensearchserverless",
	Short: "Use the Amazon OpenSearch Serverless API to create, configure, and manage OpenSearch Serverless collections and security policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverlessCmd).Standalone()

	rootCmd.AddCommand(opensearchserverlessCmd)
}
