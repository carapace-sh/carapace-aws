package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessGrantsInstanceForPrefixCmd = &cobra.Command{
	Use:   "get-access-grants-instance-for-prefix",
	Short: "Retrieve the S3 Access Grants instance that contains a particular prefix.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessGrantsInstanceForPrefixCmd).Standalone()

	s3control_getAccessGrantsInstanceForPrefixCmd.Flags().String("account-id", "", "The ID of the Amazon Web Services account that is making this request.")
	s3control_getAccessGrantsInstanceForPrefixCmd.Flags().String("s3-prefix", "", "The S3 prefix of the access grants that you would like to retrieve.")
	s3control_getAccessGrantsInstanceForPrefixCmd.MarkFlagRequired("account-id")
	s3control_getAccessGrantsInstanceForPrefixCmd.MarkFlagRequired("s3-prefix")
	s3controlCmd.AddCommand(s3control_getAccessGrantsInstanceForPrefixCmd)
}
