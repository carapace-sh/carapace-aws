package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createRetrieverCmd = &cobra.Command{
	Use:   "create-retriever",
	Short: "Adds a retriever to your Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createRetrieverCmd).Standalone()

	qbusiness_createRetrieverCmd.Flags().String("application-id", "", "The identifier of your Amazon Q Business application.")
	qbusiness_createRetrieverCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create your Amazon Q Business application retriever.")
	qbusiness_createRetrieverCmd.Flags().String("configuration", "", "")
	qbusiness_createRetrieverCmd.Flags().String("display-name", "", "The name of your retriever.")
	qbusiness_createRetrieverCmd.Flags().String("role-arn", "", "The ARN of an IAM role used by Amazon Q Business to access the basic authentication credentials stored in a Secrets Manager secret.")
	qbusiness_createRetrieverCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the retriever.")
	qbusiness_createRetrieverCmd.Flags().String("type", "", "The type of retriever you are using.")
	qbusiness_createRetrieverCmd.MarkFlagRequired("application-id")
	qbusiness_createRetrieverCmd.MarkFlagRequired("configuration")
	qbusiness_createRetrieverCmd.MarkFlagRequired("display-name")
	qbusiness_createRetrieverCmd.MarkFlagRequired("type")
	qbusinessCmd.AddCommand(qbusiness_createRetrieverCmd)
}
