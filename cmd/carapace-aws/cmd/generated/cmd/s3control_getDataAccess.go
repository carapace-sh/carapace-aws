package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getDataAccessCmd = &cobra.Command{
	Use:   "get-data-access",
	Short: "Returns a temporary access credential from S3 Access Grants to the grantee or client application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getDataAccessCmd).Standalone()

	s3control_getDataAccessCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_getDataAccessCmd.Flags().String("duration-seconds", "", "The session duration, in seconds, of the temporary access credential that S3 Access Grants vends to the grantee or client application.")
	s3control_getDataAccessCmd.Flags().String("permission", "", "The type of permission granted to your S3 data, which can be set to one of the following values:")
	s3control_getDataAccessCmd.Flags().String("privilege", "", "The scope of the temporary access credential that S3 Access Grants vends to the grantee or client application.")
	s3control_getDataAccessCmd.Flags().String("target", "", "The S3 URI path of the data to which you are requesting temporary access credentials.")
	s3control_getDataAccessCmd.Flags().String("target-type", "", "The type of `Target`.")
	s3control_getDataAccessCmd.MarkFlagRequired("account-id")
	s3control_getDataAccessCmd.MarkFlagRequired("permission")
	s3control_getDataAccessCmd.MarkFlagRequired("target")
	s3controlCmd.AddCommand(s3control_getDataAccessCmd)
}
