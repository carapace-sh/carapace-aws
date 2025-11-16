package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createReplicationConfigCmd = &cobra.Command{
	Use:   "create-replication-config",
	Short: "Creates a configuration that you can later provide to configure and start an DMS Serverless replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createReplicationConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createReplicationConfigCmd).Standalone()

		dms_createReplicationConfigCmd.Flags().String("compute-config", "", "Configuration parameters for provisioning an DMS Serverless replication.")
		dms_createReplicationConfigCmd.Flags().String("replication-config-identifier", "", "A unique identifier that you want to use to create a `ReplicationConfigArn` that is returned as part of the output from this action.")
		dms_createReplicationConfigCmd.Flags().String("replication-settings", "", "Optional JSON settings for DMS Serverless replications that are provisioned using this replication configuration.")
		dms_createReplicationConfigCmd.Flags().String("replication-type", "", "The type of DMS Serverless replication to provision using this replication configuration.")
		dms_createReplicationConfigCmd.Flags().String("resource-identifier", "", "Optional unique value or name that you set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource.")
		dms_createReplicationConfigCmd.Flags().String("source-endpoint-arn", "", "The Amazon Resource Name (ARN) of the source endpoint for this DMS Serverless replication configuration.")
		dms_createReplicationConfigCmd.Flags().String("supplemental-settings", "", "Optional JSON settings for specifying supplemental data.")
		dms_createReplicationConfigCmd.Flags().String("table-mappings", "", "JSON table mappings for DMS Serverless replications that are provisioned using this replication configuration.")
		dms_createReplicationConfigCmd.Flags().String("tags", "", "One or more optional tags associated with resources used by the DMS Serverless replication.")
		dms_createReplicationConfigCmd.Flags().String("target-endpoint-arn", "", "The Amazon Resource Name (ARN) of the target endpoint for this DMS serverless replication configuration.")
		dms_createReplicationConfigCmd.MarkFlagRequired("compute-config")
		dms_createReplicationConfigCmd.MarkFlagRequired("replication-config-identifier")
		dms_createReplicationConfigCmd.MarkFlagRequired("replication-type")
		dms_createReplicationConfigCmd.MarkFlagRequired("source-endpoint-arn")
		dms_createReplicationConfigCmd.MarkFlagRequired("table-mappings")
		dms_createReplicationConfigCmd.MarkFlagRequired("target-endpoint-arn")
	})
	dmsCmd.AddCommand(dms_createReplicationConfigCmd)
}
