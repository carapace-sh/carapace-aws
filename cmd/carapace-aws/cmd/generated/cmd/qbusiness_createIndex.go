package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Creates an Amazon Q Business index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createIndexCmd).Standalone()

		qbusiness_createIndexCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application using the index.")
		qbusiness_createIndexCmd.Flags().String("capacity-configuration", "", "The capacity units you want to provision for your index.")
		qbusiness_createIndexCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create an index.")
		qbusiness_createIndexCmd.Flags().String("description", "", "A description for the Amazon Q Business index.")
		qbusiness_createIndexCmd.Flags().String("display-name", "", "A name for the Amazon Q Business index.")
		qbusiness_createIndexCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the index.")
		qbusiness_createIndexCmd.Flags().String("type", "", "The index type that's suitable for your needs.")
		qbusiness_createIndexCmd.MarkFlagRequired("application-id")
		qbusiness_createIndexCmd.MarkFlagRequired("display-name")
	})
	qbusinessCmd.AddCommand(qbusiness_createIndexCmd)
}
