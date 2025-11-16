package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Updates the resource policy for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_putResourcePolicyCmd).Standalone()

		redshift_putResourcePolicyCmd.Flags().String("policy", "", "The content of the resource policy being updated.")
		redshift_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource of which its resource policy is updated.")
		redshift_putResourcePolicyCmd.MarkFlagRequired("policy")
		redshift_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	redshiftCmd.AddCommand(redshift_putResourcePolicyCmd)
}
