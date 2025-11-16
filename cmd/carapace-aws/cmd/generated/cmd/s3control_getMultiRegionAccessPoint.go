package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getMultiRegionAccessPointCmd = &cobra.Command{
	Use:   "get-multi-region-access-point",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getMultiRegionAccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getMultiRegionAccessPointCmd).Standalone()

		s3control_getMultiRegionAccessPointCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
		s3control_getMultiRegionAccessPointCmd.Flags().String("name", "", "The name of the Multi-Region Access Point whose configuration information you want to receive.")
		s3control_getMultiRegionAccessPointCmd.MarkFlagRequired("account-id")
		s3control_getMultiRegionAccessPointCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getMultiRegionAccessPointCmd)
}
