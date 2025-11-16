package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_describeMultiRegionAccessPointOperationCmd = &cobra.Command{
	Use:   "describe-multi-region-access-point-operation",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_describeMultiRegionAccessPointOperationCmd).Standalone()

	s3control_describeMultiRegionAccessPointOperationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the Multi-Region Access Point.")
	s3control_describeMultiRegionAccessPointOperationCmd.Flags().String("request-token-arn", "", "The request token associated with the request you want to know about.")
	s3control_describeMultiRegionAccessPointOperationCmd.MarkFlagRequired("account-id")
	s3control_describeMultiRegionAccessPointOperationCmd.MarkFlagRequired("request-token-arn")
	s3controlCmd.AddCommand(s3control_describeMultiRegionAccessPointOperationCmd)
}
