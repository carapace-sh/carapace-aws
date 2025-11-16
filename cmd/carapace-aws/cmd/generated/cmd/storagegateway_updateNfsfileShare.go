package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateNfsfileShareCmd = &cobra.Command{
	Use:   "update-nfsfile-share",
	Short: "Updates a Network File System (NFS) file share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateNfsfileShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateNfsfileShareCmd).Standalone()

		storagegateway_updateNfsfileShareCmd.Flags().String("audit-destination-arn", "", "The Amazon Resource Name (ARN) of the storage used for audit logs.")
		storagegateway_updateNfsfileShareCmd.Flags().String("cache-attributes", "", "Specifies refresh cache information for the file share.")
		storagegateway_updateNfsfileShareCmd.Flags().String("client-list", "", "The list of clients that are allowed to access the S3 File Gateway.")
		storagegateway_updateNfsfileShareCmd.Flags().String("default-storage-class", "", "The default storage class for objects put into an Amazon S3 bucket by the S3 File Gateway.")
		storagegateway_updateNfsfileShareCmd.Flags().String("encryption-type", "", "A value that specifies the type of server-side encryption that the file share will use for the data that it stores in Amazon S3.")
		storagegateway_updateNfsfileShareCmd.Flags().String("file-share-arn", "", "The Amazon Resource Name (ARN) of the file share to be updated.")
		storagegateway_updateNfsfileShareCmd.Flags().String("file-share-name", "", "The name of the file share.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("kmsencrypted", false, "Optional.")
		storagegateway_updateNfsfileShareCmd.Flags().String("kmskey", "", "Optional.")
		storagegateway_updateNfsfileShareCmd.Flags().String("nfsfile-share-defaults", "", "The default values for the file share.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("no-guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("no-kmsencrypted", false, "Optional.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("no-read-only", false, "A value that sets the write status of a file share.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("no-requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
		storagegateway_updateNfsfileShareCmd.Flags().String("notification-policy", "", "The notification policy of the file share.")
		storagegateway_updateNfsfileShareCmd.Flags().String("object-acl", "", "A value that sets the access control list (ACL) permission for objects in the S3 bucket that a S3 File Gateway puts objects into.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("read-only", false, "A value that sets the write status of a file share.")
		storagegateway_updateNfsfileShareCmd.Flags().Bool("requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
		storagegateway_updateNfsfileShareCmd.Flags().String("squash", "", "The user mapped to anonymous user.")
		storagegateway_updateNfsfileShareCmd.MarkFlagRequired("file-share-arn")
		storagegateway_updateNfsfileShareCmd.Flag("no-guess-mimetype-enabled").Hidden = true
		storagegateway_updateNfsfileShareCmd.Flag("no-kmsencrypted").Hidden = true
		storagegateway_updateNfsfileShareCmd.Flag("no-read-only").Hidden = true
		storagegateway_updateNfsfileShareCmd.Flag("no-requester-pays").Hidden = true
	})
	storagegatewayCmd.AddCommand(storagegateway_updateNfsfileShareCmd)
}
