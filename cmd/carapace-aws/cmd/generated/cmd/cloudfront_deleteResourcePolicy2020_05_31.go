package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteResourcePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "delete-resource-policy2020_05_31",
	Short: "Deletes the resource policy attached to the CloudFront resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteResourcePolicy2020_05_31Cmd).Standalone()

	cloudfront_deleteResourcePolicy2020_05_31Cmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudFront resource for which the resource policy should be deleted.")
	cloudfront_deleteResourcePolicy2020_05_31Cmd.MarkFlagRequired("resource-arn")
	cloudfrontCmd.AddCommand(cloudfront_deleteResourcePolicy2020_05_31Cmd)
}
