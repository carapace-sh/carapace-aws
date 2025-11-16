package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteIndexCmd = &cobra.Command{
	Use:   "delete-index",
	Short: "Deletes an Amazon Q Business index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deleteIndexCmd).Standalone()

		qbusiness_deleteIndexCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application the Amazon Q Business index is linked to.")
		qbusiness_deleteIndexCmd.Flags().String("index-id", "", "The identifier of the Amazon Q Business index.")
		qbusiness_deleteIndexCmd.MarkFlagRequired("application-id")
		qbusiness_deleteIndexCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deleteIndexCmd)
}
