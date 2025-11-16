package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_runJobFlowCmd = &cobra.Command{
	Use:   "run-job-flow",
	Short: "RunJobFlow creates and starts running a new cluster (job flow).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_runJobFlowCmd).Standalone()

	emr_runJobFlowCmd.Flags().String("additional-info", "", "A JSON string for selecting additional features.")
	emr_runJobFlowCmd.Flags().String("ami-version", "", "Applies only to Amazon EMR AMI versions 3.x and 2.x. For Amazon EMR releases 4.0 and later, `ReleaseLabel` is used.")
	emr_runJobFlowCmd.Flags().String("applications", "", "Applies to Amazon EMR releases 4.0 and later.")
	emr_runJobFlowCmd.Flags().String("auto-scaling-role", "", "An IAM role for automatic scaling policies.")
	emr_runJobFlowCmd.Flags().String("auto-termination-policy", "", "")
	emr_runJobFlowCmd.Flags().String("bootstrap-actions", "", "A list of bootstrap actions to run before Hadoop starts on the cluster nodes.")
	emr_runJobFlowCmd.Flags().String("configurations", "", "For Amazon EMR releases 4.0 and later.")
	emr_runJobFlowCmd.Flags().String("custom-ami-id", "", "Available only in Amazon EMR releases 5.7.0 and later.")
	emr_runJobFlowCmd.Flags().String("ebs-root-volume-iops", "", "The IOPS, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.")
	emr_runJobFlowCmd.Flags().String("ebs-root-volume-size", "", "The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.")
	emr_runJobFlowCmd.Flags().String("ebs-root-volume-throughput", "", "The throughput, in MiB/s, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.")
	emr_runJobFlowCmd.Flags().String("extended-support", "", "Reserved.")
	emr_runJobFlowCmd.Flags().String("instances", "", "A specification of the number and type of Amazon EC2 instances.")
	emr_runJobFlowCmd.Flags().String("job-flow-role", "", "Also called instance profile and Amazon EC2 role.")
	emr_runJobFlowCmd.Flags().String("kerberos-attributes", "", "Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.")
	emr_runJobFlowCmd.Flags().String("log-encryption-kms-key-id", "", "The KMS key used for encrypting log files.")
	emr_runJobFlowCmd.Flags().String("log-uri", "", "The location in Amazon S3 to write the log files of the job flow.")
	emr_runJobFlowCmd.Flags().String("managed-scaling-policy", "", "The specified managed scaling policy for an Amazon EMR cluster.")
	emr_runJobFlowCmd.Flags().String("name", "", "The name of the job flow.")
	emr_runJobFlowCmd.Flags().String("new-supported-products", "", "For Amazon EMR releases 3.x and 2.x. For Amazon EMR releases 4.x and later, use Applications.")
	emr_runJobFlowCmd.Flags().Bool("no-visible-to-all-users", false, "The VisibleToAllUsers parameter is no longer supported.")
	emr_runJobFlowCmd.Flags().String("osrelease-label", "", "Specifies a particular Amazon Linux release for all nodes in a cluster launch RunJobFlow request.")
	emr_runJobFlowCmd.Flags().String("placement-group-configs", "", "The specified placement group configuration for an Amazon EMR cluster.")
	emr_runJobFlowCmd.Flags().String("release-label", "", "The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.")
	emr_runJobFlowCmd.Flags().String("repo-upgrade-on-boot", "", "Applies only when `CustomAmiID` is used.")
	emr_runJobFlowCmd.Flags().String("scale-down-behavior", "", "Specifies the way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.")
	emr_runJobFlowCmd.Flags().String("security-configuration", "", "The name of a security configuration to apply to the cluster.")
	emr_runJobFlowCmd.Flags().String("service-role", "", "The IAM role that Amazon EMR assumes in order to access Amazon Web Services resources on your behalf.")
	emr_runJobFlowCmd.Flags().String("step-concurrency-level", "", "Specifies the number of steps that can be executed concurrently.")
	emr_runJobFlowCmd.Flags().String("steps", "", "A list of steps to run.")
	emr_runJobFlowCmd.Flags().String("supported-products", "", "For Amazon EMR releases 3.x and 2.x. For Amazon EMR releases 4.x and later, use Applications.")
	emr_runJobFlowCmd.Flags().String("tags", "", "A list of tags to associate with a cluster and propagate to Amazon EC2 instances.")
	emr_runJobFlowCmd.Flags().Bool("visible-to-all-users", false, "The VisibleToAllUsers parameter is no longer supported.")
	emr_runJobFlowCmd.MarkFlagRequired("instances")
	emr_runJobFlowCmd.MarkFlagRequired("name")
	emr_runJobFlowCmd.Flag("no-visible-to-all-users").Hidden = true
	emrCmd.AddCommand(emr_runJobFlowCmd)
}
