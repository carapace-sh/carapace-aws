package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Provides a list of groups that are mapped to users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listGroupsCmd).Standalone()

	qbusiness_listGroupsCmd.Flags().String("application-id", "", "The identifier of the application for getting a list of groups mapped to users.")
	qbusiness_listGroupsCmd.Flags().String("data-source-id", "", "The identifier of the data source for getting a list of groups mapped to users.")
	qbusiness_listGroupsCmd.Flags().String("index-id", "", "The identifier of the index for getting a list of groups mapped to users.")
	qbusiness_listGroupsCmd.Flags().String("max-results", "", "The maximum number of returned groups that are mapped to users.")
	qbusiness_listGroupsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Q Business returns a pagination token in the response.")
	qbusiness_listGroupsCmd.Flags().String("updated-earlier-than", "", "The timestamp identifier used for the latest `PUT` or `DELETE` action for mapping users to their groups.")
	qbusiness_listGroupsCmd.MarkFlagRequired("application-id")
	qbusiness_listGroupsCmd.MarkFlagRequired("index-id")
	qbusiness_listGroupsCmd.MarkFlagRequired("updated-earlier-than")
	qbusinessCmd.AddCommand(qbusiness_listGroupsCmd)
}
