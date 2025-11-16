package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_getPolicyVersionCmd = &cobra.Command{
	Use:   "get-policy-version",
	Short: "Returns details for the version of a policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_getPolicyVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_getPolicyVersionCmd).Standalone()

		mpa_getPolicyVersionCmd.Flags().String("policy-version-arn", "", "Amazon Resource Name (ARN) for the policy.")
		mpa_getPolicyVersionCmd.MarkFlagRequired("policy-version-arn")
	})
	mpaCmd.AddCommand(mpa_getPolicyVersionCmd)
}
