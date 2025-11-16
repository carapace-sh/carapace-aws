package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_putLifecycleConfigurationCmd = &cobra.Command{
	Use:   "put-lifecycle-configuration",
	Short: "Use this action to manage storage for your file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_putLifecycleConfigurationCmd).Standalone()

	efs_putLifecycleConfigurationCmd.Flags().String("file-system-id", "", "The ID of the file system for which you are creating the `LifecycleConfiguration` object (String).")
	efs_putLifecycleConfigurationCmd.Flags().String("lifecycle-policies", "", "An array of `LifecyclePolicy` objects that define the file system's `LifecycleConfiguration` object.")
	efs_putLifecycleConfigurationCmd.MarkFlagRequired("file-system-id")
	efs_putLifecycleConfigurationCmd.MarkFlagRequired("lifecycle-policies")
	efsCmd.AddCommand(efs_putLifecycleConfigurationCmd)
}
