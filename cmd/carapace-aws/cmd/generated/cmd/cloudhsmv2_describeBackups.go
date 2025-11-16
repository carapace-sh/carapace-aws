package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_describeBackupsCmd = &cobra.Command{
	Use:   "describe-backups",
	Short: "Gets information about backups of CloudHSM clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_describeBackupsCmd).Standalone()

	cloudhsmv2_describeBackupsCmd.Flags().String("filters", "", "One or more filters to limit the items returned in the response.")
	cloudhsmv2_describeBackupsCmd.Flags().String("max-results", "", "The maximum number of backups to return in the response.")
	cloudhsmv2_describeBackupsCmd.Flags().String("next-token", "", "The `NextToken` value that you received in the previous response.")
	cloudhsmv2_describeBackupsCmd.Flags().Bool("no-shared", false, "Describe backups that are shared with you.")
	cloudhsmv2_describeBackupsCmd.Flags().Bool("no-sort-ascending", false, "Designates whether or not to sort the return backups by ascending chronological order of generation.")
	cloudhsmv2_describeBackupsCmd.Flags().Bool("shared", false, "Describe backups that are shared with you.")
	cloudhsmv2_describeBackupsCmd.Flags().Bool("sort-ascending", false, "Designates whether or not to sort the return backups by ascending chronological order of generation.")
	cloudhsmv2_describeBackupsCmd.Flag("no-shared").Hidden = true
	cloudhsmv2_describeBackupsCmd.Flag("no-sort-ascending").Hidden = true
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_describeBackupsCmd)
}
