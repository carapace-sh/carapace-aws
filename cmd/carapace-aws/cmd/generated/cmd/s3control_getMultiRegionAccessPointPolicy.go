package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getMultiRegionAccessPointPolicyCmd = &cobra.Command{
	Use:   "get-multi-region-access-point-policy",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getMultiRegionAccessPointPolicyCmd).Standalone()

	s3control_getMultiRegionAccessPointPolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
	s3control_getMultiRegionAccessPointPolicyCmd.Flags().String("name", "", "Specifies the Multi-Region Access Point.")
	s3control_getMultiRegionAccessPointPolicyCmd.MarkFlagRequired("account-id")
	s3control_getMultiRegionAccessPointPolicyCmd.MarkFlagRequired("name")
	s3controlCmd.AddCommand(s3control_getMultiRegionAccessPointPolicyCmd)
}
