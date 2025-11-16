package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putMultiRegionAccessPointPolicyCmd = &cobra.Command{
	Use:   "put-multi-region-access-point-policy",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putMultiRegionAccessPointPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putMultiRegionAccessPointPolicyCmd).Standalone()

		s3control_putMultiRegionAccessPointPolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
		s3control_putMultiRegionAccessPointPolicyCmd.Flags().String("client-token", "", "An idempotency token used to identify the request and guarantee that requests are unique.")
		s3control_putMultiRegionAccessPointPolicyCmd.Flags().String("details", "", "A container element containing the details of the policy for the Multi-Region Access Point.")
		s3control_putMultiRegionAccessPointPolicyCmd.MarkFlagRequired("account-id")
		s3control_putMultiRegionAccessPointPolicyCmd.MarkFlagRequired("client-token")
		s3control_putMultiRegionAccessPointPolicyCmd.MarkFlagRequired("details")
	})
	s3controlCmd.AddCommand(s3control_putMultiRegionAccessPointPolicyCmd)
}
