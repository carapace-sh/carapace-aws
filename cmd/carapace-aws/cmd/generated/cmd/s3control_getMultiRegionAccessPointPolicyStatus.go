package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getMultiRegionAccessPointPolicyStatusCmd = &cobra.Command{
	Use:   "get-multi-region-access-point-policy-status",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getMultiRegionAccessPointPolicyStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getMultiRegionAccessPointPolicyStatusCmd).Standalone()

		s3control_getMultiRegionAccessPointPolicyStatusCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
		s3control_getMultiRegionAccessPointPolicyStatusCmd.Flags().String("name", "", "Specifies the Multi-Region Access Point.")
		s3control_getMultiRegionAccessPointPolicyStatusCmd.MarkFlagRequired("account-id")
		s3control_getMultiRegionAccessPointPolicyStatusCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getMultiRegionAccessPointPolicyStatusCmd)
}
