package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateInstanceProfileCmd = &cobra.Command{
	Use:   "update-instance-profile",
	Short: "Updates information about an existing private device instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateInstanceProfileCmd).Standalone()

	devicefarm_updateInstanceProfileCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the instance profile.")
	devicefarm_updateInstanceProfileCmd.Flags().String("description", "", "The updated description for your instance profile.")
	devicefarm_updateInstanceProfileCmd.Flags().String("exclude-app-packages-from-cleanup", "", "An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run is over.")
	devicefarm_updateInstanceProfileCmd.Flags().String("name", "", "The updated name for your instance profile.")
	devicefarm_updateInstanceProfileCmd.Flags().Bool("no-package-cleanup", false, "The updated choice for whether you want to specify package cleanup.")
	devicefarm_updateInstanceProfileCmd.Flags().Bool("no-reboot-after-use", false, "The updated choice for whether you want to reboot the device after use.")
	devicefarm_updateInstanceProfileCmd.Flags().Bool("package-cleanup", false, "The updated choice for whether you want to specify package cleanup.")
	devicefarm_updateInstanceProfileCmd.Flags().Bool("reboot-after-use", false, "The updated choice for whether you want to reboot the device after use.")
	devicefarm_updateInstanceProfileCmd.MarkFlagRequired("arn")
	devicefarm_updateInstanceProfileCmd.Flag("no-package-cleanup").Hidden = true
	devicefarm_updateInstanceProfileCmd.Flag("no-reboot-after-use").Hidden = true
	devicefarmCmd.AddCommand(devicefarm_updateInstanceProfileCmd)
}
