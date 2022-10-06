package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	cst "thy/constants"
	"thy/errors"
	"thy/internal/predictor"
	"thy/paths"
	"thy/utils"
	"thy/vaultcli"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/cli"
	"github.com/spf13/viper"
)

func GetPoolCmd() (cli.Command, error) {
	return NewCommand(CommandArgs{
		Path:         []string{cst.NounPool},
		SynopsisText: "Manage engine pools",
		HelpText:     "Work with engine pools",
		RunFunc: func(vcli vaultcli.CLI, args []string) int {
			path := viper.GetString(cst.Path)
			if path == "" && len(args) > 0 && !strings.HasPrefix(args[0], "-") {
				path = args[0]
			}
			if path == "" {
				return cli.RunResultHelp
			}
			return handlePoolRead(vcli, args)
		},
	})
}

func GetPoolCreateCmd() (cli.Command, error) {
	return NewCommand(CommandArgs{
		Path:         []string{cst.NounPool, cst.Create},
		SynopsisText: "Create a new empty pool of engines",
		HelpText: fmt.Sprintf(`
Usage:
   • %[1]s %[2]s --%[3]s mypool`, cst.NounPool, cst.Create, cst.DataName),
		FlagsPredictor: []*predictor.Params{
			{Shorthand: "n", Name: cst.DataName, Usage: fmt.Sprintf("Name of the %s (required)", cst.NounPool)},
		},
		RunFunc: func(vcli vaultcli.CLI, args []string) int {
			if OnlyGlobalArgs(args) {
				return handlePoolCreateWizard(vcli)
			}
			return handlePoolCreate(vcli, args)
		},
	})
}

func GetPoolReadCmd() (cli.Command, error) {
	return NewCommand(CommandArgs{
		Path:         []string{cst.NounPool, cst.Read},
		SynopsisText: "Get information on an existing pool of engines",
		HelpText: fmt.Sprintf(`
Usage:
   • %[1]s %[2]s --%[3]s mypool`, cst.NounPool, cst.Read, cst.DataName),
		FlagsPredictor: []*predictor.Params{
			{Shorthand: "n", Name: cst.DataName, Usage: fmt.Sprintf("Name of the %s (required)", cst.NounPool)},
		},
		MinNumberArgs: 1,
		RunFunc: func(vcli vaultcli.CLI, args []string) int {
			return handlePoolRead(vcli, args)
		},
	})
}

func GetPoolListCmd() (cli.Command, error) {
	return NewCommand(CommandArgs{
		Path:         []string{cst.NounPool, cst.List},
		SynopsisText: "List the names of all existing pools",
		HelpText: fmt.Sprintf(`
Usage:
   • %[1]s %[2]s`, cst.NounPool, cst.List),
		RunFunc: func(vcli vaultcli.CLI, args []string) int {
			return handlePoolList(vcli, args)
		},
	})
}

func GetPoolDeleteCmd() (cli.Command, error) {
	return NewCommand(CommandArgs{
		Path:         []string{cst.NounPool, cst.Delete},
		SynopsisText: "Delete an existing pool of engines",
		HelpText: fmt.Sprintf(`
Usage:
   • %[1]s %[2]s --%[3]s mypool`, cst.NounPool, cst.Delete, cst.DataName),
		FlagsPredictor: []*predictor.Params{
			{Shorthand: "n", Name: cst.DataName, Usage: fmt.Sprintf("Name of the %s (required)", cst.NounPool)},
		},
		MinNumberArgs: 1,
		RunFunc: func(vcli vaultcli.CLI, args []string) int {
			return handlePoolDelete(vcli, args)
		},
	})
}

func handlePoolRead(vcli vaultcli.CLI, args []string) int {
	name := viper.GetString(cst.DataName)
	if name == "" && len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		name = args[0]
	}
	if name == "" {
		err := errors.NewS("error: must specify " + cst.DataName)
		vcli.Out().WriteResponse(nil, err)
		return utils.GetExecStatus(err)
	}

	data, apiErr := poolRead(vcli, paths.ProcessResource(name))
	vcli.Out().WriteResponse(data, apiErr)
	return utils.GetExecStatus(apiErr)
}

func handlePoolCreate(vcli vaultcli.CLI, args []string) int {
	name := viper.GetString(cst.DataName)
	if name == "" {
		err := errors.NewS("error: must specify " + cst.DataName)
		vcli.Out().WriteResponse(nil, err)
		return utils.GetExecStatus(err)
	}
	if err := vaultcli.ValidateName(name); err != nil {
		vcli.Out().FailF("error: pool name %q is invalid: %v", name, err)
		return utils.GetExecStatus(err)
	}

	data, apiErr := poolCreate(vcli, name)
	vcli.Out().WriteResponse(data, apiErr)
	return utils.GetExecStatus(apiErr)
}

func handlePoolList(vcli vaultcli.CLI, args []string) int {
	data, apiErr := poolList(vcli)
	vcli.Out().WriteResponse(data, apiErr)
	return utils.GetExecStatus(apiErr)
}

func handlePoolDelete(vcli vaultcli.CLI, args []string) int {
	name := viper.GetString(cst.DataName)
	if name == "" && len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		name = args[0]
	}
	if name == "" {
		err := errors.NewS("error: must specify " + cst.DataName)
		vcli.Out().WriteResponse(nil, err)
		return utils.GetExecStatus(err)
	}

	data, apiErr := poolDelete(vcli, paths.ProcessResource(name))
	vcli.Out().WriteResponse(data, apiErr)
	return utils.GetExecStatus(apiErr)
}

// Wizards:

func handlePoolCreateWizard(vcli vaultcli.CLI) int {
	var name string
	namePrompt := &survey.Input{Message: "Pool name:"}
	survErr := survey.AskOne(namePrompt, &name, survey.WithValidator(vaultcli.SurveyRequiredName))
	if survErr != nil {
		vcli.Out().WriteResponse(nil, errors.New(survErr))
		return utils.GetExecStatus(survErr)
	}

	data, apiErr := poolCreate(vcli, strings.TrimSpace(name))
	vcli.Out().WriteResponse(data, apiErr)
	return utils.GetExecStatus(apiErr)
}

// API callers:

func poolCreate(vcli vaultcli.CLI, name string) ([]byte, *errors.ApiError) {
	body := map[string]string{"name": name}
	uri := paths.CreateResourceURI(cst.NounPools, "", "", true, nil)
	return vcli.HTTPClient().DoRequest(http.MethodPost, uri, &body)
}

func poolRead(vcli vaultcli.CLI, name string) ([]byte, *errors.ApiError) {
	uri := paths.CreateResourceURI(cst.NounPools, name, "", true, nil)
	return vcli.HTTPClient().DoRequest(http.MethodGet, uri, nil)
}

func poolDelete(vcli vaultcli.CLI, name string) ([]byte, *errors.ApiError) {
	query := map[string]string{"force": strconv.FormatBool(true)}
	uri := paths.CreateResourceURI(cst.NounPools, name, "", true, query)
	return vcli.HTTPClient().DoRequest(http.MethodDelete, uri, nil)
}

func poolList(vcli vaultcli.CLI) ([]byte, *errors.ApiError) {
	uri := paths.CreateResourceURI(cst.NounPools, "", "", false, nil)
	return vcli.HTTPClient().DoRequest(http.MethodGet, uri, nil)
}
