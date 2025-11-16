package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_assumeRootCmd = &cobra.Command{
	Use:   "assume-root",
	Short: "Returns a set of short term credentials you can use to perform privileged tasks on a member account in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_assumeRootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sts_assumeRootCmd).Standalone()

		sts_assumeRootCmd.Flags().String("duration-seconds", "", "The duration, in seconds, of the privileged session.")
		sts_assumeRootCmd.Flags().String("target-principal", "", "The member account principal ARN or account ID.")
		sts_assumeRootCmd.Flags().String("task-policy-arn", "", "The identity based policy that scopes the session to the privileged tasks that can be performed.")
		sts_assumeRootCmd.MarkFlagRequired("target-principal")
		sts_assumeRootCmd.MarkFlagRequired("task-policy-arn")
	})
	stsCmd.AddCommand(sts_assumeRootCmd)
}
