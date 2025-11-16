package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createInstanceProfileCmd = &cobra.Command{
	Use:   "create-instance-profile",
	Short: "Creates a profile that can be applied to one or more private fleet device instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_createInstanceProfileCmd).Standalone()

		devicefarm_createInstanceProfileCmd.Flags().String("description", "", "The description of your instance profile.")
		devicefarm_createInstanceProfileCmd.Flags().String("exclude-app-packages-from-cleanup", "", "An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.")
		devicefarm_createInstanceProfileCmd.Flags().String("name", "", "The name of your instance profile.")
		devicefarm_createInstanceProfileCmd.Flags().Bool("no-package-cleanup", false, "When set to `true`, Device Farm removes app packages after a test run.")
		devicefarm_createInstanceProfileCmd.Flags().Bool("no-reboot-after-use", false, "When set to `true`, Device Farm reboots the instance after a test run.")
		devicefarm_createInstanceProfileCmd.Flags().Bool("package-cleanup", false, "When set to `true`, Device Farm removes app packages after a test run.")
		devicefarm_createInstanceProfileCmd.Flags().Bool("reboot-after-use", false, "When set to `true`, Device Farm reboots the instance after a test run.")
		devicefarm_createInstanceProfileCmd.MarkFlagRequired("name")
		devicefarm_createInstanceProfileCmd.Flag("no-package-cleanup").Hidden = true
		devicefarm_createInstanceProfileCmd.Flag("no-reboot-after-use").Hidden = true
	})
	devicefarmCmd.AddCommand(devicefarm_createInstanceProfileCmd)
}
