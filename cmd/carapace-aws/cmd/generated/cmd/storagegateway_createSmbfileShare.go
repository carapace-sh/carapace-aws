package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createSmbfileShareCmd = &cobra.Command{
	Use:   "create-smbfile-share",
	Short: "Creates a Server Message Block (SMB) file share on an existing S3 File Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createSmbfileShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_createSmbfileShareCmd).Standalone()

		storagegateway_createSmbfileShareCmd.Flags().Bool("access-based-enumeration", false, "The files and folders on this share will only be visible to users with read access.")
		storagegateway_createSmbfileShareCmd.Flags().String("admin-user-list", "", "A list of users or groups in the Active Directory that will be granted administrator privileges on the file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("audit-destination-arn", "", "The Amazon Resource Name (ARN) of the storage used for audit logs.")
		storagegateway_createSmbfileShareCmd.Flags().String("authentication", "", "The authentication method that users use to access the file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("bucket-region", "", "Specifies the Region of the S3 bucket where the SMB file share stores files.")
		storagegateway_createSmbfileShareCmd.Flags().String("cache-attributes", "", "Specifies refresh cache information for the file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("case-sensitivity", "", "The case of an object name in an Amazon S3 bucket.")
		storagegateway_createSmbfileShareCmd.Flags().String("client-token", "", "A unique string value that you supply that is used by S3 File Gateway to ensure idempotent file share creation.")
		storagegateway_createSmbfileShareCmd.Flags().String("default-storage-class", "", "The default storage class for objects put into an Amazon S3 bucket by the S3 File Gateway.")
		storagegateway_createSmbfileShareCmd.Flags().String("encryption-type", "", "A value that specifies the type of server-side encryption that the file share will use for the data that it stores in Amazon S3.")
		storagegateway_createSmbfileShareCmd.Flags().String("file-share-name", "", "The name of the file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("gateway-arn", "", "The ARN of the S3 File Gateway on which you want to create a file share.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
		storagegateway_createSmbfileShareCmd.Flags().String("invalid-user-list", "", "A list of users or groups in the Active Directory that are not allowed to access the file share.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("kmsencrypted", false, "Optional.")
		storagegateway_createSmbfileShareCmd.Flags().String("kmskey", "", "Optional.")
		storagegateway_createSmbfileShareCmd.Flags().String("location-arn", "", "A custom ARN for the backend storage used for storing data for file shares.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-access-based-enumeration", false, "The files and folders on this share will only be visible to users with read access.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-kmsencrypted", false, "Optional.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-oplocks-enabled", false, "Specifies whether opportunistic locking is enabled for the SMB file share.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-read-only", false, "A value that sets the write status of a file share.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("no-smbaclenabled", false, "Set this value to `true` to enable access control list (ACL) on the SMB file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("notification-policy", "", "The notification policy of the file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("object-acl", "", "A value that sets the access control list (ACL) permission for objects in the S3 bucket that a S3 File Gateway puts objects into.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("oplocks-enabled", false, "Specifies whether opportunistic locking is enabled for the SMB file share.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("read-only", false, "A value that sets the write status of a file share.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
		storagegateway_createSmbfileShareCmd.Flags().String("role", "", "The ARN of the Identity and Access Management (IAM) role that an S3 File Gateway assumes when it accesses the underlying storage.")
		storagegateway_createSmbfileShareCmd.Flags().Bool("smbaclenabled", false, "Set this value to `true` to enable access control list (ACL) on the SMB file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to the NFS file share.")
		storagegateway_createSmbfileShareCmd.Flags().String("valid-user-list", "", "A list of users or groups in the Active Directory that are allowed to access the file [share.")
		storagegateway_createSmbfileShareCmd.Flags().String("vpcendpoint-dnsname", "", "Specifies the DNS name for the VPC endpoint that the SMB file share uses to connect to Amazon S3.")
		storagegateway_createSmbfileShareCmd.MarkFlagRequired("client-token")
		storagegateway_createSmbfileShareCmd.MarkFlagRequired("gateway-arn")
		storagegateway_createSmbfileShareCmd.MarkFlagRequired("location-arn")
		storagegateway_createSmbfileShareCmd.Flag("no-access-based-enumeration").Hidden = true
		storagegateway_createSmbfileShareCmd.Flag("no-guess-mimetype-enabled").Hidden = true
		storagegateway_createSmbfileShareCmd.Flag("no-kmsencrypted").Hidden = true
		storagegateway_createSmbfileShareCmd.Flag("no-oplocks-enabled").Hidden = true
		storagegateway_createSmbfileShareCmd.Flag("no-read-only").Hidden = true
		storagegateway_createSmbfileShareCmd.Flag("no-requester-pays").Hidden = true
		storagegateway_createSmbfileShareCmd.Flag("no-smbaclenabled").Hidden = true
		storagegateway_createSmbfileShareCmd.MarkFlagRequired("role")
	})
	storagegatewayCmd.AddCommand(storagegateway_createSmbfileShareCmd)
}
