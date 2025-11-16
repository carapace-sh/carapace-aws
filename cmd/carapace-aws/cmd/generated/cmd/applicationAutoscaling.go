package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscalingCmd = &cobra.Command{
	Use:   "application-autoscaling",
	Short: "With Application Auto Scaling, you can configure automatic scaling for the following resources:\n\n- Amazon AppStream 2.0 fleets\n- Amazon Aurora Replicas\n- Amazon Comprehend document classification and entity recognizer endpoints\n- Amazon DynamoDB tables and global secondary indexes throughput capacity\n- Amazon ECS services\n- Amazon ElastiCache replication groups (Redis OSS and Valkey) and Memcached clusters\n- Amazon EMR clusters\n- Amazon Keyspaces (for Apache Cassandra) tables\n- Lambda function provisioned concurrency\n- Amazon Managed Streaming for Apache Kafka broker storage\n- Amazon Neptune clusters\n- Amazon SageMaker endpoint variants\n- Amazon SageMaker inference components\n- Amazon SageMaker serverless endpoint provisioned concurrency\n- Spot Fleets (Amazon EC2)\n- Pool of WorkSpaces\n- Custom resources provided by your own applications or services\n\nTo learn more about Application Auto Scaling, see the [Application Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/application/userguide/what-is-application-auto-scaling.html).\n\n**API Summary**\n\nThe Application Auto Scaling service API includes three key sets of actions:\n\n- Register and manage scalable targets - Register Amazon Web Services or custom resources as scalable targets (a resource that Application Auto Scaling can scale), set minimum and maximum capacity limits, and retrieve information on existing scalable targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscalingCmd).Standalone()

	rootCmd.AddCommand(applicationAutoscalingCmd)
}
