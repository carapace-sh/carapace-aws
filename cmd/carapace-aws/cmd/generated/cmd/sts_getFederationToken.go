package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_getFederationTokenCmd = &cobra.Command{
	Use:   "get-federation-token",
	Short: "Returns a set of temporary security credentials (consisting of an access key ID, a secret access key, and a security token) for a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_getFederationTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sts_getFederationTokenCmd).Standalone()

		sts_getFederationTokenCmd.Flags().String("duration-seconds", "", "The duration, in seconds, that the session should last.")
		sts_getFederationTokenCmd.Flags().String("name", "", "The name of the federated user.")
		sts_getFederationTokenCmd.Flags().String("policy", "", "An IAM policy in JSON format that you want to use as an inline session policy.")
		sts_getFederationTokenCmd.Flags().String("policy-arns", "", "The Amazon Resource Names (ARNs) of the IAM managed policies that you want to use as a managed session policy.")
		sts_getFederationTokenCmd.Flags().String("tags", "", "A list of session tags.")
		sts_getFederationTokenCmd.MarkFlagRequired("name")
	})
	stsCmd.AddCommand(sts_getFederationTokenCmd)
}
