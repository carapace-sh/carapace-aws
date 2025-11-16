package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_putPolicyCmd = &cobra.Command{
	Use:   "put-policy",
	Short: "Attaches a resource-based policy to a private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_putPolicyCmd).Standalone()

	acmPca_putPolicyCmd.Flags().String("policy", "", "The path and file name of a JSON-formatted IAM policy to attach to the specified private CA resource.")
	acmPca_putPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Number (ARN) of the private CA to associate with the policy.")
	acmPca_putPolicyCmd.MarkFlagRequired("policy")
	acmPca_putPolicyCmd.MarkFlagRequired("resource-arn")
	acmPcaCmd.AddCommand(acmPca_putPolicyCmd)
}
