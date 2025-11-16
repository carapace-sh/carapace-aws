package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Stores a resource policy for the ARN of a `Project` or `ReportGroup` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_putResourcePolicyCmd).Standalone()

		codebuild_putResourcePolicyCmd.Flags().String("policy", "", "A JSON-formatted resource policy.")
		codebuild_putResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the `Project` or `ReportGroup` resource you want to associate with a resource policy.")
		codebuild_putResourcePolicyCmd.MarkFlagRequired("policy")
		codebuild_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	codebuildCmd.AddCommand(codebuild_putResourcePolicyCmd)
}
