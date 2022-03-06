package utils

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/templates"
	log "github.com/sirupsen/logrus"
)

/*
GenerateValidatorKey :
Generates a validator key using the eth2.0-deposit-cli tool.
The key can be generated using a new or existing mnemonic.
Key's path is set to $(pwd)/keystore.

params :-
a. new bool
True if the key is to be generated using a new mnemonic. False if the key is to be generated using an existing mnemonic.
b. network string
Target network.

returns :-
a. error
Error if any
*/
func GenerateValidatorKey(new bool, network string) error {
	data := DepositCLI{
		Network: network,
	}

	// Get the template file
	var rawTmp []byte
	var err error
	if new {
		rawTmp, err = templates.DepositCLI.ReadFile("deposit-cli/new.tmpl")
	} else {
		rawTmp, err = templates.DepositCLI.ReadFile("deposit-cli/existing.tmpl")
	}

	if err != nil {
		return err
	}

	// Parse the template
	tmp, err := template.New("deposit-cli").Parse(string(rawTmp))
	if err != nil {
		return err
	}

	// Print cmd
	log.Info(configs.RunningCommand)
	err = tmp.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	fmt.Println()

	// Execute cmd
	script := Script{
		Tmp:       tmp,
		GetOutput: false,
		Data:      data,
	}
	if _, err = executeScript(script); err != nil {
		var scriptBuffer *bytes.Buffer
		err = tmp.Execute(scriptBuffer, data)
		log.Error(err)
		return fmt.Errorf(configs.RunningCMDError, scriptBuffer, err)
	}

	return nil
}