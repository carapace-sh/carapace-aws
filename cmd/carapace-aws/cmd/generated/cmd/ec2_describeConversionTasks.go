package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeConversionTasksCmd = &cobra.Command{
	Use:   "describe-conversion-tasks",
	Short: "Describes the specified conversion tasks or all your conversion tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeConversionTasksCmd).Standalone()

	ec2_describeConversionTasksCmd.Flags().String("conversion-task-ids", "", "The conversion task IDs.")
	ec2_describeConversionTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeConversionTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeConversionTasksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeConversionTasksCmd)
}
