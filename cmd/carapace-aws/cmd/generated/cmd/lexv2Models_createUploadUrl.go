package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createUploadUrlCmd = &cobra.Command{
	Use:   "create-upload-url",
	Short: "Gets a pre-signed S3 write URL that you use to upload the zip archive when importing a bot or a bot locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createUploadUrlCmd).Standalone()

	lexv2ModelsCmd.AddCommand(lexv2Models_createUploadUrlCmd)
}
