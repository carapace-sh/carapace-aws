package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createAndAttachS3AccessPointCmd = &cobra.Command{
	Use:   "create-and-attach-s3-access-point",
	Short: "Creates an S3 access point and attaches it to an Amazon FSx volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createAndAttachS3AccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_createAndAttachS3AccessPointCmd).Standalone()

		fsx_createAndAttachS3AccessPointCmd.Flags().String("client-request-token", "", "")
		fsx_createAndAttachS3AccessPointCmd.Flags().String("name", "", "The name you want to assign to this S3 access point.")
		fsx_createAndAttachS3AccessPointCmd.Flags().String("open-zfsconfiguration", "", "Specifies the configuration to use when creating and attaching an S3 access point to an FSx for OpenZFS volume.")
		fsx_createAndAttachS3AccessPointCmd.Flags().String("s3-access-point", "", "Specifies the virtual private cloud (VPC) configuration if you're creating an access point that is restricted to a VPC.")
		fsx_createAndAttachS3AccessPointCmd.Flags().String("type", "", "The type of S3 access point you want to create.")
		fsx_createAndAttachS3AccessPointCmd.MarkFlagRequired("name")
		fsx_createAndAttachS3AccessPointCmd.MarkFlagRequired("type")
	})
	fsxCmd.AddCommand(fsx_createAndAttachS3AccessPointCmd)
}
