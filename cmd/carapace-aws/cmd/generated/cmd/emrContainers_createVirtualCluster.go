package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_createVirtualClusterCmd = &cobra.Command{
	Use:   "create-virtual-cluster",
	Short: "Creates a virtual cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_createVirtualClusterCmd).Standalone()

	emrContainers_createVirtualClusterCmd.Flags().String("client-token", "", "The client token of the virtual cluster.")
	emrContainers_createVirtualClusterCmd.Flags().String("container-provider", "", "The container provider of the virtual cluster.")
	emrContainers_createVirtualClusterCmd.Flags().String("name", "", "The specified name of the virtual cluster.")
	emrContainers_createVirtualClusterCmd.Flags().String("security-configuration-id", "", "The ID of the security configuration.")
	emrContainers_createVirtualClusterCmd.Flags().String("tags", "", "The tags assigned to the virtual cluster.")
	emrContainers_createVirtualClusterCmd.MarkFlagRequired("client-token")
	emrContainers_createVirtualClusterCmd.MarkFlagRequired("container-provider")
	emrContainers_createVirtualClusterCmd.MarkFlagRequired("name")
	emrContainersCmd.AddCommand(emrContainers_createVirtualClusterCmd)
}
