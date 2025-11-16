package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listFunctions2020_05_31Cmd = &cobra.Command{
	Use:   "list-functions2020_05_31",
	Short: "Gets a list of all CloudFront functions in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listFunctions2020_05_31Cmd).Standalone()

	cloudfront_listFunctions2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of functions.")
	cloudfront_listFunctions2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of functions that you want in the response.")
	cloudfront_listFunctions2020_05_31Cmd.Flags().String("stage", "", "An optional filter to return only the functions that are in the specified stage, either `DEVELOPMENT` or `LIVE`.")
	cloudfrontCmd.AddCommand(cloudfront_listFunctions2020_05_31Cmd)
}
