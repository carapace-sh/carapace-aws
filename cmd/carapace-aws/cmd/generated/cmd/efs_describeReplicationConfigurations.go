package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeReplicationConfigurationsCmd = &cobra.Command{
	Use:   "describe-replication-configurations",
	Short: "Retrieves the replication configuration for a specific file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeReplicationConfigurationsCmd).Standalone()

	efs_describeReplicationConfigurationsCmd.Flags().String("file-system-id", "", "You can retrieve the replication configuration for a specific file system by providing its file system ID.")
	efs_describeReplicationConfigurationsCmd.Flags().String("max-results", "", "(Optional) To limit the number of objects returned in a response, you can specify the `MaxItems` parameter.")
	efs_describeReplicationConfigurationsCmd.Flags().String("next-token", "", "`NextToken` is present if the response is paginated.")
	efsCmd.AddCommand(efs_describeReplicationConfigurationsCmd)
}
