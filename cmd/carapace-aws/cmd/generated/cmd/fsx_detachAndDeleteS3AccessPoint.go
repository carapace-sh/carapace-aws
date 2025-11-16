package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_detachAndDeleteS3AccessPointCmd = &cobra.Command{
	Use:   "detach-and-delete-s3-access-point",
	Short: "Detaches an S3 access point from an Amazon FSx volume and deletes the S3 access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_detachAndDeleteS3AccessPointCmd).Standalone()

	fsx_detachAndDeleteS3AccessPointCmd.Flags().String("client-request-token", "", "")
	fsx_detachAndDeleteS3AccessPointCmd.Flags().String("name", "", "The name of the S3 access point attachment that you want to delete.")
	fsx_detachAndDeleteS3AccessPointCmd.MarkFlagRequired("name")
	fsxCmd.AddCommand(fsx_detachAndDeleteS3AccessPointCmd)
}
