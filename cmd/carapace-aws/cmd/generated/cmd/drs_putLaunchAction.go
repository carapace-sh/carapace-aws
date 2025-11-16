package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_putLaunchActionCmd = &cobra.Command{
	Use:   "put-launch-action",
	Short: "Puts a resource launch action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_putLaunchActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_putLaunchActionCmd).Standalone()

		drs_putLaunchActionCmd.Flags().String("action-code", "", "Launch action code.")
		drs_putLaunchActionCmd.Flags().String("action-id", "", "")
		drs_putLaunchActionCmd.Flags().String("action-version", "", "")
		drs_putLaunchActionCmd.Flags().Bool("active", false, "Whether the launch action is active.")
		drs_putLaunchActionCmd.Flags().String("category", "", "")
		drs_putLaunchActionCmd.Flags().String("description", "", "")
		drs_putLaunchActionCmd.Flags().String("name", "", "")
		drs_putLaunchActionCmd.Flags().Bool("no-active", false, "Whether the launch action is active.")
		drs_putLaunchActionCmd.Flags().Bool("no-optional", false, "Whether the launch will not be marked as failed if this action fails.")
		drs_putLaunchActionCmd.Flags().Bool("optional", false, "Whether the launch will not be marked as failed if this action fails.")
		drs_putLaunchActionCmd.Flags().String("order", "", "")
		drs_putLaunchActionCmd.Flags().String("parameters", "", "")
		drs_putLaunchActionCmd.Flags().String("resource-id", "", "")
		drs_putLaunchActionCmd.MarkFlagRequired("action-code")
		drs_putLaunchActionCmd.MarkFlagRequired("action-id")
		drs_putLaunchActionCmd.MarkFlagRequired("action-version")
		drs_putLaunchActionCmd.MarkFlagRequired("active")
		drs_putLaunchActionCmd.MarkFlagRequired("category")
		drs_putLaunchActionCmd.MarkFlagRequired("description")
		drs_putLaunchActionCmd.MarkFlagRequired("name")
		drs_putLaunchActionCmd.Flag("no-active").Hidden = true
		drs_putLaunchActionCmd.MarkFlagRequired("no-active")
		drs_putLaunchActionCmd.Flag("no-optional").Hidden = true
		drs_putLaunchActionCmd.MarkFlagRequired("no-optional")
		drs_putLaunchActionCmd.MarkFlagRequired("optional")
		drs_putLaunchActionCmd.MarkFlagRequired("order")
		drs_putLaunchActionCmd.MarkFlagRequired("resource-id")
	})
	drsCmd.AddCommand(drs_putLaunchActionCmd)
}
