package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteMultiRegionAccessPointCmd = &cobra.Command{
	Use:   "delete-multi-region-access-point",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteMultiRegionAccessPointCmd).Standalone()

	s3control_deleteMultiRegionAccessPointCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
	s3control_deleteMultiRegionAccessPointCmd.Flags().String("client-token", "", "An idempotency token used to identify the request and guarantee that requests are unique.")
	s3control_deleteMultiRegionAccessPointCmd.Flags().String("details", "", "A container element containing details about the Multi-Region Access Point.")
	s3control_deleteMultiRegionAccessPointCmd.MarkFlagRequired("account-id")
	s3control_deleteMultiRegionAccessPointCmd.MarkFlagRequired("client-token")
	s3control_deleteMultiRegionAccessPointCmd.MarkFlagRequired("details")
	s3controlCmd.AddCommand(s3control_deleteMultiRegionAccessPointCmd)
}
