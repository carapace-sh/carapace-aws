package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_updateTrailCmd = &cobra.Command{
	Use:   "update-trail",
	Short: "Updates trail settings that control what events you are logging, and how to handle log files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_updateTrailCmd).Standalone()

	cloudtrail_updateTrailCmd.Flags().String("cloud-watch-logs-log-group-arn", "", "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs are delivered.")
	cloudtrail_updateTrailCmd.Flags().String("cloud-watch-logs-role-arn", "", "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.")
	cloudtrail_updateTrailCmd.Flags().Bool("enable-log-file-validation", false, "Specifies whether log file validation is enabled.")
	cloudtrail_updateTrailCmd.Flags().Bool("include-global-service-events", false, "Specifies whether the trail is publishing events from global services such as IAM to the log files.")
	cloudtrail_updateTrailCmd.Flags().Bool("is-multi-region-trail", false, "Specifies whether the trail applies only to the current Region or to all Regions.")
	cloudtrail_updateTrailCmd.Flags().Bool("is-organization-trail", false, "Specifies whether the trail is applied to all accounts in an organization in Organizations, or only for the current Amazon Web Services account.")
	cloudtrail_updateTrailCmd.Flags().String("kms-key-id", "", "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.")
	cloudtrail_updateTrailCmd.Flags().String("name", "", "Specifies the name of the trail or trail ARN.")
	cloudtrail_updateTrailCmd.Flags().Bool("no-enable-log-file-validation", false, "Specifies whether log file validation is enabled.")
	cloudtrail_updateTrailCmd.Flags().Bool("no-include-global-service-events", false, "Specifies whether the trail is publishing events from global services such as IAM to the log files.")
	cloudtrail_updateTrailCmd.Flags().Bool("no-is-multi-region-trail", false, "Specifies whether the trail applies only to the current Region or to all Regions.")
	cloudtrail_updateTrailCmd.Flags().Bool("no-is-organization-trail", false, "Specifies whether the trail is applied to all accounts in an organization in Organizations, or only for the current Amazon Web Services account.")
	cloudtrail_updateTrailCmd.Flags().String("s3-bucket-name", "", "Specifies the name of the Amazon S3 bucket designated for publishing log files.")
	cloudtrail_updateTrailCmd.Flags().String("s3-key-prefix", "", "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.")
	cloudtrail_updateTrailCmd.Flags().String("sns-topic-name", "", "Specifies the name or ARN of the Amazon SNS topic defined for notification of log file delivery.")
	cloudtrail_updateTrailCmd.MarkFlagRequired("name")
	cloudtrail_updateTrailCmd.Flag("no-enable-log-file-validation").Hidden = true
	cloudtrail_updateTrailCmd.Flag("no-include-global-service-events").Hidden = true
	cloudtrail_updateTrailCmd.Flag("no-is-multi-region-trail").Hidden = true
	cloudtrail_updateTrailCmd.Flag("no-is-organization-trail").Hidden = true
	cloudtrailCmd.AddCommand(cloudtrail_updateTrailCmd)
}
