package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource policy for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteResourcePolicyCmd).Standalone()

		redshift_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource of which its resource policy is deleted.")
		redshift_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	redshiftCmd.AddCommand(redshift_deleteResourcePolicyCmd)
}
