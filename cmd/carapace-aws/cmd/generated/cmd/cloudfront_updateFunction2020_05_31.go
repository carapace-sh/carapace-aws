package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateFunction2020_05_31Cmd = &cobra.Command{
	Use:   "update-function2020_05_31",
	Short: "Updates a CloudFront function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateFunction2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updateFunction2020_05_31Cmd).Standalone()

		cloudfront_updateFunction2020_05_31Cmd.Flags().String("function-code", "", "The function code.")
		cloudfront_updateFunction2020_05_31Cmd.Flags().String("function-config", "", "Configuration information about the function.")
		cloudfront_updateFunction2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the function that you are updating, which you can get using `DescribeFunction`.")
		cloudfront_updateFunction2020_05_31Cmd.Flags().String("name", "", "The name of the function that you are updating.")
		cloudfront_updateFunction2020_05_31Cmd.MarkFlagRequired("function-code")
		cloudfront_updateFunction2020_05_31Cmd.MarkFlagRequired("function-config")
		cloudfront_updateFunction2020_05_31Cmd.MarkFlagRequired("if-match")
		cloudfront_updateFunction2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_updateFunction2020_05_31Cmd)
}
