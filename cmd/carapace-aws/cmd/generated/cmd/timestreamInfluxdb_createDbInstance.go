package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_createDbInstanceCmd = &cobra.Command{
	Use:   "create-db-instance",
	Short: "Creates a new Timestream for InfluxDB DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_createDbInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_createDbInstanceCmd).Standalone()

		timestreamInfluxdb_createDbInstanceCmd.Flags().String("allocated-storage", "", "The amount of storage to allocate for your DB storage type in GiB (gibibytes).")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("bucket", "", "The name of the initial InfluxDB bucket.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("db-instance-type", "", "The Timestream for InfluxDB DB instance type to run InfluxDB on.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("db-parameter-group-identifier", "", "The id of the DB parameter group to assign to your DB instance.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("db-storage-type", "", "The Timestream for InfluxDB DB storage type to read and write InfluxDB data.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("deployment-type", "", "Specifies whether the DB instance will be deployed as a standalone instance or with a Multi-AZ standby for high availability.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("log-delivery-configuration", "", "Configuration for sending InfluxDB engine logs to a specified S3 bucket.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("name", "", "The name that uniquely identifies the DB instance when interacting with the Amazon Timestream for InfluxDB API and CLI commands.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("network-type", "", "Specifies whether the networkType of the Timestream for InfluxDB instance is IPV4, which can communicate over IPv4 protocol only, or DUAL, which can communicate over both IPv4 and IPv6 protocols.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().Bool("no-publicly-accessible", false, "Configures the DB instance with a public IP to facilitate access.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("organization", "", "The name of the initial organization for the initial admin user in InfluxDB.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("password", "", "The password of the initial admin user created in InfluxDB v2.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("port", "", "The port number on which InfluxDB accepts connections.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().Bool("publicly-accessible", false, "Configures the DB instance with a public IP to facilitate access.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the DB instance.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("username", "", "The username of the initial admin user created in InfluxDB.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security group IDs to associate with the DB instance.")
		timestreamInfluxdb_createDbInstanceCmd.Flags().String("vpc-subnet-ids", "", "A list of VPC subnet IDs to associate with the DB instance.")
		timestreamInfluxdb_createDbInstanceCmd.MarkFlagRequired("allocated-storage")
		timestreamInfluxdb_createDbInstanceCmd.MarkFlagRequired("db-instance-type")
		timestreamInfluxdb_createDbInstanceCmd.MarkFlagRequired("name")
		timestreamInfluxdb_createDbInstanceCmd.Flag("no-publicly-accessible").Hidden = true
		timestreamInfluxdb_createDbInstanceCmd.MarkFlagRequired("password")
		timestreamInfluxdb_createDbInstanceCmd.MarkFlagRequired("vpc-security-group-ids")
		timestreamInfluxdb_createDbInstanceCmd.MarkFlagRequired("vpc-subnet-ids")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_createDbInstanceCmd)
}
