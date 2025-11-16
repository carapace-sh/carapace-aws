package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_putDefaultApplicationSettingCmd = &cobra.Command{
	Use:   "put-default-application-setting",
	Short: "Sets the default application to the application with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_putDefaultApplicationSettingCmd).Standalone()

	opensearch_putDefaultApplicationSettingCmd.Flags().String("application-arn", "", "")
	opensearch_putDefaultApplicationSettingCmd.Flags().Bool("no-set-as-default", false, "Set to true to set the specified ARN as the default application.")
	opensearch_putDefaultApplicationSettingCmd.Flags().Bool("set-as-default", false, "Set to true to set the specified ARN as the default application.")
	opensearch_putDefaultApplicationSettingCmd.MarkFlagRequired("application-arn")
	opensearch_putDefaultApplicationSettingCmd.Flag("no-set-as-default").Hidden = true
	opensearch_putDefaultApplicationSettingCmd.MarkFlagRequired("no-set-as-default")
	opensearch_putDefaultApplicationSettingCmd.MarkFlagRequired("set-as-default")
	opensearchCmd.AddCommand(opensearch_putDefaultApplicationSettingCmd)
}
