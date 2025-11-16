package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createInvalidation2020_05_31Cmd = &cobra.Command{
	Use:   "create-invalidation2020_05_31",
	Short: "Create a new invalidation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createInvalidation2020_05_31Cmd).Standalone()

	cloudfront_createInvalidation2020_05_31Cmd.Flags().String("distribution-id", "", "The distribution's id.")
	cloudfront_createInvalidation2020_05_31Cmd.Flags().String("invalidation-batch", "", "The batch information for the invalidation.")
	cloudfront_createInvalidation2020_05_31Cmd.MarkFlagRequired("distribution-id")
	cloudfront_createInvalidation2020_05_31Cmd.MarkFlagRequired("invalidation-batch")
	cloudfrontCmd.AddCommand(cloudfront_createInvalidation2020_05_31Cmd)
}
