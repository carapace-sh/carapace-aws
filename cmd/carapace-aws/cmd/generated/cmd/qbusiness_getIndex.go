package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getIndexCmd = &cobra.Command{
	Use:   "get-index",
	Short: "Gets information about an existing Amazon Q Business index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getIndexCmd).Standalone()

		qbusiness_getIndexCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application connected to the index.")
		qbusiness_getIndexCmd.Flags().String("index-id", "", "The identifier of the Amazon Q Business index you want information on.")
		qbusiness_getIndexCmd.MarkFlagRequired("application-id")
		qbusiness_getIndexCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getIndexCmd)
}
