package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_listVirtualClustersCmd = &cobra.Command{
	Use:   "list-virtual-clusters",
	Short: "Lists information about the specified virtual cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_listVirtualClustersCmd).Standalone()

	emrContainers_listVirtualClustersCmd.Flags().String("container-provider-id", "", "The container provider ID of the virtual cluster.")
	emrContainers_listVirtualClustersCmd.Flags().String("container-provider-type", "", "The container provider type of the virtual cluster.")
	emrContainers_listVirtualClustersCmd.Flags().String("created-after", "", "The date and time after which the virtual clusters are created.")
	emrContainers_listVirtualClustersCmd.Flags().String("created-before", "", "The date and time before which the virtual clusters are created.")
	emrContainers_listVirtualClustersCmd.Flags().Bool("eks-access-entry-integrated", false, "Optional Boolean that specifies whether the operation should return the virtual clusters that have the access entry integration enabled or disabled.")
	emrContainers_listVirtualClustersCmd.Flags().String("max-results", "", "The maximum number of virtual clusters that can be listed.")
	emrContainers_listVirtualClustersCmd.Flags().String("next-token", "", "The token for the next set of virtual clusters to return.")
	emrContainers_listVirtualClustersCmd.Flags().Bool("no-eks-access-entry-integrated", false, "Optional Boolean that specifies whether the operation should return the virtual clusters that have the access entry integration enabled or disabled.")
	emrContainers_listVirtualClustersCmd.Flags().String("states", "", "The states of the requested virtual clusters.")
	emrContainers_listVirtualClustersCmd.Flag("no-eks-access-entry-integrated").Hidden = true
	emrContainersCmd.AddCommand(emrContainers_listVirtualClustersCmd)
}
