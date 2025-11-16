package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_createDbClusterCmd = &cobra.Command{
	Use:   "create-db-cluster",
	Short: "Creates a new Timestream for InfluxDB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_createDbClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_createDbClusterCmd).Standalone()

		timestreamInfluxdb_createDbClusterCmd.Flags().String("allocated-storage", "", "The amount of storage to allocate for your DB storage type in GiB (gibibytes).")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("bucket", "", "The name of the initial InfluxDB bucket.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("db-instance-type", "", "The Timestream for InfluxDB DB instance type to run InfluxDB on.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("db-parameter-group-identifier", "", "The ID of the DB parameter group to assign to your DB cluster.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("db-storage-type", "", "The Timestream for InfluxDB DB storage type to read and write InfluxDB data.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("deployment-type", "", "Specifies the type of cluster to create.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("failover-mode", "", "Specifies the behavior of failure recovery when the primary node of the cluster fails.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("log-delivery-configuration", "", "Configuration for sending InfluxDB engine logs to a specified S3 bucket.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("name", "", "The name that uniquely identifies the DB cluster when interacting with the Amazon Timestream for InfluxDB API and CLI commands.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("network-type", "", "Specifies whether the network type of the Timestream for InfluxDB cluster is IPv4, which can communicate over IPv4 protocol only, or DUAL, which can communicate over both IPv4 and IPv6 protocols.")
		timestreamInfluxdb_createDbClusterCmd.Flags().Bool("no-publicly-accessible", false, "Configures the Timestream for InfluxDB cluster with a public IP to facilitate access from outside the VPC.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("organization", "", "The name of the initial organization for the initial admin user in InfluxDB.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("password", "", "The password of the initial admin user created in InfluxDB.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("port", "", "The port number on which InfluxDB accepts connections.")
		timestreamInfluxdb_createDbClusterCmd.Flags().Bool("publicly-accessible", false, "Configures the Timestream for InfluxDB cluster with a public IP to facilitate access from outside the VPC.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the DB instance.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("username", "", "The username of the initial admin user created in InfluxDB.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security group IDs to associate with the Timestream for InfluxDB cluster.")
		timestreamInfluxdb_createDbClusterCmd.Flags().String("vpc-subnet-ids", "", "A list of VPC subnet IDs to associate with the DB cluster.")
		timestreamInfluxdb_createDbClusterCmd.MarkFlagRequired("db-instance-type")
		timestreamInfluxdb_createDbClusterCmd.MarkFlagRequired("name")
		timestreamInfluxdb_createDbClusterCmd.Flag("no-publicly-accessible").Hidden = true
		timestreamInfluxdb_createDbClusterCmd.MarkFlagRequired("vpc-security-group-ids")
		timestreamInfluxdb_createDbClusterCmd.MarkFlagRequired("vpc-subnet-ids")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_createDbClusterCmd)
}
