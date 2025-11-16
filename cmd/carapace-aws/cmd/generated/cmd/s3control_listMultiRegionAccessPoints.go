package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listMultiRegionAccessPointsCmd = &cobra.Command{
	Use:   "list-multi-region-access-points",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listMultiRegionAccessPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listMultiRegionAccessPointsCmd).Standalone()

		s3control_listMultiRegionAccessPointsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
		s3control_listMultiRegionAccessPointsCmd.Flags().String("max-results", "", "Not currently used.")
		s3control_listMultiRegionAccessPointsCmd.Flags().String("next-token", "", "Not currently used.")
		s3control_listMultiRegionAccessPointsCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listMultiRegionAccessPointsCmd)
}
