package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Returns a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_getResourcePolicyCmd).Standalone()

		redshiftServerless_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to return.")
		redshiftServerless_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_getResourcePolicyCmd)
}
