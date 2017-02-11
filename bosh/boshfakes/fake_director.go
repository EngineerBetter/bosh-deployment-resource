// This file was generated by counterfeiter
package boshfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-deployment-resource/bosh"
)

type FakeDirector struct {
	DeployStub        func(manifestBytes []byte, deployParams bosh.DeployParams) error
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		manifestBytes []byte
		deployParams  bosh.DeployParams
	}
	deployReturns struct {
		result1 error
	}
	DownloadManifestStub        func() ([]byte, error)
	downloadManifestMutex       sync.RWMutex
	downloadManifestArgsForCall []struct{}
	downloadManifestReturns     struct {
		result1 []byte
		result2 error
	}
	UploadReleaseStub        func(releaseURL string) error
	uploadReleaseMutex       sync.RWMutex
	uploadReleaseArgsForCall []struct {
		releaseURL string
	}
	uploadReleaseReturns struct {
		result1 error
	}
	UploadStemcellStub        func(stemcellURL string) error
	uploadStemcellMutex       sync.RWMutex
	uploadStemcellArgsForCall []struct {
		stemcellURL string
	}
	uploadStemcellReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirector) Deploy(manifestBytes []byte, deployParams bosh.DeployParams) error {
	var manifestBytesCopy []byte
	if manifestBytes != nil {
		manifestBytesCopy = make([]byte, len(manifestBytes))
		copy(manifestBytesCopy, manifestBytes)
	}
	fake.deployMutex.Lock()
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		manifestBytes []byte
		deployParams  bosh.DeployParams
	}{manifestBytesCopy, deployParams})
	fake.recordInvocation("Deploy", []interface{}{manifestBytesCopy, deployParams})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(manifestBytes, deployParams)
	}
	return fake.deployReturns.result1
}

func (fake *FakeDirector) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *FakeDirector) DeployArgsForCall(i int) ([]byte, bosh.DeployParams) {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return fake.deployArgsForCall[i].manifestBytes, fake.deployArgsForCall[i].deployParams
}

func (fake *FakeDirector) DeployReturns(result1 error) {
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) DownloadManifest() ([]byte, error) {
	fake.downloadManifestMutex.Lock()
	fake.downloadManifestArgsForCall = append(fake.downloadManifestArgsForCall, struct{}{})
	fake.recordInvocation("DownloadManifest", []interface{}{})
	fake.downloadManifestMutex.Unlock()
	if fake.DownloadManifestStub != nil {
		return fake.DownloadManifestStub()
	}
	return fake.downloadManifestReturns.result1, fake.downloadManifestReturns.result2
}

func (fake *FakeDirector) DownloadManifestCallCount() int {
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	return len(fake.downloadManifestArgsForCall)
}

func (fake *FakeDirector) DownloadManifestReturns(result1 []byte, result2 error) {
	fake.DownloadManifestStub = nil
	fake.downloadManifestReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDirector) UploadRelease(releaseURL string) error {
	fake.uploadReleaseMutex.Lock()
	fake.uploadReleaseArgsForCall = append(fake.uploadReleaseArgsForCall, struct {
		releaseURL string
	}{releaseURL})
	fake.recordInvocation("UploadRelease", []interface{}{releaseURL})
	fake.uploadReleaseMutex.Unlock()
	if fake.UploadReleaseStub != nil {
		return fake.UploadReleaseStub(releaseURL)
	}
	return fake.uploadReleaseReturns.result1
}

func (fake *FakeDirector) UploadReleaseCallCount() int {
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	return len(fake.uploadReleaseArgsForCall)
}

func (fake *FakeDirector) UploadReleaseArgsForCall(i int) string {
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	return fake.uploadReleaseArgsForCall[i].releaseURL
}

func (fake *FakeDirector) UploadReleaseReturns(result1 error) {
	fake.UploadReleaseStub = nil
	fake.uploadReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) UploadStemcell(stemcellURL string) error {
	fake.uploadStemcellMutex.Lock()
	fake.uploadStemcellArgsForCall = append(fake.uploadStemcellArgsForCall, struct {
		stemcellURL string
	}{stemcellURL})
	fake.recordInvocation("UploadStemcell", []interface{}{stemcellURL})
	fake.uploadStemcellMutex.Unlock()
	if fake.UploadStemcellStub != nil {
		return fake.UploadStemcellStub(stemcellURL)
	}
	return fake.uploadStemcellReturns.result1
}

func (fake *FakeDirector) UploadStemcellCallCount() int {
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	return len(fake.uploadStemcellArgsForCall)
}

func (fake *FakeDirector) UploadStemcellArgsForCall(i int) string {
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	return fake.uploadStemcellArgsForCall[i].stemcellURL
}

func (fake *FakeDirector) UploadStemcellReturns(result1 error) {
	fake.UploadStemcellStub = nil
	fake.uploadStemcellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDirector) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ bosh.Director = new(FakeDirector)
