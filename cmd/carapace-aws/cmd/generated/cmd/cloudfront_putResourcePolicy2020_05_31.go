package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_putResourcePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "put-resource-policy2020_05_31",
	Short: "Creates a resource control policy for a given CloudFront resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_putResourcePolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_putResourcePolicy2020_05_31Cmd).Standalone()

		cloudfront_putResourcePolicy2020_05_31Cmd.Flags().String("policy-document", "", "The JSON-formatted resource policy to create.")
		cloudfront_putResourcePolicy2020_05_31Cmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudFront resource for which the policy is being created.")
		cloudfront_putResourcePolicy2020_05_31Cmd.MarkFlagRequired("policy-document")
		cloudfront_putResourcePolicy2020_05_31Cmd.MarkFlagRequired("resource-arn")
	})
	cloudfrontCmd.AddCommand(cloudfront_putResourcePolicy2020_05_31Cmd)
}
