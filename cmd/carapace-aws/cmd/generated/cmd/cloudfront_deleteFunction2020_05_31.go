package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteFunction2020_05_31Cmd = &cobra.Command{
	Use:   "delete-function2020_05_31",
	Short: "Deletes a CloudFront function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteFunction2020_05_31Cmd).Standalone()

	cloudfront_deleteFunction2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the function that you are deleting, which you can get using `DescribeFunction`.")
	cloudfront_deleteFunction2020_05_31Cmd.Flags().String("name", "", "The name of the function that you are deleting.")
	cloudfront_deleteFunction2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfront_deleteFunction2020_05_31Cmd.MarkFlagRequired("name")
	cloudfrontCmd.AddCommand(cloudfront_deleteFunction2020_05_31Cmd)
}
