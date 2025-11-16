package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_createAssetCmd = &cobra.Command{
	Use:   "create-asset",
	Short: "Creates a new MediaPackage VOD Asset resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_createAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_createAssetCmd).Standalone()

		mediapackageVod_createAssetCmd.Flags().String("id", "", "The unique identifier for the Asset.")
		mediapackageVod_createAssetCmd.Flags().String("packaging-group-id", "", "The ID of the PackagingGroup for the Asset.")
		mediapackageVod_createAssetCmd.Flags().String("resource-id", "", "The resource ID to include in SPEKE key requests.")
		mediapackageVod_createAssetCmd.Flags().String("source-arn", "", "ARN of the source object in S3.")
		mediapackageVod_createAssetCmd.Flags().String("source-role-arn", "", "The IAM role ARN used to access the source S3 bucket.")
		mediapackageVod_createAssetCmd.Flags().String("tags", "", "")
		mediapackageVod_createAssetCmd.MarkFlagRequired("id")
		mediapackageVod_createAssetCmd.MarkFlagRequired("packaging-group-id")
		mediapackageVod_createAssetCmd.MarkFlagRequired("source-arn")
		mediapackageVod_createAssetCmd.MarkFlagRequired("source-role-arn")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_createAssetCmd)
}
