package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessGrantsInstanceCmd = &cobra.Command{
	Use:   "get-access-grants-instance",
	Short: "Retrieves the S3 Access Grants instance for a Region in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessGrantsInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessGrantsInstanceCmd).Standalone()

		s3control_getAccessGrantsInstanceCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_getAccessGrantsInstanceCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_getAccessGrantsInstanceCmd)
}
