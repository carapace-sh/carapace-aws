package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getParameterHistoryCmd = &cobra.Command{
	Use:   "get-parameter-history",
	Short: "Retrieves the history of all changes to a parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getParameterHistoryCmd).Standalone()

	ssm_getParameterHistoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_getParameterHistoryCmd.Flags().String("name", "", "The name or Amazon Resource Name (ARN) of the parameter for which you want to review history.")
	ssm_getParameterHistoryCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_getParameterHistoryCmd.Flags().Bool("no-with-decryption", false, "Return decrypted values for secure string parameters.")
	ssm_getParameterHistoryCmd.Flags().Bool("with-decryption", false, "Return decrypted values for secure string parameters.")
	ssm_getParameterHistoryCmd.MarkFlagRequired("name")
	ssm_getParameterHistoryCmd.Flag("no-with-decryption").Hidden = true
	ssmCmd.AddCommand(ssm_getParameterHistoryCmd)
}
