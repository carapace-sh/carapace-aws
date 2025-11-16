package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listTrafficPoliciesCmd = &cobra.Command{
	Use:   "list-traffic-policies",
	Short: "List traffic policy resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listTrafficPoliciesCmd).Standalone()

	mailmanager_listTrafficPoliciesCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
	mailmanager_listTrafficPoliciesCmd.Flags().String("page-size", "", "The maximum number of traffic policy resources that are returned per call.")
	mailmanagerCmd.AddCommand(mailmanager_listTrafficPoliciesCmd)
}
