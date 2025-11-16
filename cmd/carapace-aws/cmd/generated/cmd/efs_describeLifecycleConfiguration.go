package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeLifecycleConfigurationCmd = &cobra.Command{
	Use:   "describe-lifecycle-configuration",
	Short: "Returns the current `LifecycleConfiguration` object for the specified EFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeLifecycleConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_describeLifecycleConfigurationCmd).Standalone()

		efs_describeLifecycleConfigurationCmd.Flags().String("file-system-id", "", "The ID of the file system whose `LifecycleConfiguration` object you want to retrieve (String).")
		efs_describeLifecycleConfigurationCmd.MarkFlagRequired("file-system-id")
	})
	efsCmd.AddCommand(efs_describeLifecycleConfigurationCmd)
}
