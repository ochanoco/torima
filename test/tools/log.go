package tools

import (
	"log"
	"reflect"
	"runtime"

	"github.com/ochanoco/torima/core"
)

var STATE = map[int]string{
	0: "AuthNeeded",
	1: "Authed",
	2: "NoAuthNeeded",
	3: "ForceStop",
	4: "Keep",
}

type ExtensionLogger struct {
}

func (logger *ExtensionLogger) Director(count int) core.TorimaDirector {
	return func(c *core.TorimaDirectorPackageContext) (int, error) {
		Log(count, c.Proxy.Directors[count*2+1], c.PackageStatus)
		return core.Keep, nil
	}
}

func (logger *ExtensionLogger) ModifyResp(count int) core.TorimaModifyResponse {
	return func(c *core.TorimaModifyResponsePackageContext) (int, error) {
		Log(count, c.Proxy.Directors[count*2+1], c.PackageStatus)
		return core.Keep, nil
	}
}

func StartOrEndDirector[T *core.TorimaDirectorPackageContext | *core.TorimaModifyResponsePackageContext](c T) (int, error) {
	println("---------------------")
	return core.Keep, nil
}

func (logger *ExtensionLogger) InjectDirectors(source core.TorimaDirectors) core.TorimaDirectors {
	result := core.TorimaDirectors{StartOrEndDirector[*core.TorimaDirectorPackageContext]}

	for i, v := range source {
		d := logger.Director(i)
		result = append(result, d)
		result = append(result, v)
	}

	return result
}

func (logger *ExtensionLogger) InjectModifyResps(source core.TorimaModifyResponses) core.TorimaModifyResponses {
	result := core.TorimaModifyResponses{StartOrEndDirector[*core.TorimaModifyResponsePackageContext]}

	for i, v := range source {
		d := logger.ModifyResp(i)
		result = append(result, d)
		result = append(result, v)
	}

	return result
}

func Log(count int, extension any, result int) {
	rv1 := reflect.ValueOf(extension)
	ptr1 := rv1.Pointer()

	extensionName := runtime.FuncForPC(ptr1).Name()

	log.Printf("id: %v\n", count)
	log.Printf("name: %v\n", extensionName)
	log.Printf("result: %v\n", STATE[result])

}

func DirectorSliceToDirectors(slice []core.TorimaDirector) core.TorimaDirectors {
	var directors core.TorimaDirectors

	for _, s := range slice {
		directors = append(directors, s)
	}

	return directors
}

func ModifyRespSliceToModifyResps(slice []core.TorimaModifyResponse) core.TorimaModifyResponses {
	var modifyResps core.TorimaModifyResponses

	for _, s := range slice {
		modifyResps = append(modifyResps, s)
	}

	return modifyResps
}
