package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getEvidenceFileUploadUrlCmd = &cobra.Command{
	Use:   "get-evidence-file-upload-url",
	Short: "Creates a presigned Amazon S3 URL that can be used to upload a file as manual evidence.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getEvidenceFileUploadUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getEvidenceFileUploadUrlCmd).Standalone()

		auditmanager_getEvidenceFileUploadUrlCmd.Flags().String("file-name", "", "The file that you want to upload.")
		auditmanager_getEvidenceFileUploadUrlCmd.MarkFlagRequired("file-name")
	})
	auditmanagerCmd.AddCommand(auditmanager_getEvidenceFileUploadUrlCmd)
}
