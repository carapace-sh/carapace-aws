package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listPoliciesCmd = &cobra.Command{
	Use:   "list-policies",
	Short: "Returns a list of policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_listPoliciesCmd).Standalone()

		mpa_listPoliciesCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		mpa_listPoliciesCmd.Flags().String("next-token", "", "If present, indicates that more output is available than is included in the current response.")
	})
	mpaCmd.AddCommand(mpa_listPoliciesCmd)
}
