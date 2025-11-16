package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Get the resource policy for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_getResourcePolicyCmd).Standalone()

		redshift_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource of which its resource policy is fetched.")
		redshift_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	redshiftCmd.AddCommand(redshift_getResourcePolicyCmd)
}
