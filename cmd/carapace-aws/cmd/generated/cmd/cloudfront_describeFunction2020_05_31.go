package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_describeFunction2020_05_31Cmd = &cobra.Command{
	Use:   "describe-function2020_05_31",
	Short: "Gets configuration information and metadata about a CloudFront function, but not the function's code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_describeFunction2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_describeFunction2020_05_31Cmd).Standalone()

		cloudfront_describeFunction2020_05_31Cmd.Flags().String("name", "", "The name of the function that you are getting information about.")
		cloudfront_describeFunction2020_05_31Cmd.Flags().String("stage", "", "The function's stage, either `DEVELOPMENT` or `LIVE`.")
		cloudfront_describeFunction2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_describeFunction2020_05_31Cmd)
}
