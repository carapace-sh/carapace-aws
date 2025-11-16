package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_createTrailCmd = &cobra.Command{
	Use:   "create-trail",
	Short: "Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_createTrailCmd).Standalone()

	cloudtrail_createTrailCmd.Flags().String("cloud-watch-logs-log-group-arn", "", "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered.")
	cloudtrail_createTrailCmd.Flags().String("cloud-watch-logs-role-arn", "", "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.")
	cloudtrail_createTrailCmd.Flags().Bool("enable-log-file-validation", false, "Specifies whether log file integrity validation is enabled.")
	cloudtrail_createTrailCmd.Flags().Bool("include-global-service-events", false, "Specifies whether the trail is publishing events from global services such as IAM to the log files.")
	cloudtrail_createTrailCmd.Flags().Bool("is-multi-region-trail", false, "Specifies whether the trail is created in the current Region or in all Regions.")
	cloudtrail_createTrailCmd.Flags().Bool("is-organization-trail", false, "Specifies whether the trail is created for all accounts in an organization in Organizations, or only for the current Amazon Web Services account.")
	cloudtrail_createTrailCmd.Flags().String("kms-key-id", "", "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.")
	cloudtrail_createTrailCmd.Flags().String("name", "", "Specifies the name of the trail.")
	cloudtrail_createTrailCmd.Flags().Bool("no-enable-log-file-validation", false, "Specifies whether log file integrity validation is enabled.")
	cloudtrail_createTrailCmd.Flags().Bool("no-include-global-service-events", false, "Specifies whether the trail is publishing events from global services such as IAM to the log files.")
	cloudtrail_createTrailCmd.Flags().Bool("no-is-multi-region-trail", false, "Specifies whether the trail is created in the current Region or in all Regions.")
	cloudtrail_createTrailCmd.Flags().Bool("no-is-organization-trail", false, "Specifies whether the trail is created for all accounts in an organization in Organizations, or only for the current Amazon Web Services account.")
	cloudtrail_createTrailCmd.Flags().String("s3-bucket-name", "", "Specifies the name of the Amazon S3 bucket designated for publishing log files.")
	cloudtrail_createTrailCmd.Flags().String("s3-key-prefix", "", "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.")
	cloudtrail_createTrailCmd.Flags().String("sns-topic-name", "", "Specifies the name or ARN of the Amazon SNS topic defined for notification of log file delivery.")
	cloudtrail_createTrailCmd.Flags().String("tags-list", "", "")
	cloudtrail_createTrailCmd.MarkFlagRequired("name")
	cloudtrail_createTrailCmd.Flag("no-enable-log-file-validation").Hidden = true
	cloudtrail_createTrailCmd.Flag("no-include-global-service-events").Hidden = true
	cloudtrail_createTrailCmd.Flag("no-is-multi-region-trail").Hidden = true
	cloudtrail_createTrailCmd.Flag("no-is-organization-trail").Hidden = true
	cloudtrail_createTrailCmd.MarkFlagRequired("s3-bucket-name")
	cloudtrailCmd.AddCommand(cloudtrail_createTrailCmd)
}
