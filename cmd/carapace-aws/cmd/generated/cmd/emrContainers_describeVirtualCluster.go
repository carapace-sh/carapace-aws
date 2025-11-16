package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_describeVirtualClusterCmd = &cobra.Command{
	Use:   "describe-virtual-cluster",
	Short: "Displays detailed information about a specified virtual cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_describeVirtualClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_describeVirtualClusterCmd).Standalone()

		emrContainers_describeVirtualClusterCmd.Flags().String("id", "", "The ID of the virtual cluster that will be described.")
		emrContainers_describeVirtualClusterCmd.MarkFlagRequired("id")
	})
	emrContainersCmd.AddCommand(emrContainers_describeVirtualClusterCmd)
}
