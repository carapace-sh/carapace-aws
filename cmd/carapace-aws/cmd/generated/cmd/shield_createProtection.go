package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_createProtectionCmd = &cobra.Command{
	Use:   "create-protection",
	Short: "Enables Shield Advanced for a specific Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_createProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_createProtectionCmd).Standalone()

		shield_createProtectionCmd.Flags().String("name", "", "Friendly name for the `Protection` you are creating.")
		shield_createProtectionCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the resource to be protected.")
		shield_createProtectionCmd.Flags().String("tags", "", "One or more tag key-value pairs for the [Protection]() object that is created.")
		shield_createProtectionCmd.MarkFlagRequired("name")
		shield_createProtectionCmd.MarkFlagRequired("resource-arn")
	})
	shieldCmd.AddCommand(shield_createProtectionCmd)
}
