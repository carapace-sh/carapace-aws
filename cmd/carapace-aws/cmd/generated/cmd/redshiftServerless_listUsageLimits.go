package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listUsageLimitsCmd = &cobra.Command{
	Use:   "list-usage-limits",
	Short: "Lists all usage limits within Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listUsageLimitsCmd).Standalone()

	redshiftServerless_listUsageLimitsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	redshiftServerless_listUsageLimitsCmd.Flags().String("next-token", "", "If your initial `ListUsageLimits` operation returns a `nextToken`, you can include the returned `nextToken` in following `ListUsageLimits` operations, which returns results in the next page.")
	redshiftServerless_listUsageLimitsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource whose usage limits you want to list.")
	redshiftServerless_listUsageLimitsCmd.Flags().String("usage-type", "", "The Amazon Redshift Serverless feature whose limits you want to see.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listUsageLimitsCmd)
}
