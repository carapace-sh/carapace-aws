package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd = &cobra.Command{
	Use:   "update-distribution-with-staging-config2020_05_31",
	Short: "Copies the staging distribution's configuration to its corresponding primary distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd).Standalone()

	cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd.Flags().String("id", "", "The identifier of the primary distribution to which you are copying a staging distribution's configuration.")
	cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd.Flags().String("if-match", "", "The current versions (`ETag` values) of both primary and staging distributions.")
	cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd.Flags().String("staging-distribution-id", "", "The identifier of the staging distribution whose configuration you are copying to the primary distribution.")
	cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_updateDistributionWithStagingConfig2020_05_31Cmd)
}
