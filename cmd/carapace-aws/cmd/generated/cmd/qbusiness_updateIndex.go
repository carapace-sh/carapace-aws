package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateIndexCmd = &cobra.Command{
	Use:   "update-index",
	Short: "Updates an Amazon Q Business index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_updateIndexCmd).Standalone()

		qbusiness_updateIndexCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application connected to the index.")
		qbusiness_updateIndexCmd.Flags().String("capacity-configuration", "", "The storage capacity units you want to provision for your Amazon Q Business index.")
		qbusiness_updateIndexCmd.Flags().String("description", "", "The description of the Amazon Q Business index.")
		qbusiness_updateIndexCmd.Flags().String("display-name", "", "The name of the Amazon Q Business index.")
		qbusiness_updateIndexCmd.Flags().String("document-attribute-configurations", "", "Configuration information for document metadata or fields.")
		qbusiness_updateIndexCmd.Flags().String("index-id", "", "The identifier of the Amazon Q Business index.")
		qbusiness_updateIndexCmd.MarkFlagRequired("application-id")
		qbusiness_updateIndexCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_updateIndexCmd)
}
