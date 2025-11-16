package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a new CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_createClusterCmd).Standalone()

	cloudhsmv2_createClusterCmd.Flags().String("backup-retention-policy", "", "A policy that defines how the service retains backups.")
	cloudhsmv2_createClusterCmd.Flags().String("hsm-type", "", "The type of HSM to use in the cluster.")
	cloudhsmv2_createClusterCmd.Flags().String("mode", "", "The mode to use in the cluster.")
	cloudhsmv2_createClusterCmd.Flags().String("network-type", "", "The NetworkType to create a cluster with.")
	cloudhsmv2_createClusterCmd.Flags().String("source-backup-id", "", "The identifier (ID) or the Amazon Resource Name (ARN) of the cluster backup to restore.")
	cloudhsmv2_createClusterCmd.Flags().String("subnet-ids", "", "The identifiers (IDs) of the subnets where you are creating the cluster.")
	cloudhsmv2_createClusterCmd.Flags().String("tag-list", "", "Tags to apply to the CloudHSM cluster during creation.")
	cloudhsmv2_createClusterCmd.MarkFlagRequired("hsm-type")
	cloudhsmv2_createClusterCmd.MarkFlagRequired("subnet-ids")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_createClusterCmd)
}
