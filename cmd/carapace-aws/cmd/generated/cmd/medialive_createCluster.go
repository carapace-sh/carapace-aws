package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Create a new Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createClusterCmd).Standalone()

	medialive_createClusterCmd.Flags().String("cluster-type", "", "Specify a type.")
	medialive_createClusterCmd.Flags().String("instance-role-arn", "", "The ARN of the IAM role for the Node in this Cluster.")
	medialive_createClusterCmd.Flags().String("name", "", "Specify a name that is unique in the AWS account.")
	medialive_createClusterCmd.Flags().String("network-settings", "", "Network settings that connect the Nodes in the Cluster to one or more of the Networks that the Cluster is associated with.")
	medialive_createClusterCmd.Flags().String("request-id", "", "The unique ID of the request.")
	medialive_createClusterCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialiveCmd.AddCommand(medialive_createClusterCmd)
}
