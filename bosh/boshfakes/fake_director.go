// Code generated by counterfeiter. DO NOT EDIT.
package boshfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-deployment-resource/bosh"
)

type FakeDirector struct {
	DeleteStub        func(force bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		force bool
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeployStub        func(manifestBytes []byte, deployParams bosh.DeployParams) error
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		manifestBytes []byte
		deployParams  bosh.DeployParams
	}
	deployReturns struct {
		result1 error
	}
	deployReturnsOnCall map[int]struct {
		result1 error
	}
	InterpolateStub        func(manifestBytes []byte, interpolateParams bosh.InterpolateParams) ([]byte, error)
	interpolateMutex       sync.RWMutex
	interpolateArgsForCall []struct {
		manifestBytes     []byte
		interpolateParams bosh.InterpolateParams
	}
	interpolateReturns struct {
		result1 []byte
		result2 error
	}
	interpolateReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	DownloadManifestStub        func() ([]byte, error)
	downloadManifestMutex       sync.RWMutex
	downloadManifestArgsForCall []struct{}
	downloadManifestReturns     struct {
		result1 []byte
		result2 error
	}
	downloadManifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ExportReleasesStub        func(targetDirectory string, releases []bosh.ReleaseSpec) error
	exportReleasesMutex       sync.RWMutex
	exportReleasesArgsForCall []struct {
		targetDirectory string
		releases        []bosh.ReleaseSpec
	}
	exportReleasesReturns struct {
		result1 error
	}
	exportReleasesReturnsOnCall map[int]struct {
		result1 error
	}
	UploadReleaseStub        func(releaseURL string) error
	uploadReleaseMutex       sync.RWMutex
	uploadReleaseArgsForCall []struct {
		releaseURL string
	}
	uploadReleaseReturns struct {
		result1 error
	}
	uploadReleaseReturnsOnCall map[int]struct {
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
	uploadStemcellReturnsOnCall map[int]struct {
		result1 error
	}
	WaitForDeployLockStub        func() error
	waitForDeployLockMutex       sync.RWMutex
	waitForDeployLockArgsForCall []struct{}
	waitForDeployLockReturns     struct {
		result1 error
	}
	waitForDeployLockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirector) Delete(force bool) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		force bool
	}{force})
	fake.recordInvocation("Delete", []interface{}{force})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(force)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeDirector) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeDirector) DeleteArgsForCall(i int) bool {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].force
}

func (fake *FakeDirector) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) Deploy(manifestBytes []byte, deployParams bosh.DeployParams) error {
	var manifestBytesCopy []byte
	if manifestBytes != nil {
		manifestBytesCopy = make([]byte, len(manifestBytes))
		copy(manifestBytesCopy, manifestBytes)
	}
	fake.deployMutex.Lock()
	ret, specificReturn := fake.deployReturnsOnCall[len(fake.deployArgsForCall)]
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		manifestBytes []byte
		deployParams  bosh.DeployParams
	}{manifestBytesCopy, deployParams})
	fake.recordInvocation("Deploy", []interface{}{manifestBytesCopy, deployParams})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(manifestBytes, deployParams)
	}
	if specificReturn {
		return ret.result1
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

func (fake *FakeDirector) DeployReturnsOnCall(i int, result1 error) {
	fake.DeployStub = nil
	if fake.deployReturnsOnCall == nil {
		fake.deployReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deployReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) Interpolate(manifestBytes []byte, interpolateParams bosh.InterpolateParams) ([]byte, error) {
	var manifestBytesCopy []byte
	if manifestBytes != nil {
		manifestBytesCopy = make([]byte, len(manifestBytes))
		copy(manifestBytesCopy, manifestBytes)
	}
	fake.interpolateMutex.Lock()
	ret, specificReturn := fake.interpolateReturnsOnCall[len(fake.interpolateArgsForCall)]
	fake.interpolateArgsForCall = append(fake.interpolateArgsForCall, struct {
		manifestBytes     []byte
		interpolateParams bosh.InterpolateParams
	}{manifestBytesCopy, interpolateParams})
	fake.recordInvocation("Interpolate", []interface{}{manifestBytesCopy, interpolateParams})
	fake.interpolateMutex.Unlock()
	if fake.InterpolateStub != nil {
		return fake.InterpolateStub(manifestBytes, interpolateParams)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpolateReturns.result1, fake.interpolateReturns.result2
}

func (fake *FakeDirector) InterpolateCallCount() int {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return len(fake.interpolateArgsForCall)
}

func (fake *FakeDirector) InterpolateArgsForCall(i int) ([]byte, bosh.InterpolateParams) {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return fake.interpolateArgsForCall[i].manifestBytes, fake.interpolateArgsForCall[i].interpolateParams
}

func (fake *FakeDirector) InterpolateReturns(result1 []byte, result2 error) {
	fake.InterpolateStub = nil
	fake.interpolateReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDirector) InterpolateReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.InterpolateStub = nil
	if fake.interpolateReturnsOnCall == nil {
		fake.interpolateReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.interpolateReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDirector) DownloadManifest() ([]byte, error) {
	fake.downloadManifestMutex.Lock()
	ret, specificReturn := fake.downloadManifestReturnsOnCall[len(fake.downloadManifestArgsForCall)]
	fake.downloadManifestArgsForCall = append(fake.downloadManifestArgsForCall, struct{}{})
	fake.recordInvocation("DownloadManifest", []interface{}{})
	fake.downloadManifestMutex.Unlock()
	if fake.DownloadManifestStub != nil {
		return fake.DownloadManifestStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
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

func (fake *FakeDirector) DownloadManifestReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.DownloadManifestStub = nil
	if fake.downloadManifestReturnsOnCall == nil {
		fake.downloadManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.downloadManifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDirector) ExportReleases(targetDirectory string, releases []bosh.ReleaseSpec) error {
	var releasesCopy []bosh.ReleaseSpec
	if releases != nil {
		releasesCopy = make([]bosh.ReleaseSpec, len(releases))
		copy(releasesCopy, releases)
	}
	fake.exportReleasesMutex.Lock()
	ret, specificReturn := fake.exportReleasesReturnsOnCall[len(fake.exportReleasesArgsForCall)]
	fake.exportReleasesArgsForCall = append(fake.exportReleasesArgsForCall, struct {
		targetDirectory string
		releases        []bosh.ReleaseSpec
	}{targetDirectory, releasesCopy})
	fake.recordInvocation("ExportReleases", []interface{}{targetDirectory, releasesCopy})
	fake.exportReleasesMutex.Unlock()
	if fake.ExportReleasesStub != nil {
		return fake.ExportReleasesStub(targetDirectory, releases)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.exportReleasesReturns.result1
}

func (fake *FakeDirector) ExportReleasesCallCount() int {
	fake.exportReleasesMutex.RLock()
	defer fake.exportReleasesMutex.RUnlock()
	return len(fake.exportReleasesArgsForCall)
}

func (fake *FakeDirector) ExportReleasesArgsForCall(i int) (string, []bosh.ReleaseSpec) {
	fake.exportReleasesMutex.RLock()
	defer fake.exportReleasesMutex.RUnlock()
	return fake.exportReleasesArgsForCall[i].targetDirectory, fake.exportReleasesArgsForCall[i].releases
}

func (fake *FakeDirector) ExportReleasesReturns(result1 error) {
	fake.ExportReleasesStub = nil
	fake.exportReleasesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) ExportReleasesReturnsOnCall(i int, result1 error) {
	fake.ExportReleasesStub = nil
	if fake.exportReleasesReturnsOnCall == nil {
		fake.exportReleasesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.exportReleasesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) UploadRelease(releaseURL string) error {
	fake.uploadReleaseMutex.Lock()
	ret, specificReturn := fake.uploadReleaseReturnsOnCall[len(fake.uploadReleaseArgsForCall)]
	fake.uploadReleaseArgsForCall = append(fake.uploadReleaseArgsForCall, struct {
		releaseURL string
	}{releaseURL})
	fake.recordInvocation("UploadRelease", []interface{}{releaseURL})
	fake.uploadReleaseMutex.Unlock()
	if fake.UploadReleaseStub != nil {
		return fake.UploadReleaseStub(releaseURL)
	}
	if specificReturn {
		return ret.result1
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

func (fake *FakeDirector) UploadReleaseReturnsOnCall(i int, result1 error) {
	fake.UploadReleaseStub = nil
	if fake.uploadReleaseReturnsOnCall == nil {
		fake.uploadReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) UploadStemcell(stemcellURL string) error {
	fake.uploadStemcellMutex.Lock()
	ret, specificReturn := fake.uploadStemcellReturnsOnCall[len(fake.uploadStemcellArgsForCall)]
	fake.uploadStemcellArgsForCall = append(fake.uploadStemcellArgsForCall, struct {
		stemcellURL string
	}{stemcellURL})
	fake.recordInvocation("UploadStemcell", []interface{}{stemcellURL})
	fake.uploadStemcellMutex.Unlock()
	if fake.UploadStemcellStub != nil {
		return fake.UploadStemcellStub(stemcellURL)
	}
	if specificReturn {
		return ret.result1
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

func (fake *FakeDirector) UploadStemcellReturnsOnCall(i int, result1 error) {
	fake.UploadStemcellStub = nil
	if fake.uploadStemcellReturnsOnCall == nil {
		fake.uploadStemcellReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadStemcellReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) WaitForDeployLock() error {
	fake.waitForDeployLockMutex.Lock()
	ret, specificReturn := fake.waitForDeployLockReturnsOnCall[len(fake.waitForDeployLockArgsForCall)]
	fake.waitForDeployLockArgsForCall = append(fake.waitForDeployLockArgsForCall, struct{}{})
	fake.recordInvocation("WaitForDeployLock", []interface{}{})
	fake.waitForDeployLockMutex.Unlock()
	if fake.WaitForDeployLockStub != nil {
		return fake.WaitForDeployLockStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.waitForDeployLockReturns.result1
}

func (fake *FakeDirector) WaitForDeployLockCallCount() int {
	fake.waitForDeployLockMutex.RLock()
	defer fake.waitForDeployLockMutex.RUnlock()
	return len(fake.waitForDeployLockArgsForCall)
}

func (fake *FakeDirector) WaitForDeployLockReturns(result1 error) {
	fake.WaitForDeployLockStub = nil
	fake.waitForDeployLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) WaitForDeployLockReturnsOnCall(i int, result1 error) {
	fake.WaitForDeployLockStub = nil
	if fake.waitForDeployLockReturnsOnCall == nil {
		fake.waitForDeployLockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitForDeployLockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	fake.exportReleasesMutex.RLock()
	defer fake.exportReleasesMutex.RUnlock()
	fake.uploadReleaseMutex.RLock()
	defer fake.uploadReleaseMutex.RUnlock()
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	fake.waitForDeployLockMutex.RLock()
	defer fake.waitForDeployLockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
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
