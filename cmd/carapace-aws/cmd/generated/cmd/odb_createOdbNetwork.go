package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_createOdbNetworkCmd = &cobra.Command{
	Use:   "create-odb-network",
	Short: "Creates an ODB network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_createOdbNetworkCmd).Standalone()

	odb_createOdbNetworkCmd.Flags().String("availability-zone", "", "The Amazon Web Services Availability Zone (AZ) where the ODB network is located.")
	odb_createOdbNetworkCmd.Flags().String("availability-zone-id", "", "The AZ ID of the AZ where the ODB network is located.")
	odb_createOdbNetworkCmd.Flags().String("backup-subnet-cidr", "", "The CIDR range of the backup subnet for the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("client-subnet-cidr", "", "The CIDR range of the client subnet for the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	odb_createOdbNetworkCmd.Flags().String("custom-domain-name", "", "The domain name to use for the resources in the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("default-dns-prefix", "", "The DNS prefix to the default DNS domain name.")
	odb_createOdbNetworkCmd.Flags().String("display-name", "", "A user-friendly name for the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("s3-access", "", "Specifies the configuration for Amazon S3 access from the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("s3-policy-document", "", "Specifies the endpoint policy for Amazon S3 access from the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("tags", "", "The list of resource tags to apply to the ODB network.")
	odb_createOdbNetworkCmd.Flags().String("zero-etl-access", "", "Specifies the configuration for Zero-ETL access from the ODB network.")
	odb_createOdbNetworkCmd.MarkFlagRequired("client-subnet-cidr")
	odb_createOdbNetworkCmd.MarkFlagRequired("display-name")
	odbCmd.AddCommand(odb_createOdbNetworkCmd)
}
