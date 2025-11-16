package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the specified resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_deleteResourcePolicyCmd).Standalone()

		redshiftServerless_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the policy to delete.")
		redshiftServerless_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteResourcePolicyCmd)
}
