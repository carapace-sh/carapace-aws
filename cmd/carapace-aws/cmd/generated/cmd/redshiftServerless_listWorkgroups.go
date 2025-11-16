package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listWorkgroupsCmd = &cobra.Command{
	Use:   "list-workgroups",
	Short: "Returns information about a list of specified workgroups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listWorkgroupsCmd).Standalone()

	redshiftServerless_listWorkgroupsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	redshiftServerless_listWorkgroupsCmd.Flags().String("next-token", "", "If your initial ListWorkgroups operation returns a `nextToken`, you can include the returned `nextToken` in following ListNamespaces operations, which returns results in the next page.")
	redshiftServerless_listWorkgroupsCmd.Flags().String("owner-account", "", "The owner Amazon Web Services account for the Amazon Redshift Serverless workgroup.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listWorkgroupsCmd)
}
