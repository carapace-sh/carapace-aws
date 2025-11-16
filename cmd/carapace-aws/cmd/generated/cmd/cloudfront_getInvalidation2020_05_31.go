package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getInvalidation2020_05_31Cmd = &cobra.Command{
	Use:   "get-invalidation2020_05_31",
	Short: "Get the information about an invalidation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getInvalidation2020_05_31Cmd).Standalone()

	cloudfront_getInvalidation2020_05_31Cmd.Flags().String("distribution-id", "", "The distribution's ID.")
	cloudfront_getInvalidation2020_05_31Cmd.Flags().String("id", "", "The identifier for the invalidation request, for example, `IDFDVBD632BHDS5`.")
	cloudfront_getInvalidation2020_05_31Cmd.MarkFlagRequired("distribution-id")
	cloudfront_getInvalidation2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getInvalidation2020_05_31Cmd)
}
