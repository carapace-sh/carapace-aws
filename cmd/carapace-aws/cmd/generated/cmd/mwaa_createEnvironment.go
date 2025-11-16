package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Creates an Amazon Managed Workflows for Apache Airflow (Amazon MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_createEnvironmentCmd).Standalone()

	mwaa_createEnvironmentCmd.Flags().String("airflow-configuration-options", "", "A list of key-value pairs containing the Apache Airflow configuration options you want to attach to your environment.")
	mwaa_createEnvironmentCmd.Flags().String("airflow-version", "", "The Apache Airflow version for your environment.")
	mwaa_createEnvironmentCmd.Flags().String("dag-s3-path", "", "The relative path to the DAGs folder on your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("endpoint-management", "", "Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.")
	mwaa_createEnvironmentCmd.Flags().String("environment-class", "", "The environment class type.")
	mwaa_createEnvironmentCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the execution role for your environment.")
	mwaa_createEnvironmentCmd.Flags().String("kms-key", "", "The Amazon Web Services Key Management Service (KMS) key to encrypt the data in your environment.")
	mwaa_createEnvironmentCmd.Flags().String("logging-configuration", "", "Defines the Apache Airflow logs to send to CloudWatch Logs.")
	mwaa_createEnvironmentCmd.Flags().String("max-webservers", "", "The maximum number of web servers that you want to run in your environment.")
	mwaa_createEnvironmentCmd.Flags().String("max-workers", "", "The maximum number of workers that you want to run in your environment.")
	mwaa_createEnvironmentCmd.Flags().String("min-webservers", "", "The minimum number of web servers that you want to run in your environment.")
	mwaa_createEnvironmentCmd.Flags().String("min-workers", "", "The minimum number of workers that you want to run in your environment.")
	mwaa_createEnvironmentCmd.Flags().String("name", "", "The name of the Amazon MWAA environment.")
	mwaa_createEnvironmentCmd.Flags().String("network-configuration", "", "The VPC networking components used to secure and enable network traffic between the Amazon Web Services resources for your environment.")
	mwaa_createEnvironmentCmd.Flags().String("plugins-s3-object-version", "", "The version of the plugins.zip file on your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("plugins-s3-path", "", "The relative path to the `plugins.zip` file on your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("requirements-s3-object-version", "", "The version of the `requirements.txt` file on your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("requirements-s3-path", "", "The relative path to the `requirements.txt` file on your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("schedulers", "", "The number of Apache Airflow schedulers to run in your environment.")
	mwaa_createEnvironmentCmd.Flags().String("source-bucket-arn", "", "The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored.")
	mwaa_createEnvironmentCmd.Flags().String("startup-script-s3-object-version", "", "The version of the startup shell script in your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("startup-script-s3-path", "", "The relative path to the startup shell script in your Amazon S3 bucket.")
	mwaa_createEnvironmentCmd.Flags().String("tags", "", "The key-value tag pairs you want to associate to your environment.")
	mwaa_createEnvironmentCmd.Flags().String("webserver-access-mode", "", "Defines the access mode for the Apache Airflow *web server*.")
	mwaa_createEnvironmentCmd.Flags().String("weekly-maintenance-window-start", "", "The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM`.")
	mwaa_createEnvironmentCmd.MarkFlagRequired("dag-s3-path")
	mwaa_createEnvironmentCmd.MarkFlagRequired("execution-role-arn")
	mwaa_createEnvironmentCmd.MarkFlagRequired("name")
	mwaa_createEnvironmentCmd.MarkFlagRequired("network-configuration")
	mwaa_createEnvironmentCmd.MarkFlagRequired("source-bucket-arn")
	mwaaCmd.AddCommand(mwaa_createEnvironmentCmd)
}
