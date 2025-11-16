package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_createUploadUrlCmd = &cobra.Command{
	Use:   "create-upload-url",
	Short: "Generates a pre-signed URL, request headers used to upload a code resource, and code artifact identifier for the uploaded resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_createUploadUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_createUploadUrlCmd).Standalone()

		codeguruSecurity_createUploadUrlCmd.Flags().String("scan-name", "", "The name of the scan that will use the uploaded resource.")
		codeguruSecurity_createUploadUrlCmd.MarkFlagRequired("scan-name")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_createUploadUrlCmd)
}
