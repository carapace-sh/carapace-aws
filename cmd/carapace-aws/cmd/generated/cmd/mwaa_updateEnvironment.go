package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Updates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_updateEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaa_updateEnvironmentCmd).Standalone()

		mwaa_updateEnvironmentCmd.Flags().String("airflow-configuration-options", "", "A list of key-value pairs containing the Apache Airflow configuration options you want to attach to your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("airflow-version", "", "The Apache Airflow version for your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("dag-s3-path", "", "The relative path to the DAGs folder on your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("environment-class", "", "The environment class type.")
		mwaa_updateEnvironmentCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to access Amazon Web Services resources in your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("logging-configuration", "", "The Apache Airflow log types to send to CloudWatch Logs.")
		mwaa_updateEnvironmentCmd.Flags().String("max-webservers", "", "The maximum number of web servers that you want to run in your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("max-workers", "", "The maximum number of workers that you want to run in your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("min-webservers", "", "The minimum number of web servers that you want to run in your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("min-workers", "", "The minimum number of workers that you want to run in your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("name", "", "The name of your Amazon MWAA environment.")
		mwaa_updateEnvironmentCmd.Flags().String("network-configuration", "", "The VPC networking components used to secure and enable network traffic between the Amazon Web Services resources for your environment.")
		mwaa_updateEnvironmentCmd.Flags().String("plugins-s3-object-version", "", "The version of the plugins.zip file on your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("plugins-s3-path", "", "The relative path to the `plugins.zip` file on your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("requirements-s3-object-version", "", "The version of the requirements.txt file on your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("requirements-s3-path", "", "The relative path to the `requirements.txt` file on your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("schedulers", "", "The number of Apache Airflow schedulers to run in your Amazon MWAA environment.")
		mwaa_updateEnvironmentCmd.Flags().String("source-bucket-arn", "", "The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored.")
		mwaa_updateEnvironmentCmd.Flags().String("startup-script-s3-object-version", "", "The version of the startup shell script in your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("startup-script-s3-path", "", "The relative path to the startup shell script in your Amazon S3 bucket.")
		mwaa_updateEnvironmentCmd.Flags().String("webserver-access-mode", "", "The Apache Airflow *Web server* access mode.")
		mwaa_updateEnvironmentCmd.Flags().String("weekly-maintenance-window-start", "", "The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM`.")
		mwaa_updateEnvironmentCmd.Flags().String("worker-replacement-strategy", "", "The worker replacement strategy to use when updating the environment.")
		mwaa_updateEnvironmentCmd.MarkFlagRequired("name")
	})
	mwaaCmd.AddCommand(mwaa_updateEnvironmentCmd)
}
