package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listQueryLoggingConfigsCmd = &cobra.Command{
	Use:   "list-query-logging-configs",
	Short: "Lists the configurations for DNS query logging that are associated with the current Amazon Web Services account or the configuration that is associated with a specified hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listQueryLoggingConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listQueryLoggingConfigsCmd).Standalone()

		route53_listQueryLoggingConfigsCmd.Flags().String("hosted-zone-id", "", "(Optional) If you want to list the query logging configuration that is associated with a hosted zone, specify the ID in `HostedZoneId`.")
		route53_listQueryLoggingConfigsCmd.Flags().String("max-results", "", "(Optional) The maximum number of query logging configurations that you want Amazon Route 53 to return in response to the current request.")
		route53_listQueryLoggingConfigsCmd.Flags().String("next-token", "", "(Optional) If the current Amazon Web Services account has more than `MaxResults` query logging configurations, use `NextToken` to get the second and subsequent pages of results.")
	})
	route53Cmd.AddCommand(route53_listQueryLoggingConfigsCmd)
}
