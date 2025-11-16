package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_updateOdbNetworkCmd = &cobra.Command{
	Use:   "update-odb-network",
	Short: "Updates properties of a specified ODB network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_updateOdbNetworkCmd).Standalone()

	odb_updateOdbNetworkCmd.Flags().String("display-name", "", "The new user-friendly name of the ODB network.")
	odb_updateOdbNetworkCmd.Flags().String("odb-network-id", "", "The unique identifier of the ODB network to update.")
	odb_updateOdbNetworkCmd.Flags().String("peered-cidrs-to-be-added", "", "The list of CIDR ranges from the peered VPC that allow access to the ODB network.")
	odb_updateOdbNetworkCmd.Flags().String("peered-cidrs-to-be-removed", "", "The list of CIDR ranges from the peered VPC to remove from the ODB network.")
	odb_updateOdbNetworkCmd.Flags().String("s3-access", "", "Specifies the updated configuration for Amazon S3 access from the ODB network.")
	odb_updateOdbNetworkCmd.Flags().String("s3-policy-document", "", "Specifies the updated endpoint policy for Amazon S3 access from the ODB network.")
	odb_updateOdbNetworkCmd.Flags().String("zero-etl-access", "", "Specifies the updated configuration for Zero-ETL access from the ODB network.")
	odb_updateOdbNetworkCmd.MarkFlagRequired("odb-network-id")
	odbCmd.AddCommand(odb_updateOdbNetworkCmd)
}
