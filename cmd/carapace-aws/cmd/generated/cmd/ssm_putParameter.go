package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_putParameterCmd = &cobra.Command{
	Use:   "put-parameter",
	Short: "Create or update a parameter in Parameter Store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_putParameterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_putParameterCmd).Standalone()

		ssm_putParameterCmd.Flags().String("allowed-pattern", "", "A regular expression used to validate the parameter value.")
		ssm_putParameterCmd.Flags().String("data-type", "", "The data type for a `String` parameter.")
		ssm_putParameterCmd.Flags().String("description", "", "Information about the parameter that you want to add to the system.")
		ssm_putParameterCmd.Flags().String("key-id", "", "The Key Management Service (KMS) ID that you want to use to encrypt a parameter.")
		ssm_putParameterCmd.Flags().String("name", "", "The fully qualified name of the parameter that you want to create or update.")
		ssm_putParameterCmd.Flags().Bool("no-overwrite", false, "Overwrite an existing parameter.")
		ssm_putParameterCmd.Flags().Bool("overwrite", false, "Overwrite an existing parameter.")
		ssm_putParameterCmd.Flags().String("policies", "", "One or more policies to apply to a parameter.")
		ssm_putParameterCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		ssm_putParameterCmd.Flags().String("tier", "", "The parameter tier to assign to a parameter.")
		ssm_putParameterCmd.Flags().String("type", "", "The type of parameter that you want to create.")
		ssm_putParameterCmd.Flags().String("value", "", "The parameter value that you want to add to the system.")
		ssm_putParameterCmd.MarkFlagRequired("name")
		ssm_putParameterCmd.Flag("no-overwrite").Hidden = true
		ssm_putParameterCmd.MarkFlagRequired("value")
	})
	ssmCmd.AddCommand(ssm_putParameterCmd)
}
