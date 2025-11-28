package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["iot.create-keys-and-certificate"] = func(cmd *command.Command) error {
		if cmd.Completion.Flag == nil {
			cmd.Completion.Flag = map[string][]string{}
		}

		cmd.AddFlag(command.Flag{
			Longhand:    "public-key-outfile",
			Description: "Saves the command output contents of keyPair.PublicKey to the given filename",
			Value:       true,
		})
		cmd.Completion.Flag["public-key-outfile"] = []string{"$files"}

		cmd.AddFlag(command.Flag{
			Longhand:    "certificate-pem-outfile",
			Description: "Saves the command output contents of certificatePem to the given filename",
			Value:       true,
		})
		cmd.Completion.Flag["certificate-pem-outfile"] = []string{"$files"}

		cmd.AddFlag(command.Flag{
			Longhand:    "private-key-outfile",
			Description: "Saves the command output contents of keyPair.PrivateKey to the given filename",
			Value:       true,
		})
		cmd.Completion.Flag["private-key-outfile"] = []string{"$files"}

		return nil
	}

	customizations["iot.create-certificate-from-csr"] = func(cmd *command.Command) error {
		if cmd.Completion.Flag == nil {
			cmd.Completion.Flag = map[string][]string{}
		}

		cmd.AddFlag(command.Flag{
			Longhand:    "certificate-pem-outfile",
			Description: "Saves the command output contents of certificatePem to the given filename",
			Value:       true,
		})
		cmd.Completion.Flag["certificate-pem-outfile"] = []string{"$files"}

		return nil
	}
}
