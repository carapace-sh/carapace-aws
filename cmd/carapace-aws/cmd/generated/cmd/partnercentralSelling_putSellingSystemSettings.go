package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_putSellingSystemSettingsCmd = &cobra.Command{
	Use:   "put-selling-system-settings",
	Short: "Updates the currently set system settings, which include the IAM Role used for resource snapshot jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_putSellingSystemSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_putSellingSystemSettingsCmd).Standalone()

		partnercentralSelling_putSellingSystemSettingsCmd.Flags().String("catalog", "", "Specifies the catalog in which the settings will be updated.")
		partnercentralSelling_putSellingSystemSettingsCmd.Flags().String("resource-snapshot-job-role-identifier", "", "Specifies the ARN of the IAM Role used for resource snapshot job executions.")
		partnercentralSelling_putSellingSystemSettingsCmd.MarkFlagRequired("catalog")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_putSellingSystemSettingsCmd)
}
