package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getSampledRequestsCmd = &cobra.Command{
	Use:   "get-sampled-requests",
	Short: "Gets detailed information about a specified number of requests--a sample--that WAF randomly selects from among the first 5,000 requests that your Amazon Web Services resource received during a time range that you choose.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getSampledRequestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_getSampledRequestsCmd).Standalone()

		wafv2_getSampledRequestsCmd.Flags().String("max-items", "", "The number of requests that you want WAF to return from among the first 5,000 requests that your Amazon Web Services resource received during the time range.")
		wafv2_getSampledRequestsCmd.Flags().String("rule-metric-name", "", "The metric name assigned to the `Rule` or `RuleGroup` dimension for which you want a sample of requests.")
		wafv2_getSampledRequestsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_getSampledRequestsCmd.Flags().String("time-window", "", "The start date and time and the end date and time of the range for which you want `GetSampledRequests` to return a sample of requests.")
		wafv2_getSampledRequestsCmd.Flags().String("web-acl-arn", "", "The Amazon resource name (ARN) of the `WebACL` for which you want a sample of requests.")
		wafv2_getSampledRequestsCmd.MarkFlagRequired("max-items")
		wafv2_getSampledRequestsCmd.MarkFlagRequired("rule-metric-name")
		wafv2_getSampledRequestsCmd.MarkFlagRequired("scope")
		wafv2_getSampledRequestsCmd.MarkFlagRequired("time-window")
		wafv2_getSampledRequestsCmd.MarkFlagRequired("web-acl-arn")
	})
	wafv2Cmd.AddCommand(wafv2_getSampledRequestsCmd)
}
