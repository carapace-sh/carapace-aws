package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getMultiRegionAccessPointRoutesCmd = &cobra.Command{
	Use:   "get-multi-region-access-point-routes",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getMultiRegionAccessPointRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getMultiRegionAccessPointRoutesCmd).Standalone()

		s3control_getMultiRegionAccessPointRoutesCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
		s3control_getMultiRegionAccessPointRoutesCmd.Flags().String("mrap", "", "The Multi-Region Access Point ARN.")
		s3control_getMultiRegionAccessPointRoutesCmd.MarkFlagRequired("account-id")
		s3control_getMultiRegionAccessPointRoutesCmd.MarkFlagRequired("mrap")
	})
	s3controlCmd.AddCommand(s3control_getMultiRegionAccessPointRoutesCmd)
}
