package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyEndpointCmd = &cobra.Command{
	Use:   "modify-endpoint",
	Short: "Modifies the specified endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyEndpointCmd).Standalone()

	dms_modifyEndpointCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the certificate used for SSL connection.")
	dms_modifyEndpointCmd.Flags().String("database-name", "", "The name of the endpoint database.")
	dms_modifyEndpointCmd.Flags().String("dms-transfer-settings", "", "The settings in JSON format for the DMS transfer type of source endpoint.")
	dms_modifyEndpointCmd.Flags().String("doc-db-settings", "", "Settings in JSON format for the source DocumentDB endpoint.")
	dms_modifyEndpointCmd.Flags().String("dynamo-db-settings", "", "Settings in JSON format for the target Amazon DynamoDB endpoint.")
	dms_modifyEndpointCmd.Flags().String("elasticsearch-settings", "", "Settings in JSON format for the target OpenSearch endpoint.")
	dms_modifyEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
	dms_modifyEndpointCmd.Flags().String("endpoint-identifier", "", "The database endpoint identifier.")
	dms_modifyEndpointCmd.Flags().String("endpoint-type", "", "The type of endpoint.")
	dms_modifyEndpointCmd.Flags().String("engine-name", "", "The database engine name.")
	dms_modifyEndpointCmd.Flags().String("exact-settings", "", "If this attribute is Y, the current call to `ModifyEndpoint` replaces all existing endpoint settings with the exact settings that you specify in this call.")
	dms_modifyEndpointCmd.Flags().String("external-table-definition", "", "The external table definition.")
	dms_modifyEndpointCmd.Flags().String("extra-connection-attributes", "", "Additional attributes associated with the connection.")
	dms_modifyEndpointCmd.Flags().String("gcp-my-sqlsettings", "", "Settings in JSON format for the source GCP MySQL endpoint.")
	dms_modifyEndpointCmd.Flags().String("ibmdb2-settings", "", "Settings in JSON format for the source IBM Db2 LUW endpoint.")
	dms_modifyEndpointCmd.Flags().String("kafka-settings", "", "Settings in JSON format for the target Apache Kafka endpoint.")
	dms_modifyEndpointCmd.Flags().String("kinesis-settings", "", "Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.")
	dms_modifyEndpointCmd.Flags().String("microsoft-sqlserver-settings", "", "Settings in JSON format for the source and target Microsoft SQL Server endpoint.")
	dms_modifyEndpointCmd.Flags().String("mongo-db-settings", "", "Settings in JSON format for the source MongoDB endpoint.")
	dms_modifyEndpointCmd.Flags().String("my-sqlsettings", "", "Settings in JSON format for the source and target MySQL endpoint.")
	dms_modifyEndpointCmd.Flags().String("neptune-settings", "", "Settings in JSON format for the target Amazon Neptune endpoint.")
	dms_modifyEndpointCmd.Flags().String("oracle-settings", "", "Settings in JSON format for the source and target Oracle endpoint.")
	dms_modifyEndpointCmd.Flags().String("password", "", "The password to be used to login to the endpoint database.")
	dms_modifyEndpointCmd.Flags().String("port", "", "The port used by the endpoint database.")
	dms_modifyEndpointCmd.Flags().String("postgre-sqlsettings", "", "Settings in JSON format for the source and target PostgreSQL endpoint.")
	dms_modifyEndpointCmd.Flags().String("redis-settings", "", "Settings in JSON format for the Redis target endpoint.")
	dms_modifyEndpointCmd.Flags().String("redshift-settings", "", "")
	dms_modifyEndpointCmd.Flags().String("s3-settings", "", "Settings in JSON format for the target Amazon S3 endpoint.")
	dms_modifyEndpointCmd.Flags().String("server-name", "", "The name of the server where the endpoint database resides.")
	dms_modifyEndpointCmd.Flags().String("service-access-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role you want to use to modify the endpoint.")
	dms_modifyEndpointCmd.Flags().String("ssl-mode", "", "The SSL mode used to connect to the endpoint.")
	dms_modifyEndpointCmd.Flags().String("sybase-settings", "", "Settings in JSON format for the source and target SAP ASE endpoint.")
	dms_modifyEndpointCmd.Flags().String("timestream-settings", "", "Settings in JSON format for the target Amazon Timestream endpoint.")
	dms_modifyEndpointCmd.Flags().String("username", "", "The user name to be used to login to the endpoint database.")
	dms_modifyEndpointCmd.MarkFlagRequired("endpoint-arn")
	dmsCmd.AddCommand(dms_modifyEndpointCmd)
}
