package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listIndicesCmd = &cobra.Command{
	Use:   "list-indices",
	Short: "Lists the Amazon Q Business indices you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listIndicesCmd).Standalone()

	qbusiness_listIndicesCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application connected to the index.")
	qbusiness_listIndicesCmd.Flags().String("max-results", "", "The maximum number of indices to return.")
	qbusiness_listIndicesCmd.Flags().String("next-token", "", "If the maxResults response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
	qbusiness_listIndicesCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_listIndicesCmd)
}
