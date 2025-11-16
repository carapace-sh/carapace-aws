package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_putResourcePolicyCmd).Standalone()

	redshiftServerless_putResourcePolicyCmd.Flags().String("policy", "", "The policy to create or update.")
	redshiftServerless_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the account to create or update a resource policy for.")
	redshiftServerless_putResourcePolicyCmd.MarkFlagRequired("policy")
	redshiftServerless_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	redshiftServerlessCmd.AddCommand(redshiftServerless_putResourcePolicyCmd)
}
