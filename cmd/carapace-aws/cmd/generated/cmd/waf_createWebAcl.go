package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createWebAclCmd = &cobra.Command{
	Use:   "create-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createWebAclCmd).Standalone()

		waf_createWebAclCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createWebAclCmd.Flags().String("default-action", "", "The action that you want AWS WAF to take when a request doesn't match the criteria specified in any of the `Rule` objects that are associated with the `WebACL`.")
		waf_createWebAclCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `WebACL`.The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one.")
		waf_createWebAclCmd.Flags().String("name", "", "A friendly name or description of the [WebACL]().")
		waf_createWebAclCmd.Flags().String("tags", "", "")
		waf_createWebAclCmd.MarkFlagRequired("change-token")
		waf_createWebAclCmd.MarkFlagRequired("default-action")
		waf_createWebAclCmd.MarkFlagRequired("metric-name")
		waf_createWebAclCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createWebAclCmd)
}
