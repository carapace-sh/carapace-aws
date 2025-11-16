package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_updateFlywheelCmd = &cobra.Command{
	Use:   "update-flywheel",
	Short: "Update the configuration information for an existing flywheel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_updateFlywheelCmd).Standalone()

	comprehend_updateFlywheelCmd.Flags().String("active-model-arn", "", "The Amazon Resource Number (ARN) of the active model version.")
	comprehend_updateFlywheelCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend permission to access the flywheel data.")
	comprehend_updateFlywheelCmd.Flags().String("data-security-config", "", "Flywheel data security configuration.")
	comprehend_updateFlywheelCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel to update.")
	comprehend_updateFlywheelCmd.MarkFlagRequired("flywheel-arn")
	comprehendCmd.AddCommand(comprehend_updateFlywheelCmd)
}
