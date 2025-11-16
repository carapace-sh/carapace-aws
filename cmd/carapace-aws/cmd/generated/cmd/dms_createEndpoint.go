package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createEndpointCmd = &cobra.Command{
	Use:   "create-endpoint",
	Short: "Creates an endpoint using the provided settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createEndpointCmd).Standalone()

		dms_createEndpointCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) for the certificate.")
		dms_createEndpointCmd.Flags().String("database-name", "", "The name of the endpoint database.")
		dms_createEndpointCmd.Flags().String("dms-transfer-settings", "", "The settings in JSON format for the DMS transfer type of source endpoint.")
		dms_createEndpointCmd.Flags().String("doc-db-settings", "", "")
		dms_createEndpointCmd.Flags().String("dynamo-db-settings", "", "Settings in JSON format for the target Amazon DynamoDB endpoint.")
		dms_createEndpointCmd.Flags().String("elasticsearch-settings", "", "Settings in JSON format for the target OpenSearch endpoint.")
		dms_createEndpointCmd.Flags().String("endpoint-identifier", "", "The database endpoint identifier.")
		dms_createEndpointCmd.Flags().String("endpoint-type", "", "The type of endpoint.")
		dms_createEndpointCmd.Flags().String("engine-name", "", "The type of engine for the endpoint.")
		dms_createEndpointCmd.Flags().String("external-table-definition", "", "The external table definition.")
		dms_createEndpointCmd.Flags().String("extra-connection-attributes", "", "Additional attributes associated with the connection.")
		dms_createEndpointCmd.Flags().String("gcp-my-sqlsettings", "", "Settings in JSON format for the source GCP MySQL endpoint.")
		dms_createEndpointCmd.Flags().String("ibmdb2-settings", "", "Settings in JSON format for the source IBM Db2 LUW endpoint.")
		dms_createEndpointCmd.Flags().String("kafka-settings", "", "Settings in JSON format for the target Apache Kafka endpoint.")
		dms_createEndpointCmd.Flags().String("kinesis-settings", "", "Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.")
		dms_createEndpointCmd.Flags().String("kms-key-id", "", "An KMS key identifier that is used to encrypt the connection parameters for the endpoint.")
		dms_createEndpointCmd.Flags().String("microsoft-sqlserver-settings", "", "Settings in JSON format for the source and target Microsoft SQL Server endpoint.")
		dms_createEndpointCmd.Flags().String("mongo-db-settings", "", "Settings in JSON format for the source MongoDB endpoint.")
		dms_createEndpointCmd.Flags().String("my-sqlsettings", "", "Settings in JSON format for the source and target MySQL endpoint.")
		dms_createEndpointCmd.Flags().String("neptune-settings", "", "Settings in JSON format for the target Amazon Neptune endpoint.")
		dms_createEndpointCmd.Flags().String("oracle-settings", "", "Settings in JSON format for the source and target Oracle endpoint.")
		dms_createEndpointCmd.Flags().String("password", "", "The password to be used to log in to the endpoint database.")
		dms_createEndpointCmd.Flags().String("port", "", "The port used by the endpoint database.")
		dms_createEndpointCmd.Flags().String("postgre-sqlsettings", "", "Settings in JSON format for the source and target PostgreSQL endpoint.")
		dms_createEndpointCmd.Flags().String("redis-settings", "", "Settings in JSON format for the target Redis endpoint.")
		dms_createEndpointCmd.Flags().String("redshift-settings", "", "")
		dms_createEndpointCmd.Flags().String("resource-identifier", "", "A friendly name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.")
		dms_createEndpointCmd.Flags().String("s3-settings", "", "Settings in JSON format for the target Amazon S3 endpoint.")
		dms_createEndpointCmd.Flags().String("server-name", "", "The name of the server where the endpoint database resides.")
		dms_createEndpointCmd.Flags().String("service-access-role-arn", "", "The Amazon Resource Name (ARN) for the service access role that you want to use to create the endpoint.")
		dms_createEndpointCmd.Flags().String("ssl-mode", "", "The Secure Sockets Layer (SSL) mode to use for the SSL connection.")
		dms_createEndpointCmd.Flags().String("sybase-settings", "", "Settings in JSON format for the source and target SAP ASE endpoint.")
		dms_createEndpointCmd.Flags().String("tags", "", "One or more tags to be assigned to the endpoint.")
		dms_createEndpointCmd.Flags().String("timestream-settings", "", "Settings in JSON format for the target Amazon Timestream endpoint.")
		dms_createEndpointCmd.Flags().String("username", "", "The user name to be used to log in to the endpoint database.")
		dms_createEndpointCmd.MarkFlagRequired("endpoint-identifier")
		dms_createEndpointCmd.MarkFlagRequired("endpoint-type")
		dms_createEndpointCmd.MarkFlagRequired("engine-name")
	})
	dmsCmd.AddCommand(dms_createEndpointCmd)
}
