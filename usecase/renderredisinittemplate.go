package usecase

import (
	"fmt"
	"os"

	coreusecase "github.com/golangspell/golangspell-core/usecase"
	"github.com/golangspell/golangspell-redis/appcontext"
	"github.com/golangspell/golangspell-redis/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
)

func addComponentConstantToContext(currentPath string, componentName string) error {
	return coreusecase.GetAddComponentConstantToContext().Execute(currentPath, componentName)
}

func addImportToMain(moduleName string, currentPath string, importPath string) error {
	return coreusecase.GetAddPackageImportToMain().Execute(moduleName, currentPath, importPath)
}

func addEnvironmentVariables(currentPath string) error {
	err := coreusecase.GetAddEnvironmentVariable().Execute(currentPath, "RedisAddress", "string", "`env:\"REDIS_ADDRESS\" envDefault:\"\"`")
	if err != nil {
		fmt.Printf("An error occurred while trying to update the environment file. Error: %s\n", err.Error())
		return err
	}
	return coreusecase.GetAddEnvironmentVariable().Execute(currentPath, "RedisPassword", "string", "`env:\"REDIS_PASSWORD\" envDefault:\"\"`")
}

// RenderredisinitTemplate renders the templates defined to the redisinit command with the proper variables
func RenderredisinitTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to init redis. Error: %s\n", err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	globalVariables := map[string]interface{}{
		"ModuleName": moduleName,
	}

	err = renderer.RenderTemplate(spell, "redisinit", globalVariables, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		"Cache")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		"Lock")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addImportToMain(
		moduleName,
		currentPath,
		fmt.Sprintf("%s/gateway/redis", moduleName))
	if err != nil {
		fmt.Printf("An error occurred while trying to add the import to main. Error: %s\n", err.Error())
		return err
	}

	err = addEnvironmentVariables(currentPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to configure the environment. Error: %s\n", err.Error())
		return err
	}

	return nil
}
