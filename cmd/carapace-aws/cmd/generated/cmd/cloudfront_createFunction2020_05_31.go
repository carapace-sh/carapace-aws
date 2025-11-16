package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createFunction2020_05_31Cmd = &cobra.Command{
	Use:   "create-function2020_05_31",
	Short: "Creates a CloudFront function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createFunction2020_05_31Cmd).Standalone()

	cloudfront_createFunction2020_05_31Cmd.Flags().String("function-code", "", "The function code.")
	cloudfront_createFunction2020_05_31Cmd.Flags().String("function-config", "", "Configuration information about the function, including an optional comment and the function's runtime.")
	cloudfront_createFunction2020_05_31Cmd.Flags().String("name", "", "A name to identify the function.")
	cloudfront_createFunction2020_05_31Cmd.MarkFlagRequired("function-code")
	cloudfront_createFunction2020_05_31Cmd.MarkFlagRequired("function-config")
	cloudfront_createFunction2020_05_31Cmd.MarkFlagRequired("name")
	cloudfrontCmd.AddCommand(cloudfront_createFunction2020_05_31Cmd)
}
