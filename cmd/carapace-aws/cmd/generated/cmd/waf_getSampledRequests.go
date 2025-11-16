package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getSampledRequestsCmd = &cobra.Command{
	Use:   "get-sampled-requests",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getSampledRequestsCmd).Standalone()

	waf_getSampledRequestsCmd.Flags().String("max-items", "", "The number of requests that you want AWS WAF to return from among the first 5,000 requests that your AWS resource received during the time range.")
	waf_getSampledRequestsCmd.Flags().String("rule-id", "", "`RuleId` is one of three values:")
	waf_getSampledRequestsCmd.Flags().String("time-window", "", "The start date and time and the end date and time of the range for which you want `GetSampledRequests` to return a sample of requests.")
	waf_getSampledRequestsCmd.Flags().String("web-acl-id", "", "The `WebACLId` of the `WebACL` for which you want `GetSampledRequests` to return a sample of requests.")
	waf_getSampledRequestsCmd.MarkFlagRequired("max-items")
	waf_getSampledRequestsCmd.MarkFlagRequired("rule-id")
	waf_getSampledRequestsCmd.MarkFlagRequired("time-window")
	waf_getSampledRequestsCmd.MarkFlagRequired("web-acl-id")
	wafCmd.AddCommand(waf_getSampledRequestsCmd)
}
