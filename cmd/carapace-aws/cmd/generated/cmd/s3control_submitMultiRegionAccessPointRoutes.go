package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_submitMultiRegionAccessPointRoutesCmd = &cobra.Command{
	Use:   "submit-multi-region-access-point-routes",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_submitMultiRegionAccessPointRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_submitMultiRegionAccessPointRoutesCmd).Standalone()

		s3control_submitMultiRegionAccessPointRoutesCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
		s3control_submitMultiRegionAccessPointRoutesCmd.Flags().String("mrap", "", "The Multi-Region Access Point ARN.")
		s3control_submitMultiRegionAccessPointRoutesCmd.Flags().String("route-updates", "", "The different routes that make up the new route configuration.")
		s3control_submitMultiRegionAccessPointRoutesCmd.MarkFlagRequired("account-id")
		s3control_submitMultiRegionAccessPointRoutesCmd.MarkFlagRequired("mrap")
		s3control_submitMultiRegionAccessPointRoutesCmd.MarkFlagRequired("route-updates")
	})
	s3controlCmd.AddCommand(s3control_submitMultiRegionAccessPointRoutesCmd)
}
