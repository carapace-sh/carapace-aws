package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createNfsfileShareCmd = &cobra.Command{
	Use:   "create-nfsfile-share",
	Short: "Creates a Network File System (NFS) file share on an existing S3 File Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createNfsfileShareCmd).Standalone()

	storagegateway_createNfsfileShareCmd.Flags().String("audit-destination-arn", "", "The Amazon Resource Name (ARN) of the storage used for audit logs.")
	storagegateway_createNfsfileShareCmd.Flags().String("bucket-region", "", "Specifies the Region of the S3 bucket where the NFS file share stores files.")
	storagegateway_createNfsfileShareCmd.Flags().String("cache-attributes", "", "Specifies refresh cache information for the file share.")
	storagegateway_createNfsfileShareCmd.Flags().String("client-list", "", "The list of clients that are allowed to access the S3 File Gateway.")
	storagegateway_createNfsfileShareCmd.Flags().String("client-token", "", "A unique string value that you supply that is used by S3 File Gateway to ensure idempotent file share creation.")
	storagegateway_createNfsfileShareCmd.Flags().String("default-storage-class", "", "The default storage class for objects put into an Amazon S3 bucket by the S3 File Gateway.")
	storagegateway_createNfsfileShareCmd.Flags().String("encryption-type", "", "A value that specifies the type of server-side encryption that the file share will use for the data that it stores in Amazon S3.")
	storagegateway_createNfsfileShareCmd.Flags().String("file-share-name", "", "The name of the file share.")
	storagegateway_createNfsfileShareCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the S3 File Gateway on which you want to create a file share.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("kmsencrypted", false, "Optional.")
	storagegateway_createNfsfileShareCmd.Flags().String("kmskey", "", "Optional.")
	storagegateway_createNfsfileShareCmd.Flags().String("location-arn", "", "A custom ARN for the backend storage used for storing data for file shares.")
	storagegateway_createNfsfileShareCmd.Flags().String("nfsfile-share-defaults", "", "File share default values.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("no-guess-mimetype-enabled", false, "A value that enables guessing of the MIME type for uploaded objects based on file extensions.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("no-kmsencrypted", false, "Optional.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("no-read-only", false, "A value that sets the write status of a file share.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("no-requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
	storagegateway_createNfsfileShareCmd.Flags().String("notification-policy", "", "The notification policy of the file share.")
	storagegateway_createNfsfileShareCmd.Flags().String("object-acl", "", "A value that sets the access control list (ACL) permission for objects in the S3 bucket that a S3 File Gateway puts objects into.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("read-only", false, "A value that sets the write status of a file share.")
	storagegateway_createNfsfileShareCmd.Flags().Bool("requester-pays", false, "A value that sets who pays the cost of the request and the cost associated with data download from the S3 bucket.")
	storagegateway_createNfsfileShareCmd.Flags().String("role", "", "The ARN of the Identity and Access Management (IAM) role that an S3 File Gateway assumes when it accesses the underlying storage.")
	storagegateway_createNfsfileShareCmd.Flags().String("squash", "", "A value that maps a user to anonymous user.")
	storagegateway_createNfsfileShareCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to the NFS file share.")
	storagegateway_createNfsfileShareCmd.Flags().String("vpcendpoint-dnsname", "", "Specifies the DNS name for the VPC endpoint that the NFS file share uses to connect to Amazon S3.")
	storagegateway_createNfsfileShareCmd.MarkFlagRequired("client-token")
	storagegateway_createNfsfileShareCmd.MarkFlagRequired("gateway-arn")
	storagegateway_createNfsfileShareCmd.MarkFlagRequired("location-arn")
	storagegateway_createNfsfileShareCmd.Flag("no-guess-mimetype-enabled").Hidden = true
	storagegateway_createNfsfileShareCmd.Flag("no-kmsencrypted").Hidden = true
	storagegateway_createNfsfileShareCmd.Flag("no-read-only").Hidden = true
	storagegateway_createNfsfileShareCmd.Flag("no-requester-pays").Hidden = true
	storagegateway_createNfsfileShareCmd.MarkFlagRequired("role")
	storagegatewayCmd.AddCommand(storagegateway_createNfsfileShareCmd)
}
