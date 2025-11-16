package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateRetrieverCmd = &cobra.Command{
	Use:   "update-retriever",
	Short: "Updates the retriever used for your Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateRetrieverCmd).Standalone()

	qbusiness_updateRetrieverCmd.Flags().String("application-id", "", "The identifier of your Amazon Q Business application.")
	qbusiness_updateRetrieverCmd.Flags().String("configuration", "", "")
	qbusiness_updateRetrieverCmd.Flags().String("display-name", "", "The name of your retriever.")
	qbusiness_updateRetrieverCmd.Flags().String("retriever-id", "", "The identifier of your retriever.")
	qbusiness_updateRetrieverCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the retriever and required resources.")
	qbusiness_updateRetrieverCmd.MarkFlagRequired("application-id")
	qbusiness_updateRetrieverCmd.MarkFlagRequired("retriever-id")
	qbusinessCmd.AddCommand(qbusiness_updateRetrieverCmd)
}
