package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceCpuOptionsCmd = &cobra.Command{
	Use:   "modify-instance-cpu-options",
	Short: "By default, all vCPUs for the instance type are active when you launch an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceCpuOptionsCmd).Standalone()

	ec2_modifyInstanceCpuOptionsCmd.Flags().String("core-count", "", "The number of CPU cores to activate for the specified instance.")
	ec2_modifyInstanceCpuOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceCpuOptionsCmd.Flags().String("instance-id", "", "The ID of the instance to update.")
	ec2_modifyInstanceCpuOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceCpuOptionsCmd.Flags().String("threads-per-core", "", "The number of threads to run for each CPU core.")
	ec2_modifyInstanceCpuOptionsCmd.MarkFlagRequired("core-count")
	ec2_modifyInstanceCpuOptionsCmd.MarkFlagRequired("instance-id")
	ec2_modifyInstanceCpuOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyInstanceCpuOptionsCmd.MarkFlagRequired("threads-per-core")
	ec2Cmd.AddCommand(ec2_modifyInstanceCpuOptionsCmd)
}
