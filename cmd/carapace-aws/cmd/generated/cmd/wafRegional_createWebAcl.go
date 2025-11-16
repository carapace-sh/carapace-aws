package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createWebAclCmd = &cobra.Command{
	Use:   "create-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createWebAclCmd).Standalone()

	wafRegional_createWebAclCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_createWebAclCmd.Flags().String("default-action", "", "The action that you want AWS WAF to take when a request doesn't match the criteria specified in any of the `Rule` objects that are associated with the `WebACL`.")
	wafRegional_createWebAclCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `WebACL`.The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one.")
	wafRegional_createWebAclCmd.Flags().String("name", "", "A friendly name or description of the [WebACL]().")
	wafRegional_createWebAclCmd.Flags().String("tags", "", "")
	wafRegional_createWebAclCmd.MarkFlagRequired("change-token")
	wafRegional_createWebAclCmd.MarkFlagRequired("default-action")
	wafRegional_createWebAclCmd.MarkFlagRequired("metric-name")
	wafRegional_createWebAclCmd.MarkFlagRequired("name")
	wafRegionalCmd.AddCommand(wafRegional_createWebAclCmd)
}
