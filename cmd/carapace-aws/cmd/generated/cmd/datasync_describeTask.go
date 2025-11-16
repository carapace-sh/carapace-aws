package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeTaskCmd = &cobra.Command{
	Use:   "describe-task",
	Short: "Provides information about a *task*, which defines where and how DataSync transfers your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeTaskCmd).Standalone()

	datasync_describeTaskCmd.Flags().String("task-arn", "", "Specifies the Amazon Resource Name (ARN) of the transfer task that you want information about.")
	datasync_describeTaskCmd.MarkFlagRequired("task-arn")
	datasyncCmd.AddCommand(datasync_describeTaskCmd)
}
