package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_copyDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "copy-distribution2020_05_31",
	Short: "Creates a staging distribution using the configuration of the provided primary distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_copyDistribution2020_05_31Cmd).Standalone()

	cloudfront_copyDistribution2020_05_31Cmd.Flags().String("caller-reference", "", "A value that uniquely identifies a request to create a resource.")
	cloudfront_copyDistribution2020_05_31Cmd.Flags().String("enabled", "", "A Boolean flag to specify the state of the staging distribution when it's created.")
	cloudfront_copyDistribution2020_05_31Cmd.Flags().String("if-match", "", "The version identifier of the primary distribution whose configuration you are copying.")
	cloudfront_copyDistribution2020_05_31Cmd.Flags().String("primary-distribution-id", "", "The identifier of the primary distribution whose configuration you are copying.")
	cloudfront_copyDistribution2020_05_31Cmd.Flags().String("staging", "", "The type of distribution that your primary distribution will be copied to.")
	cloudfront_copyDistribution2020_05_31Cmd.MarkFlagRequired("caller-reference")
	cloudfront_copyDistribution2020_05_31Cmd.MarkFlagRequired("primary-distribution-id")
	cloudfrontCmd.AddCommand(cloudfront_copyDistribution2020_05_31Cmd)
}
