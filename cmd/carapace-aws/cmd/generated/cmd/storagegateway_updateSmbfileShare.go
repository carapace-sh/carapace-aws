package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateSmbfileShareCmd = &cobra.Command{
	Use:   "update-smbfile-share",
	Short: "Updates a Server Message Block (SMB) file share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateSmbfileShareCmd).Standalone()

	storagegateway_updateSmbfileShareCmd.Flags().Bool("access-based-enumeration", false, "The files and folders on this share will only be visible to users with read access.")
	storagegateway_updateSmbfileShareCmd.Flags().String("admin-user-list", "", "A list of users or groups in the Active Directory that have administrator rights to the file share.")
	storagegateway_updateSmbfileShareCmd.Flags().String("audit-destination-arn", "", "The Amazon Resource Name (ARN) of the storage used for audit logs.")
	storagegateway_updateSmbfileShareCmd.Flags().String("cache-attributes", "", "Specifies refresh cache information for the file share.")
	storagegateway_updateSmbfileShareCmd.Flags().String("case-sensitivity", "", "The case of an object name in an Amazon S3 bucket.")
	storagegateway_updateSmbfileShareCmd.Flags().String("default-storage-class", "", "The default storage class for objects put into an Amazon S3 bucket by the S3 File Gateway.")
	storagegateway_updateSmbfileShareCmd.Flags().String("encryption-type", "", "A value that specifies the type of server-side encryption that the file share will use for the data that it stores in Amazon S3.")
	storagegateway_updateSmbfileShareCmd.Flags().String("file-share-arn", "", "The Amazon Resource Name (ARN) of the SMB file share that you want to update.")
	storagegateway_updateSmbfileShareCmd.Flags().String("file-share-name", "", "The name of the file share.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
	storagegateway_updateSmbfileShareCmd.Flags().String("invalid-user-list", "", "A list of users or groups in the Active Directory that are not allowed to access the file share.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("kmsencrypted", false, "Optional.")
	storagegateway_updateSmbfileShareCmd.Flags().String("kmskey", "", "Optional.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-access-based-enumeration", false, "The files and folders on this share will only be visible to users with read access.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-kmsencrypted", false, "Optional.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-oplocks-enabled", false, "Specifies whether opportunistic locking is enabled for the SMB file share.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-read-only", false, "A value that sets the write status of a file share.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("no-smbaclenabled", false, "Set this value to `true` to enable access control list (ACL) on the SMB file share.")
	storagegateway_updateSmbfileShareCmd.Flags().String("notification-policy", "", "The notification policy of the file share.")
	storagegateway_updateSmbfileShareCmd.Flags().String("object-acl", "", "A value that sets the access control list (ACL) permission for objects in the S3 bucket that a S3 File Gateway puts objects into.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("oplocks-enabled", false, "Specifies whether opportunistic locking is enabled for the SMB file share.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("read-only", false, "A value that sets the write status of a file share.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
	storagegateway_updateSmbfileShareCmd.Flags().Bool("smbaclenabled", false, "Set this value to `true` to enable access control list (ACL) on the SMB file share.")
	storagegateway_updateSmbfileShareCmd.Flags().String("valid-user-list", "", "A list of users or groups in the Active Directory that are allowed to access the file share.")
	storagegateway_updateSmbfileShareCmd.MarkFlagRequired("file-share-arn")
	storagegateway_updateSmbfileShareCmd.Flag("no-access-based-enumeration").Hidden = true
	storagegateway_updateSmbfileShareCmd.Flag("no-guess-mimetype-enabled").Hidden = true
	storagegateway_updateSmbfileShareCmd.Flag("no-kmsencrypted").Hidden = true
	storagegateway_updateSmbfileShareCmd.Flag("no-oplocks-enabled").Hidden = true
	storagegateway_updateSmbfileShareCmd.Flag("no-read-only").Hidden = true
	storagegateway_updateSmbfileShareCmd.Flag("no-requester-pays").Hidden = true
	storagegateway_updateSmbfileShareCmd.Flag("no-smbaclenabled").Hidden = true
	storagegatewayCmd.AddCommand(storagegateway_updateSmbfileShareCmd)
}
